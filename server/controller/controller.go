package controller

import (
	"log"
	"net/http"
	"server/service"
	"server/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*
The Controller handles sending data to the user using g.Context objects.
*/

type Controller struct {
	serv service.Service
}

func New(s service.Service) Controller {
	return Controller{serv: s}
}

func (c *Controller) GetAllCasesForLawyer(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	lIDString := params.Get("lawyer_id")
	if lIDString == "" {
		log.Println("no Lawyer ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no lawyer id provided"})
		return
	}

	lID, _ := strconv.ParseInt(params.Get("lawyer_id"), 10, 64)

	cases, err := c.serv.GetAllCases(int(lID))

	if err == service.ErrInvalidLawyerID {
		log.Println("Invalid Lawyer ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.GetAllCasesForLawyer()")
	ctx.JSON(http.StatusOK, cases)
}

func (c *Controller) CreateNewCase(ctx *gin.Context) {

	var newCase types.Case
	var resp types.NewCaseResp
	var err error

	if err = ctx.BindJSON(&newCase); err != nil {
		log.Println("incorrect case format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect case format"})
		return
	}

	resp.CaseID, err = c.serv.CreateNewCase(newCase)

	if err == service.ErrQueryFailure {
		log.Println("db error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.CreateNewCase()")
	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) CreateNewLawyer(ctx *gin.Context) {
	log.Println("inside controller.CreateNewLawyer()")
	var newLawyer types.Lawyer
	var resp types.NewLawyerResp
	var err error

	if err = ctx.BindJSON(&newLawyer); err != nil {
		log.Println("incorrect lawyer format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect lawyer format"})
		return
	}

	log.Println("newLawyer: ", newLawyer)

	resp.LawyerEmail, err = c.serv.CreateNewLawyer(newLawyer)

	if err == service.ErrQueryFailure {
		log.Println("db error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.CreateNewLawyer()")
	ctx.JSON(http.StatusOK, resp)
}

func (c *Controller) AuthenticateLawyer(ctx *gin.Context) {
	var lawyer types.LawyerLogin
	var err error

	if err = ctx.BindJSON(&lawyer); err != nil {
		log.Println("incorrect lawyer login format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "incorrect lawyer login format"})
		return
	}

	isAuthenticated, LawyerID, err := c.serv.AuthenticateLawyer(lawyer)

	if err == service.ErrQueryFailure {
		log.Println("db error")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !isAuthenticated {
		log.Println("incorrect email or password")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect email or password"})
		return
	}

	log.Println("success for controller.AuthenticateLawyer()")
	ctx.JSON(http.StatusOK, gin.H{"message": "authenticated", "lawyer_id": LawyerID})

	// I tried to do the Gorilla Cookie thing. That didn't work. I'm commenting out all the Gorilla stuff and will delete later.
	/*
		request := ctx.Request
		session, err := c.store.Get(ctx.Request, "session-name")
		if err != nil {
			log.Println("Error retrieving session:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.Set("session", session)
		ctx.Set("lawyer_id", LawyerID)

		// Set lawyer ID in session
		session.Values["lawyer_id"] = LawyerID

		session.Options = &sessions.Options{
			Path:     "/",   // Available throughout the site
			MaxAge:   86400, // Expires after 1 day
			HttpOnly: false, // Make accessible via JavaScript
			Secure:   false, // When running in localhost, we are not over HTTPS
			SameSite: http.SameSiteLaxMode,
		}

		// Save the session
		err = session.Save(request, ctx.Writer)
		if err != nil {
			log.Println("Error saving session:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	*/

}
