package controllers

import (
	"go-pudthai/internal/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	logger "go-pudthai/internal/infrastructure/logger"
)

type dbController struct {
	db *gorm.DB
}

func NewDatabaseController(db *gorm.DB) DatabaseController {
	return &dbController{
		db: db,
	}
}

type DatabaseController interface {
	AutoMigrate(c *fiber.Ctx) error
}

func (s dbController) AutoMigrate(c *fiber.Ctx) error {
	if err := s.db.AutoMigrate(
		&entities.Merchant{},
		&entities.User{},
		&entities.Role{},
		&entities.OauthAccessToken{},
		&entities.OauthRefreshToken{},
	); err != nil {
		logger.Log.Err(err).Msg("Migration failed")
	}

	return c.Status(fiber.StatusOK).
		JSON(fiber.Map{
			"message": "Migration success",
		})
}
