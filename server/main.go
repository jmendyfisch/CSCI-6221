// Main file, application server runs here.
// Directs all traffic to the appropriate controller functions.
// Original for project, based on controller-service-repository architecture (ChatGPT assisted with development)

package main

import (
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

	r.GET("/case-details/:case_id", func(ctx *gin.Context) {
		caseID := ctx.Param("case_id")
		if caseID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "No case id provided"})
			return
		}

		// Pass caseID to the CheckLogin
		c.CheckLogin(ctx, "case-details.html", "/display-cases", caseID, "")

	})

	r.GET("/get-case-details", func(ctx *gin.Context) {
		c.GetCaseDetails(ctx)
	})

	r.GET("/meeting-details/:case_id/:meeting_id", func(ctx *gin.Context) {
		caseID := ctx.Param("case_id")
		meetingID := ctx.Param("meeting_id")

		if caseID == "" || meetingID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing case or meeting id"})
			return
		}

		// Pass caseID and meeting id to the CheckLogin
		c.CheckLogin(ctx, "meeting-details.html", "/display-cases", caseID, meetingID)

	})

	r.GET("/get-all-meetings", func(ctx *gin.Context) {
		c.GetAllMeetings(ctx)
	})

	r.GET("/get-meetings-details", func(ctx *gin.Context) {
		c.GetMeetingDetails(ctx)
	})

	r.GET("/display-cases", func(ctx *gin.Context) {

		c.CheckLogin(ctx, "display-cases.html", "/lawyer-login", "", "")

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

		if c.CheckLogin(ctx, "", "", caseID, "") {

			//create a meeting

			meetingID, err := c.CreateNewMeeting(caseID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating meeting"})
				return
			}

			// Pass caseID to the template
			ctx.HTML(http.StatusOK, "intake.html", gin.H{"case_id": caseID, "meeting_id": meetingID})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Lawyer not authenticated for this case"})
			return
		}
	})

	r.GET("/cases", func(ctx *gin.Context) {
		c.GetAllCasesForLawyer(ctx)
	})

	r.GET("/check_login", func(ctx *gin.Context) {
		c.CheckLogin(ctx, "", "", "", "")
	})

	r.GET("/assign-case", func(ctx *gin.Context) {
		c.AssignCaseToLawyer(ctx)
	})

	r.GET("/delete-meeting/:case_id/:meeting_id", func(ctx *gin.Context) {

		caseID := ctx.Param("case_id")
		meetingID := ctx.Param("meeting_id")

		if c.CheckLogin(ctx, "", "", caseID, "") { //lawyer is authenticated for this case

			//delete the meeting
			err := c.DeleteMeeting(meetingID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting meeting"})
				return
			}

			//if the meeting got deleted, redirect to the case details page
			ctx.Redirect(http.StatusFound, "/case-details/"+caseID)

		} else {
			//this would only occur if someone hit the page directly and not from a button while viewing the meeting or
			//while on the intake page and they weren't logged in or assigned to the case
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Lawyer not authenticated for this case"})
			return
		}
	})

	r.POST("/create_case", func(ctx *gin.Context) {
		c.CreateNewCase(ctx)
	})

	r.POST("/save_lawyer_notes", func(ctx *gin.Context) {
		c.AddNotesToMeeting(ctx)
	})

	r.POST("/lawyer_login", func(ctx *gin.Context) {
		c.AuthenticateLawyer(ctx)
	})

	r.POST("/create_lawyer_account", func(ctx *gin.Context) {
		c.CreateNewLawyer(ctx)
	})

}
