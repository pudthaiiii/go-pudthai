package controllers

import (
	"go-ibooking/internal/entities"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	logger "go-ibooking/internal/infrastructure/logger"
)

type featuresController struct {
	db *gorm.DB
}

func NewFeaturesController(db *gorm.DB) FeaturesController {
	return &featuresController{
		db: db,
	}
}

type FeaturesController interface {
	AutoMigrate(c *fiber.Ctx) error
}

func (s featuresController) AutoMigrate(c *fiber.Ctx) error {
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
