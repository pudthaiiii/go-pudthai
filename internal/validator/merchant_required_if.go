// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/validator/merchant_required_if.go

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
	validate.RegisterValidation("required_if_type_not_admin", merchantRequiredIfNotAdmin)

	validate.RegisterTranslation("required_if_type_not_admin", trans, func(ut ut.Translator) error {
		return ut.Add("required_if_type_not_admin", "{0} is required when Type is not ADMIN", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_if_type_not_admin", fe.Field())
		return t
	})
}

func merchantRequiredIfNotAdmin(fl validator.FieldLevel) bool {
	typeField := fl.Parent().FieldByName("Type").String()

	merchantID := fl.Field().Uint()

	if strings.ToUpper(typeField) != "ADMIN" {
		return merchantID != 0
	}

	return true
}
