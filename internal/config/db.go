package config

import (
	"agro-mart/internal/constant"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgresDB() (*gorm.DB, error) {
	dsn := getDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN() string {
	dbHost := os.Getenv(constant.DB_HOST_POSTGRES)
	dbUser := os.Getenv(constant.DB_USER_POSTGRES)
	dbPassword := os.Getenv(constant.DB_PASSWORD_POSTGRES)
	dbName := os.Getenv(constant.DB_NAME_POSTGRES)
	dbPort := os.Getenv(constant.DB_PORT_POSTGRES)
	dbSslMode := os.Getenv(constant.DB_SSLMODE_POSTGRES)
	dbTimeZone := os.Getenv(constant.DB_TIMEZONE_POSTGRES)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
		dbSslMode,
		dbTimeZone,
	)

	return dsn
}
