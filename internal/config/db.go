package config

import (
	"agro-mart/internal/constant"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

func NewPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     os.Getenv(constant.DB_HOST_POSTGRES),
		User:     os.Getenv(constant.DB_USER_POSTGRES),
		Password: os.Getenv(constant.DB_PASSWORD_POSTGRES),
		Name:     os.Getenv(constant.DB_NAME_POSTGRES),
		Port:     os.Getenv(constant.DB_PORT_POSTGRES),
	}
}

func (c PostgresConfig) PostgresDSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
	)
}

func (c PostgresConfig) MigrationDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
	)
}

func InitPostgresDB() (*gorm.DB, error) {
	cfg := NewPostgresConfig()
	dsn := cfg.PostgresDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
