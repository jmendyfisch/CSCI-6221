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

// Sets all the endpoints for the Gin router
func SetEndpoints(r *gin.Engine, c *controller.Controller) {

	r.GET("/cases", func(ctx *gin.Context) {
		c.GetAllCasesForLawyer(ctx)
	})

	r.POST("/create_case", func(ctx *gin.Context) {
		c.CreateNewCase(ctx)
	})

}
