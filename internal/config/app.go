// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/config/app.go

/*
Package config ใช้สำหรับการตั้งค่าของแอปพลิเคชัน

โดยจะมีการโหลดค่าต่างๆ จากไฟล์ .env และตั้งค่าต่างๆ ของแอปพลิเคชัน
*/
package config

import (
	"go-pudthai/internal/utils"
)

// กำหนดค่าเริ่มต้นของ bodyLimit และ bufferSize โดยใช้ฟังก์ชัน CalFileSize จาก utils
var (
	bodyLimit, _  = utils.CalFileSize("10mb") // กำหนดขนาดของ body request เป็น 10mb
	bufferSize, _ = utils.CalFileSize("10mb") // กำหนดขนาดของ buffer เป็น 10mb
)

// loadAppConfig โหลดค่าการตั้งค่าต่างๆ ของแอปพลิเคชัน
func (c *Config) loadAppConfig() {
	// กำหนดค่าต่างๆ ลงในแผนที่ cfg
	c.cfg = map[string]interface{}{
		// logger config การตั้งค่าสำหรับการบันทึก log
		"LoggerConfig": map[string]interface{}{
			"FileLog": "storage/logs.log",                        // ไฟล์ที่จะเก็บ log
			"Enabled": getEnv("AWS_CLOUDWATCH_ENABLED", "false"), // เปิดใช้ CloudWatch log หรือไม่
		},

		// fiber config การตั้งค่าสำหรับ Fiber framework
		"FiberConfig": map[string]interface{}{
			"BodyLimit":       int(bodyLimit),         // ขนาดสูงสุดของ body ที่รับได้
			"ReadBufferSize":  int(bufferSize),        // ขนาด buffer สำหรับการอ่าน request
			"WriteBufferSize": int(bufferSize),        // ขนาด buffer สำหรับการเขียน response
			"Port":            getEnv("PORT", "3000"), // พอร์ตที่แอปพลิเคชันจะรัน
		},

		// Postgresql config การตั้งค่าสำหรับการเชื่อมต่อกับ PostgreSQL
		"Postgresql": map[string]interface{}{
			"Host":     getEnv("DB_HOST"),     // ที่อยู่ของฐานข้อมูล
			"Port":     getEnv("DB_PORT"),     // พอร์ตของฐานข้อมูล
			"User":     getEnv("DB_USERNAME"), // ชื่อผู้ใช้สำหรับการเชื่อมต่อฐานข้อมูล
			"Password": getEnv("DB_PASSWORD"), // รหัสผ่านสำหรับการเชื่อมต่อฐานข้อมูล
			"DBName":   getEnv("DB_DATABASE"), // ชื่อฐานข้อมูล
			"SSL":      getEnv("DB_SSL"),      // การตั้งค่า SSL สำหรับการเชื่อมต่อ
		},

		// Redis config การตั้งค่าสำหรับการเชื่อมต่อกับ Redis
		"Redis": map[string]interface{}{
			"ClusterEnabled": utils.StringToBool(getEnv("REDIS_CLUSTER_ENABLED")), // เปิดใช้ Redis Cluster หรือไม่
			"Host":           getEnv("REDIS_HOST"),                                // ที่อยู่ของ Redis
			"Port":           getEnv("REDIS_PORT"),                                // พอร์ตของ Redis
			"Password":       getEnv("REDIS_PASSWORD"),                            // รหัสผ่านของ Redis
			"DB":             getEnv("REDIS_DB"),                                  // เลขที่ของฐานข้อมูลใน Redis
			"ClusterNodes":   getEnv("REDIS_CLUSTER_NODES"),                       // รายการ node ของ Redis Cluster
		},

		// S3 config การตั้งค่าสำหรับการเชื่อมต่อกับ Amazon S3
		"S3": map[string]interface{}{
			"AccessKeyID":     getEnv("AWS_S3_ACCESS_KEY_ID"),     // Access Key ของ AWS S3
			"SecretAccessKey": getEnv("AWS_S3_SECRET_ACCESS_KEY"), // Secret Access Key ของ AWS S3
			"Region":          getEnv("AWS_S3_REGION"),            // ภูมิภาคของ S3 bucket
			"Bucket":          getEnv("AWS_S3_BUCKET"),            // ชื่อ S3 bucket
		},

		// CloudWatch config การตั้งค่าสำหรับ Amazon CloudWatch
		"CloudWatch": map[string]interface{}{
			"AccessKeyID":     getEnv("AWS_CLOUDWATCH_ACCESS_KEY_ID"),     // Access Key สำหรับ CloudWatch
			"SecretAccessKey": getEnv("AWS_CLOUDWATCH_SECRET_ACCESS_KEY"), // Secret Access Key สำหรับ CloudWatch
			"Region":          getEnv("AWS_CLOUDWATCH_REGION"),            // ภูมิภาคของ CloudWatch
			"LogGroupName":    getEnv("AWS_CLOUDWATCH_LOG_GROUP_NAME"),    // ชื่อกลุ่ม log ใน CloudWatch
			"LogStreamName":   getEnv("AWS_CLOUDWATCH_LOG_STREAM_NAME"),   // ชื่อ stream log ใน CloudWatch
		},

		// GoogleRecaptcha config การตั้งค่าสำหรับ Google Recaptcha
		"GoogleRecaptcha": map[string]interface{}{
			"RecaptchaSecretKey": getEnv("GOOGLE_RECAPTCHA_SECRET_KEY"), // กุญแจลับสำหรับ Google Recaptcha
			"RecaptchaEnabled":   getEnv("GOOGLE_RECAPTCHA_ENABLED"),    // เปิดใช้ Google Recaptcha หรือไม่
		},

		// MailServer config การตั้งค่าสำหรับการส่งอีเมล
		"MailServer": map[string]interface{}{
			"SmtpServer":  getEnv("MAIL_HOST"),         // ที่อยู่ของเซิร์ฟเวอร์ SMTP
			"SmtpPort":    getEnv("MAIL_PORT"),         // พอร์ตของเซิร์ฟเวอร์ SMTP
			"Username":    getEnv("MAIL_USERNAME"),     // ชื่อผู้ใช้สำหรับการเชื่อมต่อ SMTP
			"Password":    getEnv("MAIL_PASSWORD"),     // รหัสผ่านสำหรับการเชื่อมต่อ SMTP
			"FromAddress": getEnv("MAIL_FROM_ADDRESS"), // ที่อยู่อีเมลผู้ส่ง
			"Encryption":  getEnv("MAIL_ENCRYPTION"),   // ประเภทของการเข้ารหัส (เช่น SSL หรือ TLS)
		},

		// JWT config การตั้งค่าสำหรับการจัดการ JWT
		"JWT": map[string]interface{}{
			"JwtSecret":               getEnv("JWT_SECRET"),                        // กุญแจลับสำหรับ JWT
			"JwtSecretAdmin":          getEnv("JWT_SECRET_ADMIN"),                  // กุญแจลับสำหรับ Admin JWT
			"JwtSecretBackend":        getEnv("JWT_SECRET_BACKEND"),                // กุญแจลับสำหรับ Backend JWT
			"JwtAccessExpiresInHour":  getEnv("JWT_ACCESS_TOKEN_EXPIRES_IN_HOUR"),  // ระยะเวลาหมดอายุของ Access Token (ชั่วโมง)
			"JwtRefreshExpiresInHour": getEnv("JWT_REFRESH_TOKEN_EXPIRES_IN_HOUR"), // ระยะเวลาหมดอายุของ Refresh Token (ชั่วโมง)
		},

		// Cookie config การตั้งค่าการใช้ Cookie
		"Cookie": map[string]interface{}{
			"Name":   getEnv("COOKIE_NAME", "console"),   // ชื่อ cookie
			"Secret": getEnv("COOKIE_SECRET", "pudthai"), // กุญแจลับสำหรับการเข้ารหัส cookie
		},
	}
}
