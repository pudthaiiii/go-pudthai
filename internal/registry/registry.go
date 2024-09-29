// Code by พิเชษฐ์ ขุนใจ (คุณผัดไท)
// source: internal/registry/registry.go

/*
Package registry ใช้สำหรับจัดการการลงทะเบียน dependencies ต่างๆ ที่จำเป็น
เช่น Controllers, Middleware และอื่นๆ โดยเป็นตัวเชื่อมระหว่างโครงสร้างพื้นฐาน
และส่วนต่างๆ ของแอปพลิเคชัน
*/
package registry

import (
	cc "go-pudthai/internal/adapter/console/controllers"
	cm "go-pudthai/internal/adapter/console/middleware"
	sm "go-pudthai/internal/adapter/shared/middleware"
	ac "go-pudthai/internal/adapter/v1/admin/controllers"
	ab "go-pudthai/internal/adapter/v1/backend/controllers"
	af "go-pudthai/internal/adapter/v1/frontend/controllers"
	"go-pudthai/internal/config"
	"go-pudthai/internal/events"
	"go-pudthai/internal/infrastructure/cache"
	"go-pudthai/internal/infrastructure/datastore"
	"go-pudthai/internal/infrastructure/recaptcha"

	"gorm.io/gorm"
)

// Registry interface คือส่วนที่กำหนดฟังก์ชันสำหรับสร้าง Controller และ Middleware ต่างๆ
type Registry interface {
	NewAdminController() ac.AdminController
	NewBackendController() ab.BackendController
	NewFrontendController() af.FrontendController
	NewConsoleController() cc.ConsoleController
	NewConsoleMiddleware() cm.Middleware
	NewSharedMiddleware() sm.Middleware
}

// registry struct คือโครงสร้างที่ใช้ในการเก็บ dependencies ต่างๆ ที่จำเป็น
type registry struct {
	db           *gorm.DB
	s3           *datastore.S3Datastore
	cfg          *config.Config
	recaptcha    *recaptcha.RecaptchaProvider
	cacheManager *cache.CacheManager
	listener     events.EventListener
}

// NewRegistry เป็นฟังก์ชันที่ใช้สร้าง registry ใหม่พร้อม dependencies ที่จำเป็น
func NewRegistry(
	db *gorm.DB,
	s3 *datastore.S3Datastore,
	cfg *config.Config,
	recaptcha *recaptcha.RecaptchaProvider,
	cacheManager *cache.CacheManager,
	listener events.EventListener,
) Registry {
	return &registry{
		db:           db,
		s3:           s3,
		cfg:          cfg,
		recaptcha:    recaptcha,
		cacheManager: cacheManager,
		listener:     listener,
	}
}

// NewConsoleController เป็นฟังก์ชันที่ใช้สร้าง ConsoleController
func (r *registry) NewConsoleController() cc.ConsoleController {
	return cc.ConsoleController{
		DatabaseController: r.NewConsoleDatabaseController(),
	}
}

// NewAdminController เป็นฟังก์ชันที่ใช้สร้าง AdminController
func (r *registry) NewAdminController() ac.AdminController {
	return ac.AdminController{
		AuthController:  r.NewAdminAuthController(),
		UsersController: r.NewAdminUsersController(),
	}
}

// NewBackendController เป็นฟังก์ชันที่ใช้สร้าง BackendController
func (r *registry) NewBackendController() ab.BackendController {
	return ab.BackendController{
		AuthController: r.NewBackendAuthController(),
	}
}

// NewFrontendController เป็นฟังก์ชันที่ใช้สร้าง FrontendController
func (r *registry) NewFrontendController() af.FrontendController {
	return af.FrontendController{
		AuthController: r.NewFrontendAuthController(),
	}
}

// NewConsoleMiddleware เป็นฟังก์ชันที่ใช้สร้าง middleware สำหรับ Console
func (r *registry) NewConsoleMiddleware() cm.Middleware {
	return cm.NewConsoleMiddleware(r.cfg)
}

// NewSharedMiddleware เป็นฟังก์ชันที่ใช้สร้าง shared middleware
func (r *registry) NewSharedMiddleware() sm.Middleware {
	return sm.NewSharedMiddleware(r.cfg, r.cacheManager, r.db, r.recaptcha)
}
