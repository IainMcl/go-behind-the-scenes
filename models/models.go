package models

import (
	"fmt"

	"github.com/IainMcl/go-behind-the-scenes/internal/logging"
	"github.com/IainMcl/go-behind-the-scenes/internal/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Setup() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		settings.DatabaseSettings.Host,
		settings.DatabaseSettings.User,
		settings.DatabaseSettings.Password,
		settings.DatabaseSettings.DbName,
		settings.DatabaseSettings.Port,
		settings.DatabaseSettings.SSLMode,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Fatal("models.Setup err: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logging.Fatal("models.Setup err: %v", err)
	}
	sqlDB.SetMaxIdleConns(settings.DatabaseSettings.MaxIdleConns)
	sqlDB.SetMaxOpenConns(settings.DatabaseSettings.MaxOpenConns)

	if err := db.AutoMigrate(&Role{}); err != nil {
		logging.Fatal("models.Setup err: %v", err)
	}

	adminRole := Role{Name: "admin", RateLimit: 1000}
	userRole := Role{Name: "user", RateLimit: 100}

	db.Create(&userRole)
	db.Create(&adminRole)

	if err := db.AutoMigrate(&User{}); err != nil {
		logging.Fatal("models.Setup err: %v", err)
	}
}
