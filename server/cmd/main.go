package main

import (
	"fmt"
	"server/config"
	"server/controller"
	"server/database"
	"server/service"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("heello")
	database.Init()

	router := gin.Default()
	serv := service.New()
	cont := controller.New(serv)

	SetEndpoints(router, &cont)

	router.Run(config.RunPort)

}
