package config

import (
	"fmt"
	"log"
	"os"
	"time"

	models "github.com/tabed23/social-media-api/graph/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type GormDB struct {
	db *gorm.DB
}

// NewGormDB creates a new GormDB instance
func ConnectDB() (*GormDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("dbhost"), os.Getenv("dbport"), os.Getenv("dbuser"), os.Getenv("dbpass"), os.Getenv("dbname"), "disable",
	)
	DB, err := gorm.Open(postgres.Open(dsn), initConfig())
	if err != nil {
		return nil, err
	}

	return &GormDB{db: DB}, nil
}

// InitConfig Initialize Config
func initConfig() *gorm.Config {
	return &gorm.Config{
		Logger:         initLog(),
		NamingStrategy: initNamingStrategy(),
	}
}

// InitLog Connection Log Configuration
func initLog() logger.Interface {

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful:      true,
			LogLevel:      logger.Info,
			SlowThreshold: time.Second,
		})
	return newLogger
}

// InitNamingStrategy Init NamingStrategy
func initNamingStrategy() *schema.NamingStrategy {
	return &schema.NamingStrategy{
		SingularTable: false,
		TablePrefix:   "",
	}
}

// Schema Migration
func (p *GormDB) StartMigration() error {
	return p.db.AutoMigrate(&models.User{}, &models.Post{}, &models.Like{}, &models.Comment{})
}

// Get the database
func (p *GormDB) GetDB() *gorm.DB {
	return p.db
}
