package main

import (
	"fmt"
	"net/http"
	"server/config"
	"server/controller"
	"server/database"
	"server/service"
	"strings"
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
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		audioType := ctx.PostForm("type")
		if audioType == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No audio type provided"})
			return
		}
		audioType = strings.TrimPrefix(audioType, "audio/")

		case_id := ctx.PostForm("case_id")
		if case_id == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No case id provided"})
			return
		}

		filename := fmt.Sprintf("rec_case_%s_%v.%s", case_id, time.Now().UnixNano(), audioType)
		path := fmt.Sprintf("/tempaudio/%s", filename)

		if err := ctx.SaveUploadedFile(file, path); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "Audio saved successfully"})
	})

	r.GET("", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	//intake file accepts a case_id as a parameter
	r.GET("/intake/:case_id", func(ctx *gin.Context) {
		caseID := ctx.Param("case_id")
		if caseID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No case id provided"})
			return
		}

		// Pass caseID to the template
		ctx.HTML(http.StatusOK, "intake.html", gin.H{"case_id": caseID})
	})

	r.GET("/cases", func(ctx *gin.Context) {
		c.GetAllCasesForLawyer(ctx)
	})

	r.POST("/create_case", func(ctx *gin.Context) {
		c.CreateNewCase(ctx)
	})

}
