package config

import (
	"log"

	"github.com/gin-gonic/gin"
)

func InitApp(gin *gin.Engine) {
	db, err := InitPostgresDB()
	if err != nil {
		log.Fatal("failed connecting to database: ", err.Error())
	}

	container := NewContainer(db)
	InitRoute(gin, container)
}
