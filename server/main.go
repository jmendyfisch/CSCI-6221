package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/controller"
	"server/database"
	"server/service"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("heello")
	database.Init()

	router := gin.Default()

	router.Static("/templates", "./templates")

	serv := service.New()
	cont := controller.New(serv)

	SetEndpoints(router, &cont)

	router.Run(config.RunPort)

}

// Sets all the endpoints for the Gin router
func SetEndpoints(r *gin.Engine, c *controller.Controller) {

	r.LoadHTMLFiles("templates/index.html", "templates/intake.html")

	r.POST("/save-audio", func(ctx *gin.Context) {
		file, err := ctx.FormFile("audio")
		if err != nil {
			ctx.String(http.StatusBadRequest, "Get form err: %s", err.Error())
			return
		}

		// Generate a unique filename, e.g., using a timestamp
		filename := fmt.Sprintf("recording_%v.wav", time.Now().UnixNano())
		path := fmt.Sprintf("/tempaudio/%s", filename)

		// Save the uploaded file
		if err := ctx.SaveUploadedFile(file, path); err != nil {
			ctx.String(http.StatusInternalServerError, "Save uploaded file err: %s", err.Error())
		} else {
			ctx.String(http.StatusOK, "File uploaded successfully: %s", filename)
		}
	})

	r.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/intake", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "intake.html", gin.H{})
	})

	r.GET("/cases", func(ctx *gin.Context) {
		c.GetAllCasesForLawyer(ctx)
	})

	r.POST("/create_case", func(ctx *gin.Context) {
		c.CreateNewCase(ctx)
	})

}
