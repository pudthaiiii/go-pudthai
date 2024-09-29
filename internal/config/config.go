// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/config/config.go

/*
Package config ใช้สำหรับการตั้งค่าของแอปพลิเคชัน

โดยจะมีการโหลดค่าต่างๆ จากไฟล์ .env และตั้งค่าต่างๆ ของแอปพลิเคชัน
*/
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config โครงสร้างสำหรับจัดเก็บ configuration ของแอปพลิเคชัน
type Config struct {
	cfg map[string]interface{} // แผนที่สำหรับเก็บค่า configuration
}

// NewConfig โหลดไฟล์ .env และสร้าง instance ของ Config
// ถ้าไม่สามารถโหลดไฟล์ได้จะส่ง error กลับมา
func NewConfig() (*Config, error) {
	// โหลดค่า environment จากไฟล์ .env
	if err := godotenv.Load(); err != nil {
		return nil, err // ถ้าโหลดไฟล์ .env ไม่สำเร็จจะ return error
	}

	// คืนค่า instance ของ Config
	return &Config{
		cfg: make(map[string]interface{}), // สร้างแผนที่เก็บค่า config
	}, nil
}

// Get คืนค่า configuration ที่เก็บในแผนที่ตาม key ที่ระบุ
// ถ้าไม่พบ key ที่ต้องการ จะคืนค่า nil
func (c *Config) Get(key string) map[string]interface{} {
	// ตรวจสอบว่ามี key นี้อยู่ในแผนที่หรือไม่
	if _, exists := c.cfg[key]; !exists {
		return nil // คืนค่า nil ถ้าไม่พบ key
	}

	// คืนค่าแผนที่ที่เก็บอยู่ใน key นั้น
	return c.cfg[key].(map[string]interface{})
}

// Add เพิ่มค่า configuration ลงในแผนที่โดยใช้ key และ value ที่ระบุ
func (c *Config) Add(key string, value map[string]interface{}) {
	c.cfg[key] = value // เพิ่มคู่ key-value ลงในแผนที่ cfg
}

// Initialize เรียกฟังก์ชันเพื่อโหลด configuration พื้นฐาน
func (c *Config) Initialize() {
	c.loadAppConfig() // เรียกฟังก์ชันเพื่อโหลด config ของแอปพลิเคชัน
}

// getEnv อ่านค่า environment จาก key ที่ระบุ
// ถ้าไม่พบค่าใน environment จะคืนค่าจาก defaultValue ที่กำหนด
func getEnv(key string, defaultValue ...string) string {
	var defaultVal string

	// ถ้ามี defaultValue ส่งเข้ามา จะกำหนดให้ defaultVal
	if len(defaultValue) > 0 {
		defaultVal = defaultValue[0]
	}

	// ถ้าพบ key ใน environment จะคืนค่าจาก environment
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	// คืนค่า defaultValue ถ้าไม่พบค่าใน environment
	return defaultVal
}
