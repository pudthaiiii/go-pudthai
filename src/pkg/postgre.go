// package pkg

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"github.com/pudthaiiii/golang-cms/src/types"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"

// 	dbConfig "github.com/pudthaiiii/golang-cms/src/config/database"
// )

// type Datastore interface {
// 	connectDB(config types.PGConfig)
// }

// type PgDatastore struct {
// 	DB *gorm.DB
// }

// func ConnectPgSql() PgDatastore {
// 	pgConfig := dbConfig.GetPGConfig()

// 	var pg PgDatastore
// 	pg.connectDB(pgConfig)

// 	return pg
// }

// func (p *PgDatastore) connectDB(config types.PGConfig) {
// 	pgSqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
// 		config.Host, config.Port, config.User, config.Password, config.DBDatabase)

// 	db, err := gorm.Open(postgres.Open(pgSqlInfo), &gorm.Config{
// 		Logger: &sqlLogger{},
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	// db.AutoMigrate(&entities.UsageItem{}, &entities.Product{}, &entities.Category{})

// 	p.DB = db

// 	sqlDB, err := p.DB.DB()
// 	if err != nil {
// 		panic(err)
// 	}

// 	sqlDB.SetConnMaxLifetime(0)
// 	sqlDB.SetMaxIdleConns(0)

// 	log.Printf("%s, DB: %s", "Successfully connected to PostgreSQL", config.DBDatabase)
// }

// type sqlLogger struct {
// 	logger.Interface
// }

// func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
// 	sql, _ := fc()

// 	log.Printf("SQL: %s\n==============================================================================\n", sql)
// }

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
)

type PgDatastore struct {
	DB *gorm.DB
}

func NewPgDatastore(config types.PGConfig) *PgDatastore {
	db := connectPgSql(config)
	return &PgDatastore{DB: db}
}

func (pg *PgDatastore) Close() error {
	sqlDB, err := pg.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to retrieve database connection: %w", err)
	}
	return sqlDB.Close()
}

func connectPgSql(config types.PGConfig) *gorm.DB {
	pgSqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBDatabase)

	db, err := gorm.Open(postgres.Open(pgSqlInfo), &gorm.Config{
		Logger: &sqlLogger{},
	})
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve database connection: %v", err)
	}

	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(0)

	log.Printf("Successfully connected to PostgreSQL, DB: %s", config.DBDatabase)
	return db
}

type sqlLogger struct {
	logger.Interface
}

func (l sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, _ := fc()
	log.Printf("SQL: %s\n==============================================================================\n", sql)
}
