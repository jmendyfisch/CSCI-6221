package controller

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"server/config"
	"server/database"
	"server/service"
	"server/types"
	"strconv"
	"strings"
	"time"

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
	//Adding our own Cookie security.  This will set a hash string dependent on our CookieKey that is
	//only known in our backend.  This would prevent someone from just spoofing
	//a cookie containing a LawyerID integer. In implementation the whole site would be served over https.
	timeStamp := time.Now().UnixNano()
	timeStampStr := strconv.FormatInt(timeStamp, 10)

	data := config.CookieKey + timeStampStr + strconv.Itoa(LawyerID)
	hash := sha256.Sum256([]byte(data))
	securityString := hex.EncodeToString(hash[:])

	log.Println("success for controller.AuthenticateLawyer()")
	ctx.JSON(http.StatusOK, gin.H{"message": "authenticated", "lawyer_id": LawyerID, "timestamp": timeStampStr, "securitystring": securityString})

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
		err = session.Save(request, response)
		if err != nil {
			log.Println("Error saving session:", err)
			http.Error(response, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		//session_test, _ := store.Get(request, "session-name")
		//lawyer_id_test := session_test.Values["lawyer_id"]
		//log.Println("lawyer_id in session test:", lawyer_id_test)
	*/
}

// func (c *Controller) ReturnLawyerSession(request *http.Request) (int, error) {
// 	session, err := store.Get(request, "session-name")
// 	if err != nil {
// 		log.Println("Error retrieving session:", err)
// 		return 0, err
// 	}

// 	lawyer_id := session.Values["lawyer_id"]
// 	log.Println("lawyer_id in session:", lawyer_id)
// 	return lawyer_id.(int), nil
// }

func (c *Controller) ProcessInterview(ctx *gin.Context) {
	// get audio file
	file, err := ctx.FormFile("audio")
	if err != nil {
		log.Println("could not find file in request")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	audioType := ctx.PostForm("type")
	if audioType == "" {
		log.Println("no audio type provided")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No audio type provided"})
		return
	}

	audioType = strings.TrimPrefix(audioType, "audio/")

	if audioType != config.AudioFileExtension {
		log.Println("incorrect audio type provided")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect audio type provided"})
		return
	}

	case_id := ctx.PostForm("case_id")
	if case_id == "" {
		log.Println("No case id provided")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No case id provided"})
		return
	}

	filename := fmt.Sprintf("rec_case_%s_%v.%s", case_id, time.Now().UnixNano(), audioType)
	path := fmt.Sprintf("/tempaudio/%s", filename)

	if err := ctx.SaveUploadedFile(file, path); err != nil {
		log.Println("err: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	buf := bytes.NewBuffer(nil)
	fileTemp, tempErr := file.Open()
	if _, err := io.Copy(buf, fileTemp); err != nil || tempErr != nil {
		log.Println("Could not open audio file ", filename, ", err: ", err.Error(), "or", tempErr)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	caseIDInt, _ := strconv.ParseInt(case_id, 10, 64)
	gptResp, err := c.serv.ProcessInterview(int(caseIDInt), buf.Bytes(), path)
	if err != nil {
		log.Println("Unsuccessful process, err: ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gptResp)
}

func (c *Controller) AddNotesToMeeting(ctx *gin.Context) {

	var notesType types.Notes
	if err := ctx.BindJSON(&notesType); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no meeting id provided"})
		return
	}

	err := c.serv.AddNotesToMeeting(notesType.MeetingID, notesType.Notes)
	if err == database.ErrNoMeetingFound {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid meeting id provided"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
