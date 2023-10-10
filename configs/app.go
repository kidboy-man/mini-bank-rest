package configs

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

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

	JWTConfig
}

type JWTConfig struct {
	JWTSignatureKey     string
	JWTPublicKey        string
	JWTExpirationSecond time.Duration
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

	c.JWTSignatureKey = os.Getenv("JWT_SIGNATURE_KEY")
	if c.JWTSignatureKey == "" {
		err = errors.New("JWT_SIGNATURE_KEY is not set")
		return
	}

	c.JWTPublicKey = os.Getenv("JWT_PUBLIC_KEY")
	if c.JWTPublicKey == "" {
		err = errors.New("JWT_PUBLIC_KEY is not set")
		return
	}

	strJWTExpirationSecond := os.Getenv("JWT_EXPIRATION_SECOND")
	jwtExpirationSecond, _ := strconv.Atoi(strJWTExpirationSecond)
	if jwtExpirationSecond == 0 {
		jwtExpirationSecond = 3600
	}

	c.JWTExpirationSecond = time.Duration(jwtExpirationSecond) * time.Second
	return
}
