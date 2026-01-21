package main

import (
	"agro-mart/internal/config"
	"agro-mart/internal/constant"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	config.InitApp(app)

	address := os.Getenv(constant.GIN_ADDRESS)
	port := os.Getenv(constant.GIN_PORT)

	err := app.Run(fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal("failed to running program: " + err.Error())
	}
}
