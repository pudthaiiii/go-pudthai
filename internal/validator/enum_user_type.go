// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/validator/enum_user_type.go

/*
Package validator เป็น package ที่ใช้สำหรับการตรวจสอบความถูกต้องของข้อมูล

ไฟล์นี้ใช้สำหรับจัดการการตรวจสอบความถูกต้องของข้อมูล (Validation) ในแอปพลิเคชัน
*/
package validator

import (
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func init() {
	validate.RegisterValidation("enum_user_type", enumUserType)

	validate.RegisterTranslation("enum_user_type", trans, func(ut ut.Translator) error {
		return ut.Add("enum_user_type", "{0} must be one of [Admin, Merchant, User]", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("enum_user_type", fe.Field())
		return t
	})
}

func enumUserType(fl validator.FieldLevel) bool {
	typeField := strings.ToLower(fl.Field().String())

	// ตรวจสอบว่า typeField ต้องเป็น admin, merchant หรือ user
	return typeField == "admin" || typeField == "merchant" || typeField == "user"
}
