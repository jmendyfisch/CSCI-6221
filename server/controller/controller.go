package controller

import (
	"log"
	"net/http"
	"server/service"
	"server/types"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
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

	// Start a session
	c.StartLawyerSession(LawyerID, ctx.Writer, ctx.Request)
}

var (
	key   = []byte("lawyer-login-authentication-key-1234")
	store = sessions.NewCookieStore(key)
)

// StartLawyerSession starts a session for a lawyer

func (c *Controller) StartLawyerSession(lawyerID int, response http.ResponseWriter, request *http.Request) {
	session, err := store.Get(request, "session-name")
	if err != nil {
		log.Println("Error retrieving session:", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set lawyer ID in session
	session.Values["lawyer_id"] = lawyerID
	log.Println("lawyer_id in session:", lawyerID)
	session.Options = &sessions.Options{
		Path:     "/",   // Available throughout the site
		MaxAge:   86400, // Expires after 1 day
		HttpOnly: false, // Make accessible via JavaScript
		Secure:   false, // When running in localhost, we are not over HTTPS
	}

	// Save the session
	err = session.Save(request, response)
	if err != nil {
		log.Println("Error saving session:", err)
		http.Error(response, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	//session_test, _ := store.Get(request, "session-name")
	//lawyer_id_test := session_test.Values["lawyer_id"]
	//log.Println("lawyer_id in session test:", lawyer_id_test)
}

func (c *Controller) ReturnLawyerSession(request *http.Request) (int, error) {
	session, err := store.Get(request, "session-name")
	if err != nil {
		log.Println("Error retrieving session:", err)
		return 0, err
	}

	lawyer_id := session.Values["lawyer_id"]
	log.Println("lawyer_id in session:", lawyer_id)
	return lawyer_id.(int), nil
}
