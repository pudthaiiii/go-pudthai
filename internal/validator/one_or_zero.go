// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/validator/one_or_zero.go

/*
Package validator เป็น package ที่ใช้สำหรับการตรวจสอบความถูกต้องของข้อมูล

ไฟล์นี้ใช้สำหรับจัดการการตรวจสอบความถูกต้องของข้อมูล (Validation) ในแอปพลิเคชัน
*/
package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ฟังก์ชัน init สำหรับลงทะเบียน custom validation "oneOrZero"
func init() {
	validate.RegisterValidation("one_or_zero", validateOneOrZero)

	validate.RegisterTranslation("one_or_zero", trans, func(ut ut.Translator) error {
		return ut.Add("one_or_zero", "{0} must be either 0 or 1", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("one_or_zero", fe.Field())
		return t
	})
}

// ฟังก์ชัน validateOneOrZero ใช้ตรวจสอบว่า ค่าที่ตรวจสอบต้องเป็น 0 หรือ 1 เท่านั้น
func validateOneOrZero(fl validator.FieldLevel) bool {
	value := fl.Field().Int() // ดึงค่าที่จะตรวจสอบ

	// คืนค่า true ถ้า value เป็น 0 หรือ 1, คืนค่า false ถ้าไม่ใช่
	return value == 0 || value == 1
}
