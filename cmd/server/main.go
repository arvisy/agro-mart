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
	r := gin.Default()
	config.InitServer(r)

	address := os.Getenv(constant.GIN_ADDRESS)
	port := os.Getenv(constant.GIN_PORT)

	err := r.Run(fmt.Sprintf("%s:%s", address, port))
	if err != nil {
		log.Fatal("failed to running program: " + err.Error())
	}
}
