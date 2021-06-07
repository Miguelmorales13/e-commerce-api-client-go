package providers

import (
	"crud/entities"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	Database *gorm.DB
)

func InitDbPostgres() {
	dns := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSMODE"))
	var err error = nil
	anyLog := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)

	Database, err = gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: anyLog,
	})
	Database.Session(&gorm.Session{Logger: anyLog})
	if err != nil {
		log.Fatal("[ERROR_DATABASE_CONNECTION][", err.Error(), "]")
	}
	err = Database.AutoMigrate(&entities.User{}, &entities.UsersAddresses{}, &entities.Product{}, &entities.ProductImages{}, &entities.Category{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
