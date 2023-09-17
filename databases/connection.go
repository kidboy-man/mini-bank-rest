package databases

import (
	"fmt"
	"log"

	"github.com/kidboy-man/mini-bank-rest/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (err error) {
	dbUri := fmt.Sprintf(
		"host=%s port=%s, user=%s dbname=%s sslmode=disable password=%s",
		configs.AppConfig.DBHost,
		configs.AppConfig.DBPort,
		configs.AppConfig.DBUser,
		configs.AppConfig.DBName,
		configs.AppConfig.DBPassword,
	)

	log.Println(dbUri)
	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	configs.AppConfig.DBConnection = db
	return
}

// returns a handle to the DB object
func GetDB() *gorm.DB {
	return configs.AppConfig.DBConnection
}
