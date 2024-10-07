// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: cmd/api/shared.go

/*
Package api ใช้สำหรับการสร้างแอปพลิเคชัน API

โดยจะมีการตั้งค่าและเริ่มต้นการทำงานของแอปพลิเคชัน
*/
package api

import (
	"encoding/json"
	"fmt"
	"go-pudthai/internal/model/technical"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type validationError struct {
	Errors  map[string][]string `json:"errors"`
	Message string              `json:"message"`
}

type errorConstant struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

// errorHandler จัดการข้อผิดพลาดที่เกิดขึ้นในแอปพลิเคชัน
func errorHandler(c *fiber.Ctx, err error) error {

	fmt.Println("errorHandler", err, c)
	statusCode := http.StatusInternalServerError
	errorCode := 0
	errorMessage := "Internal Server Error"
	exception := []string{}

	if fiberErr, ok := err.(*fiber.Error); ok {
		handleFiberError(fiberErr, &statusCode, &errorCode, &errorMessage, &exception)
	} else {
		handleUnexpectedError(err, &errorCode, &errorMessage, &exception)
	}

	response := createErrorResponse(statusCode, errorMessage, errorCode, exception)

	return c.Status(statusCode).JSON(response)
}

// handleFiberError จัดการข้อผิดพลาดที่เกิดจาก Fiber
func handleFiberError(err *fiber.Error, statusCode, errorCode *int, errorMessage *string, exception *[]string) {
	*statusCode = err.Code
	*errorCode = err.Code
	*errorMessage = err.Message

	fmt.Println("handleFiberError", err)

	if strings.Contains(err.Error(), "VALIDATE_ERROR") {
		handleValidationError(err.Error(), statusCode, errorCode, errorMessage, exception)
	} else {
		handleGeneralError(errorMessage, errorCode, exception)
	}
}

// handleValidationError จัดการข้อผิดพลาดการตรวจสอบข้อมูล
func handleValidationError(errMsg string, statusCode, errorCode *int, errorMessage *string, exception *[]string) {
	*statusCode = http.StatusUnprocessableEntity
	*errorCode = 900422
	*errorMessage = "VALIDATE_ERROR"

	var validationErr validationError
	if json.Unmarshal([]byte(errMsg), &validationErr) == nil {
		for _, errs := range validationErr.Errors {
			*exception = append(*exception, errs...)
		}
	}
}

// handleGeneralError จัดการข้อผิดพลาดทั่วไป
func handleGeneralError(errorMessage *string, errorCode *int, exception *[]string) {
	var result errorConstant
	if err := json.Unmarshal([]byte(*errorMessage), &result); err == nil {
		*errorCode = result.Code
		*errorMessage = result.Message
		if result.Error != "" {
			*exception = append(*exception, result.Error)
		}
	}
}

// handleUnexpectedError จัดการข้อผิดพลาดที่ไม่คาดคิด
func handleUnexpectedError(err error, errorCode *int, errorMessage *string, exception *[]string) {
	fmt.Println("handleUnexpectedError", err)

	*errorMessage = err.Error()
	*exception = append(*exception, "An unexpected error occurred: "+*errorMessage)
	*errorCode = 1000
}

// createErrorResponse สร้างการตอบกลับข้อผิดพลาด
func createErrorResponse(statusCode int, errorMessage string, errorCode int, exception []string) fiber.Map {
	response := fiber.Map{
		"status": fiber.Map{
			"code":    statusCode,
			"message": technical.HttpStatusMessages[statusCode],
		},
		"error": fiber.Map{
			"code":    errorCode,
			"message": errorMessage,
		},
	}

	// ตรวจสอบการตั้งค่าว่าจะบันทึกข้อผิดพลาดหรือไม่
	isEnabled, _ := strconv.ParseBool(os.Getenv("EXCEPTION_LOG_ENABLED"))
	isValidate := strings.Contains(errorMessage, "VALIDATE_ERROR")

	// ถ้ามีข้อผิดพลาดเพิ่มเติมให้เพิ่มในการตอบกลับ
	if (isEnabled && len(exception) > 0) || (isValidate && len(exception) > 0) {
		if statusMap, ok := response["error"].(fiber.Map); ok {
			if len(exception) > 0 {
				statusMap["errors"] = exception
			}
		}
	}

	return response
}
