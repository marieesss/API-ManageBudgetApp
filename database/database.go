package database

import (
	"fmt"
	"log"
	"os"

	"github.com/marieesss/API-ManageBudgetApp/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Db DbInstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Europe/Paris",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to db \n", err)
		os.Exit(2)
	}

	log.Println("connected")

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{})

	Db = DbInstance{
		Db: db,
	}

}
