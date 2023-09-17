package configs

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	AppPort string

	DBHost       string
	DBPort       string
	DBName       string
	DBUser       string
	DBPassword   string
	DBConnection *gorm.DB
}

var AppConfig Config

func (c *Config) Setup() (err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return
	}

	c.AppPort = os.Getenv("APP_PORT")
	if c.AppPort == "" {
		err = errors.New("APP_PORT is not set")
		return
	}

	c.DBHost = os.Getenv("DB_HOST")
	if c.DBHost == "" {
		err = errors.New("DB_HOST is not set")
		return
	}

	c.DBPort = os.Getenv("DB_PORT")
	if c.DBPort == "" {
		err = errors.New("DB_PORT is not set")
		return
	}

	c.DBName = os.Getenv("DB_NAME")
	if c.DBName == "" {
		err = errors.New("DB_NAME is not set")
		return
	}

	c.DBUser = os.Getenv("DB_USER")
	if c.DBUser == "" {
		err = errors.New("DB_USER is not set")
		return
	}

	c.DBPassword = os.Getenv("DB_PASS")
	if c.DBPassword == "" {
		err = errors.New("DB_PASS is not set")
		return
	}

	return
}
