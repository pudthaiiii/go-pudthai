package datastore

import (
	"context"
	"fmt"
	"go-ibooking/internal/config"
	log "go-ibooking/internal/infrastructure/logger"
	"net/url"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPgDatastore(cfg *config.Config) *gorm.DB {
	config := cfg.Get("Postgresql")

	host := config["Host"].(string)
	port := config["Port"].(string)
	user := config["User"].(string)
	password := config["Password"].(string)
	dbName := config["DBName"].(string)
	ssl := config["SSL"].(string)

	u := url.URL{
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%s", host, port),
		User:   url.UserPassword(user, password),
		Path:   dbName,
	}

	query := u.Query()
	query.Set("sslmode", ssl)
	u.RawQuery = query.Encode()

	dsn := u.String()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: &sqlLogger{},
	})

	if err != nil {
		log.Log.Err(err).Msg("failed to connect to PostgreSQL")
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Log.Err(err).Msg("failed to get SQL database instance")
		return nil
	}

	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)

	log.Write.Info().Msg("Successfully connected to PostgreSQL")
	return db
}

type sqlLogger struct {
	logger.Interface
}

func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if os.Getenv("DB_LOGGING") == "false" {
		return
	}

	sql, _ := fc()
	log.Write.Printf("SQL: %s\n==============================================================================\n", sql)
}
