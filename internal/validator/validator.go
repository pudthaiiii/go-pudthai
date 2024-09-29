// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/validator/validator.go

/*
Package validator เป็น package ที่ใช้สำหรับการตรวจสอบความถูกต้องของข้อมูล

ไฟล์นี้ใช้สำหรับจัดการการตรวจสอบความถูกต้องของข้อมูล (Validation) ในแอปพลิเคชัน
*/
package validator

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/iancoleman/strcase"
)

// ใช้ validator จาก go-playground
var validate = validator.New()

type (
	// errorResponse ใช้ในการเก็บข้อมูลข้อผิดพลาดเมื่อการตรวจสอบล้มเหลว
	errorResponse struct {
		Error       bool
		FailedField string
		Tag         string
	}

	// xValidator struct ที่เป็นตัวแทนของ validator ที่ใช้งาน
	xValidator struct {
		validator *validator.Validate
	}
)

// ฟังก์ชัน inValidate สำหรับตรวจสอบ struct ที่ได้รับ
func (v xValidator) inValidate(data interface{}) []errorResponse {
	validationErrors := []errorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem errorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Error = true

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
	switch c.Method() {
	case "GET":
		err = c.QueryParser(data)
	default:
		err = c.BodyParser(data)
	}

	if err != nil {
		msg := fmt.Sprintf("%v", err)
		return handleValidationError(c, msg) // ส่ง error กลับไปยัง Fiber
	}

	if errs := myValidator.inValidate(data); len(errs) > 0 && errs[0].Error {
		errorMap := make(map[string][]string)
		for _, err := range errs {
			failedField := strcase.ToLowerCamel(err.FailedField)

			if _, ok := errorMap[failedField]; !ok {
				errorMap[failedField] = []string{}
			}

			errorMsg := fmt.Sprintf("%v failed on the '%s'", failedField, err.Tag)
			errorMap[failedField] = append(errorMap[failedField], errorMsg)
		}

		return handleValidationError(c, errorMap) // ส่ง error กลับไปยัง Fiber
	}

	return nil
}

// ฟังก์ชัน handleValidationError สำหรับจัดการ error ที่เกิดจากการตรวจสอบ
func handleValidationError(_ *fiber.Ctx, errors interface{}) error {
	response := map[string]interface{}{
		"message": "VALIDATE_ERROR",
		"errors":  errors,
	}

	jsonData, _ := json.Marshal(response)

	jsonString := string(jsonData)
	// คืนค่า error ที่ถูก format เพื่อให้ Fiber จัดการ
	return fiber.NewError(fiber.StatusUnprocessableEntity, jsonString)
}
