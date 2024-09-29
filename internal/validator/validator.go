// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: cmd/api/application.go

/*
Package validator เป็น package ที่ใช้สำหรับการตรวจสอบความถูกต้องของข้อมูล
ไฟล์นี้ใช้สำหรับจัดการการตรวจสอบความถูกต้องของข้อมูล (Validation) ในแอปพลิเคชัน
*/
package validator

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

// ใช้ validator จาก go-playground
var (
	validate = validator.New()
	enLocale = en.New()
	uni      = ut.New(enLocale, enLocale)
	trans, _ = uni.GetTranslator("en")
)

func init() {
	// Register default translations for built-in validators
	enTranslations.RegisterDefaultTranslations(validate, trans)
}

type (
	// errorResponse ใช้ในการเก็บข้อมูลข้อผิดพลาดเมื่อการตรวจสอบล้มเหลว
	errorResponse struct {
		FailedField string `json:"failedField"`
		Error       string `json:"error"`
	}

	// xValidator struct ที่เป็นตัวแทนของ validator ที่ใช้งาน
	xValidator struct {
		validator *validator.Validate
	}
)

// ฟังก์ชัน inValidate สำหรับตรวจสอบ struct ที่ได้รับและรวม error message
func (v xValidator) inValidate(data interface{}) []errorResponse {
	validationErrors := []errorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		// Loop through the validation errors
		for _, err := range errs.(validator.ValidationErrors) {
			var elem errorResponse

			message := strings.ReplaceAll(err.Translate(trans), err.Field(), strcase.ToLowerCamel(err.Field()))

			elem.FailedField = err.Field()
			elem.Error = message

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

// ฟังก์ชัน Validate สำหรับตรวจสอบข้อมูลที่ได้รับจาก request
func Validate[T any](c *fiber.Ctx, data *T) error {
	myValidator := &xValidator{
		validator: validate,
	}

	var err error
	// Check request method and parse data
	switch c.Method() {
	case "GET":
		err = c.QueryParser(data)
	default:
		err = c.BodyParser(data)
	}

	// Return parsing error if any
	if err != nil {
		msg := fmt.Sprintf("%v", err)
		return handleValidationError(c, msg) // ส่ง error กลับไปยัง Fiber
	}

	// Validate the parsed data and check for validation errors
	if errs := myValidator.inValidate(data); len(errs) > 0 {
		// Map the errors for better representation
		errorMap := make(map[string][]string)
		for _, err := range errs {
			failedField := err.FailedField
			if _, ok := errorMap[failedField]; !ok {
				errorMap[failedField] = []string{}
			}

			// Append the translated error message
			errorMap[failedField] = append(errorMap[failedField], err.Error)
		}

		return handleValidationError(c, errorMap) // ส่ง error กลับไปยัง Fiber
	}

	return nil
}

// ฟังก์ชัน handleValidationError สำหรับจัดการ error ที่เกิดจากการตรวจสอบ
func handleValidationError(_ *fiber.Ctx, errors interface{}) error {
	// Create a response map with error details
	response := map[string]interface{}{
		"message": "VALIDATE_ERROR",
		"errors":  errors,
	}

	// Convert the map to JSON format
	jsonData, _ := json.Marshal(response)

	jsonString := string(jsonData)
	// Return the formatted error response
	return fiber.NewError(fiber.StatusUnprocessableEntity, jsonString)
}
