package main

import (
	"crypto/sha256"
	"encoding/hex"
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

// gorilla session middleware
/*  This was me trying to implement gorilla sessions. I couldn't get it to work.
func GorillaSessionMiddleware(store *sessions.CookieStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Load session
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			log.Printf("Error loading session: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Make the session available in the context
		c.Set("session", session)

		c.Next()

		// Save the session
		err = sessions.Save(c.Request, c.Writer)
		if err != nil {
			log.Printf("Error saving session: %v", err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
}
*/

// Sets all the endpoints for the Gin router
func SetEndpoints(r *gin.Engine, c *controller.Controller) {

	r.LoadHTMLFiles("templates/index.html", "templates/intake.html", "templates/lawyer-login.html", "templates/display-cases.html", "templates/new-account.html")

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

	r.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/lawyer-login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lawyer-login.html", gin.H{})
	})

	r.GET("/new-account", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "new-account.html", gin.H{})
	})

	r.GET("/display-cases", func(ctx *gin.Context) {

		//login implementing custom-built security

		// Check if lawyerID cookie exists.
		lawyerID, err1 := ctx.Cookie("lawyer_id")
		securitystring, err2 := ctx.Cookie("securitystring")
		timestamp, err3 := ctx.Cookie("securitytimestamp")

		if err1 != nil || err2 != nil || err3 != nil {
			// If the cookie doesn't exist, redirect to the login page
			ctx.Redirect(http.StatusFound, "/lawyer-login")
			fmt.Println("Missing one or more required cookies.")
			return
		}

		data := config.CookieKey + timestamp + lawyerID
		hash := sha256.Sum256([]byte(data))

		if hex.EncodeToString(hash[:]) == securitystring {
			// If everything matches, proceed to display cases.
			ctx.HTML(http.StatusOK, "display-cases.html", gin.H{"lawyer_id": lawyerID})
		} else {
			ctx.Redirect(http.StatusFound, "/lawyer-login")
			fmt.Println("Cookie security data invalid.")
			return
		}

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

	r.POST("/lawyer_login", func(ctx *gin.Context) {
		c.AuthenticateLawyer(ctx)
	})

	r.POST("/create_lawyer_account", func(ctx *gin.Context) {
		c.CreateNewLawyer(ctx)
	})

}
