// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/validator/one_or_zero.go

/*
Package validator เป็น package ที่ใช้สำหรับการตรวจสอบความถูกต้องของข้อมูล

ไฟล์นี้ใช้สำหรับจัดการการตรวจสอบความถูกต้องของข้อมูล (Validation) ในแอปพลิเคชัน
*/
package validator

import "github.com/go-playground/validator/v10"

// ฟังก์ชัน init สำหรับลงทะเบียน custom validation "oneOrZero"
func init() {
	validate.RegisterValidation("oneOrZero", validateOneOrZero)
}

// ฟังก์ชัน validateOneOrZero ใช้ตรวจสอบว่า ค่าที่ตรวจสอบต้องเป็น 0 หรือ 1 เท่านั้น
func validateOneOrZero(fl validator.FieldLevel) bool {
	value := fl.Field().Int() // ดึงค่าที่จะตรวจสอบ

	// คืนค่า true ถ้า value เป็น 0 หรือ 1, คืนค่า false ถ้าไม่ใช่
	return value == 0 || value == 1
}
