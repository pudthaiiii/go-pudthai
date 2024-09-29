// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: cmd/api/shared.go

/*
Package api ใช้สำหรับการสร้างแอปพลิเคชัน API

โดยจะมีการตั้งค่าและเริ่มต้นการทำงานของแอปพลิเคชัน
*/
package api

import (
	"encoding/json"
	"go-pudthai/internal/utils"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// validationError โครงสร้างสำหรับเก็บข้อผิดพลาดในการตรวจสอบข้อมูล
type validationError struct {
	Errors  map[string][]string `json:"errors"`  // แผนที่ของข้อผิดพลาด
	Message string              `json:"message"` // ข้อความของข้อผิดพลาด
}

// errorHandler จัดการข้อผิดพลาดที่เกิดขึ้นในแอปพลิเคชัน
func errorHandler(c *fiber.Ctx, err error) error {
	// กำหนดค่าพื้นฐานสำหรับสถานะข้อผิดพลาด
	statusCode, errorCode, errorMessage, exception := http.StatusInternalServerError, 0, "Internal Server Error", []string{}

	// ตรวจสอบว่าข้อผิดพลาดคือ fiber.Error หรือไม่
	if fiberErr, ok := err.(*fiber.Error); ok {
		handleFiberError(fiberErr, &statusCode, &errorCode, &errorMessage, &exception)
	} else if err != nil {
		handleGenericError(err, &statusCode, &errorCode, &errorMessage, &exception)
	}

	// สร้างการตอบกลับข้อผิดพลาด
	response := createErrorResponse(errorMessage, errorCode, exception)
	return c.Status(statusCode).JSON(response) // ส่งการตอบกลับ JSON
}

// handleFiberError จัดการข้อผิดพลาดที่เกิดจาก Fiber
func handleFiberError(err *fiber.Error, statusCode, errorCode *int, errorMessage *string, exception *[]string) {
	*statusCode, *errorCode, *errorMessage = err.Code, err.Code, err.Message

	// ตรวจสอบข้อผิดพลาดการตรวจสอบข้อมูล
	if strings.Contains(err.Error(), "VALIDATE_ERROR") {
		*statusCode = http.StatusUnprocessableEntity
		*errorCode = 900422
		*errorMessage = "VALIDATE_ERROR"

		// แปลงข้อผิดพลาด JSON เป็น validationError
		var validationErr validationError
		if json.Unmarshal([]byte(err.Error()), &validationErr) == nil {
			for _, errs := range validationErr.Errors {
				*exception = append(*exception, errs...) // เก็บข้อผิดพลาดทั้งหมด
			}
		}
	}
}

// handleGenericError จัดการข้อผิดพลาดทั่วไป
func handleGenericError(err error, statusCode, errorCode *int, errorMessage *string, exception *[]string) {
	throwException := utils.FilterThrowExceptions(err.Error())
	if len(throwException) > 0 {
		*statusCode = http.StatusUnprocessableEntity

		// ตรวจสอบรหัสข้อผิดพลาด
		if len(throwException) > 1 && throwException[1] != "" {
			*errorCode, _ = strconv.Atoi(throwException[1])
		}

		// จัดการข้อผิดพลาดเฉพาะ
		if len(throwException) > 2 && throwException[2] != "" {
			handleSpecificErrorCases(throwException[2], statusCode, errorMessage)
		}

		// เพิ่มข้อผิดพลาดเพิ่มเติม
		if len(throwException) > 3 && throwException[3] != "" {
			*exception = append(*exception, throwException[3])
		}
	}
}

// handleSpecificErrorCases จัดการกรณีข้อผิดพลาดเฉพาะ
func handleSpecificErrorCases(errorDetail string, statusCode *int, errorMessage *string) {
	// ตรวจสอบข้อผิดพลาดด้านการรับรองความถูกต้อง
	if strings.Contains(errorDetail, "AUTH_") {
		*statusCode = http.StatusUnauthorized
	}

	// ตรวจสอบข้อผิดพลาดที่ไม่พบ
	if strings.Contains(errorDetail, "NOT_FOUND") {
		*statusCode = http.StatusNotFound
	}

	*errorMessage = errorDetail
}

// createErrorResponse สร้างการตอบกลับข้อผิดพลาด
func createErrorResponse(errorMessage string, errorCode int, exception []string) fiber.Map {
	response := fiber.Map{
		"message": errorMessage, // ข้อความข้อผิดพลาด
		"code":    errorCode,    // รหัสข้อผิดพลาด
	}

	// ตรวจสอบการตั้งค่าว่าจะบันทึกข้อผิดพลาดหรือไม่
	isEnabled, _ := strconv.ParseBool(os.Getenv("EXCEPTION_LOG_ENABLED"))
	isValidate := strings.Contains(errorMessage, "VALIDATE_ERROR")

	// ถ้ามีข้อผิดพลาดเพิ่มเติมให้เพิ่มในการตอบกลับ
	if (isEnabled && len(exception) > 0) || (isValidate && len(exception) > 0) {
		response["errors"] = exception // เพิ่มข้อผิดพลาดลงใน response
	}

	return response // ส่งกลับ response
}
