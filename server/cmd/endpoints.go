package main

import (
	"server/controller"

	"github.com/gin-gonic/gin"
)

// Sets all the endpoints for the Gin router
func SetEndpoints(r *gin.Engine, c *controller.Controller) {

	r.GET("/cases", func(ctx *gin.Context) {
		c.GetAllCasesForLawyer(ctx)
	})

}
