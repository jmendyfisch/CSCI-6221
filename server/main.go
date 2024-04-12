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

	"github.com/gin-gonic/gin"
)

func main() {
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

	r.LoadHTMLFiles("templates/index.html", "templates/intake.html", "templates/lawyer-login.html", "templates/display-cases.html", "templates/new-account.html", "templates/case-details.html", "templates/meeting-details.html")

	r.POST("/save-audio", func(ctx *gin.Context) {
		c.ProcessInterview(ctx)
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

	r.GET("/case-details", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "case-details.html", gin.H{})
	})

	r.GET("/get-case-details", func(ctx *gin.Context) {
		c.GetCaseDetails(ctx)
	})

	r.GET("/meeting-details", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "meeting-details.html", gin.H{})
	})

	r.GET("/get-all-meetings", func(ctx *gin.Context) {
		c.GetAllMeetings(ctx)
	})

	r.GET("/get-meetings-details", func(ctx *gin.Context) {
		c.GetMeetingDetails(ctx)
	})

	r.GET("/display-cases", func(ctx *gin.Context) {

		//Login implementing simple custom-built security. This passes three cookies.
		//The security string is based on the lawyerID, a secret (CookieKey), and a timestamp.
		//Without knowing the secret word or stealing the cookies, an attacker wouldn't
		//be able to guess the security string.

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

	r.GET("/log-out", func(ctx *gin.Context) {
		ctx.SetCookie("lawyer_id", "", -1, "/", "", false, true)
		ctx.SetCookie("securitystring", "", -1, "/", "", false, true)
		ctx.SetCookie("securitytimestamp", "", -1, "/", "", false, true)

		// Redirect the user to the login page or home page after logging out
		ctx.Redirect(http.StatusFound, "/index")
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
