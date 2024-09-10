package pkg

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pudthaiiii/golang-cms/src/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	dbConfig "github.com/pudthaiiii/golang-cms/src/config/database"
)

type sqlLogger struct {
	logger.Interface
}

type PgDatastore struct {
	DB *gorm.DB
}

func ConnectPgSql() PgDatastore {
	pgConfig := dbConfig.GetPGConfig()

	var pg PgDatastore
	pg.connectDB(pgConfig)

	return pg
}

func (p *PgDatastore) connectDB(config types.PGConfig) {
	pgSqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBDatabase)

	db, err := gorm.Open(postgres.Open(pgSqlInfo), &gorm.Config{
		Logger: &sqlLogger{},
	})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(&entities.UsageItem{}, &entities.Product{}, &entities.Category{})

	p.DB = db

	sqlDB, err := p.DB.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(0)

	log.Printf("%s, DB: %s", "Successfully connected to PostgreSQL", config.DBDatabase)
}

func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()

	log.Printf("SQL: %s\n==============================================================================\n", sql)
}
