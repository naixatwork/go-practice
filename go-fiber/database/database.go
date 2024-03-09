package database

import (
	"fmt"
	"github.com/naixatwork/trivia/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	//defer (func() {
	//	sqlDb, _ := db.DB()
	//	sqlDb.Close()
	//})()

	if err != nil {
		log.Fatal("Failed to connect t o database. \n", err)
	}

	log.Println("connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Printf("running migrations")
	if err := db.AutoMigrate(&models.Fact{}); err != nil {
		log.Printf("error on auto migrating models.Fact %s", err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
