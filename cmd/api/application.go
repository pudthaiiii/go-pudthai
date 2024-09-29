// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: cmd/api/application.go

/*
Package api ใช้สำหรับการสร้างแอปพลิเคชัน API

โดยจะมีการตั้งค่าและเริ่มต้นการทำงานของแอปพลิเคชัน
*/
package api

import (
	"go-pudthai/internal/config"
	"go-pudthai/internal/events"
	"go-pudthai/internal/infrastructure/cache"
	"go-pudthai/internal/infrastructure/datastore"
	"go-pudthai/internal/infrastructure/logger"
	"go-pudthai/internal/infrastructure/mailer"
	"go-pudthai/internal/infrastructure/recaptcha"
	"go-pudthai/internal/registry"
	"go-pudthai/internal/router"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type application struct {
	fiber *fiber.App
	cfg   *config.Config
}

// NewApiApplication ทำการสร้างและตั้งค่าแอปพลิเคชัน API ใหม่
func NewApiApplication() *application {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("ไม่สามารถโหลดการตั้งค่าได้: %v", err)
	}

	cfg.Initialize()
	app := &application{
		cfg:   cfg,
		fiber: newFiberApp(cfg),
	}

	return app
}

// newFiberApp ตั้งค่าและส่งกลับ instance ของ Fiber application ใหม่
func newFiberApp(cfg *config.Config) *fiber.App {
	fiberConfig := fiber.Config{
		BodyLimit:       cfg.Get("FiberConfig")["BodyLimit"].(int),
		ReadBufferSize:  cfg.Get("FiberConfig")["ReadBufferSize"].(int),
		WriteBufferSize: cfg.Get("FiberConfig")["WriteBufferSize"].(int),
		ReadTimeout:     10 * time.Second,
		WriteTimeout:    10 * time.Second,
		ErrorHandler:    errorHandler,
	}

	return fiber.New(fiberConfig)
}

// Fiber ส่งกลับ instance ของ Fiber application
func (app *application) Fiber() *fiber.App {
	return app.fiber
}

// Config ส่งกลับการตั้งค่าของแอปพลิเคชัน
func (app *application) Config() *config.Config {
	return app.cfg
}

// Boot เริ่มต้นการทำงานของแอปพลิเคชันและฟังก์ชันต่างๆ
func (app *application) Boot() {
	app.setupLogger()
	app.setupRouter()
	app.startListening()
}

// setupRouter ตั้งค่า router สำหรับแอปพลิเคชัน
func (app *application) setupRouter() {
	reg := app.initializeRegistry()
	router.InitializeRoute(app.fiber, reg)
}

// setupLogger ตั้งค่า logger สำหรับแอปพลิเคชัน
func (app *application) setupLogger() {
	logger.NewInitializeLogger(app.cfg)
}

// initializeRegistry ตั้งค่าระบบ registry ของแอปพลิเคชัน
func (app *application) initializeRegistry() registry.Registry {
	s3 := datastore.NewS3Datastore(app.cfg)
	db := datastore.NewPgDatastore(app.cfg)
	redis := datastore.NewRedisDatastore(app.cfg)
	recaptchaProvider := recaptcha.NewRecaptchaProvider(app.cfg)

	cacheManager := cache.NewCacheManager(redis.Client, 5*time.Minute)
	listener := app.createEventListener(db, cacheManager)

	return registry.NewRegistry(db, s3, app.cfg, recaptchaProvider, cacheManager, listener)
}

// createEventListener ตั้งค่า event listener สำหรับแอปพลิเคชัน
func (app *application) createEventListener(db *gorm.DB, cacheManager *cache.CacheManager) events.EventListener {
	mailerInstance := mailer.NewMailer(app.cfg)
	listener := events.NewEventListener(mailerInstance, db, cacheManager)

	go listener.Listen() // รัน listener ใน goroutine

	return listener
}

// startListening เริ่มฟังการร้องขอบนพอร์ตที่กำหนด
func (app *application) startListening() {
	port := app.cfg.Get("FiberConfig")["Port"].(string)
	if err := app.fiber.Listen(":" + port); err != nil {
		log.Fatalf("เซิร์ฟเวอร์ไม่สามารถเริ่มต้นได้: %v", err)
	}
}

// DeferClose ทำการทำความสะอาดก่อนปิดแอปพลิเคชัน
func (app *application) DeferClose() {
	defer logger.CloseLogger()
}
