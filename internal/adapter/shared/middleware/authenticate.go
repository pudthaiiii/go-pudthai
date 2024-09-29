// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/adapter/shared/middleware/authenticate.go

/*
Package middleware เป็น package ที่ใช้สำหรับการจัดการ middleware ต่างๆ ของแอปพลิเคชัน

ไฟล์นี้ใช้สำหรับการจัดการ middleware ที่เกี่ยวกับการตรวจสอบและยืนยัน token ของผู้ใช้
*/
package middleware

import (
	"context"
	"fmt"
	"go-pudthai/internal/model/business"
	t "go-pudthai/internal/model/technical"
	"go-pudthai/internal/throw"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// Authenticate ทำหน้าที่ตรวจสอบและยืนยัน token ที่ส่งมาจากผู้ใช้
func (m *middleware) Authenticate(handler fiber.Handler, action string, subject string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// หาก action เป็น NONE ให้ข้ามการตรวจสอบ token
		if isAuthenticate(action) {
			return handler(c)
		}

		// ดึง token จาก Authorization header
		tokenString := c.Get("Authorization")
		if err := m.validateTokenString(tokenString); err != nil {
			return err
		}

		// ดึง secret สำหรับการยืนยัน token
		secret := m.getSecret(c.Route().Path)
		token, err := m.parseToken(tokenString[7:], secret) // ลบ "Bearer " ออกจาก token
		if err != nil {
			return err
		}

		// ตรวจสอบ claims ของ token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return m.unauthorizedResponse() // ถ้าไม่สามารถแปลง claims ได้ ให้ส่ง response Unauthorized
		}

		// ค้นหาผู้ใช้โดยใช้ token
		user, err := m.accessTokenRepo.FindUserByToken(c.Context(), claims["token"].(string))
		if err != nil {
			return m.unauthorizedResponse()
		}

		// ตั้งค่าข้อมูลผู้ใช้ใน context
		m.setUserLocals(c, user)

		// หากไม่ใช่เส้นทางของ admin ให้ตั้งค่า merchant ใน context
		if !strings.Contains(c.Route().Path, "/v1/admin") {
			if err := m.setMerchantLocals(c, user.MerchantID); err != nil {
				return err // หากไม่พบ merchant ให้คืนค่า error
			}
		}

		return handler(c) // ส่งต่อ request ไปยัง handler ถ้าผ่านการตรวจสอบ
	}
}

// validateTokenString ตรวจสอบว่า token ที่ส่งมาถูกต้องหรือไม่
func (m *middleware) validateTokenString(tokenString string) error {
	if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer ") {
		return m.unauthorizedResponse() // ส่ง response Unauthorized หาก token ไม่ถูกต้อง
	}
	return nil
}

// parseToken ทำการแปลง token และตรวจสอบความถูกต้อง
func (m *middleware) parseToken(tokenString, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, throw.InvalidJwtToken(fmt.Errorf("unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, m.unauthorizedResponse() // ส่ง response Unauthorized หาก token ไม่ถูกต้อง
	}

	return token, nil // คืนค่า token ที่ถูกต้อง
}

// unauthorizedResponse ส่ง response ที่มีสถานะ Unauthorized
func (m *middleware) unauthorizedResponse() error {
	return throw.UserCredentialMismatch()
}

// isAuthenticate ตรวจสอบว่า action เป็น NONE หรือไม่
func isAuthenticate(action string) bool {
	return action == string(t.NONE)
}

// getSecret ดึง secret key สำหรับการยืนยัน token ตามเส้นทางที่ระบุ
func (m *middleware) getSecret(path string) string {
	secrets := m.cfg.Get("JWT")
	switch {
	case strings.Contains(path, "/v1/admin"):
		return secrets["JwtSecretAdmin"].(string)
	case strings.Contains(path, "/v1/backend"):
		return secrets["JwtSecretBackend"].(string)
	case strings.Contains(path, "/v1/frontend"):
		return secrets["JwtSecret"].(string)
	default:
		return secrets["JwtSecret"].(string)
	}
}

// setUserLocals ตั้งค่าข้อมูลผู้ใช้ใน context ของ Fiber
func (m *middleware) setUserLocals(c *fiber.Ctx, user business.GetUserResult) {
	c.Locals(t.Member, user)                              // ตั้งค่าข้อมูลผู้ใช้ใน Locals
	ctx := context.WithValue(c.Context(), t.Member, user) // เพิ่มข้อมูลผู้ใช้ใน context
	c.SetUserContext(ctx)                                 // ตั้งค่า context ใหม่
}

// setMerchantLocals ค้นหาและตั้งค่าข้อมูล merchant ใน context ของ Fiber
func (m *middleware) setMerchantLocals(c *fiber.Ctx, merchantID uint) error {
	merchant, err := m.merchantRepo.FindByID(c.Context(), merchantID)
	if err != nil {
		return throw.MerchantNotFound() // ส่ง error หากไม่พบ merchant
	}

	c.Locals(t.Merchant, merchant)     // ตั้งค่า merchant ใน Locals
	c.Locals(t.MerchantID, merchantID) // ตั้งค่า MerchantID ใน Locals

	ctx := context.WithValue(c.UserContext(), t.MerchantID, merchantID) // เพิ่ม MerchantID ใน context
	ctx = context.WithValue(ctx, t.Merchant, merchant)                  // เพิ่มข้อมูล merchant ใน context
	c.SetUserContext(ctx)                                               // ตั้งค่า context ใหม่

	return nil // คืนค่า nil เมื่อสำเร็จ
}
