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
The Controller handles sending data to the user using gin.Context objects.
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

func (c *Controller) GetCaseDetails(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	cIDString := params.Get("case_id")
	if cIDString == "" {
		log.Println("no Case ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no case id provided"})
		return
	}

	cID, _ := strconv.ParseInt(cIDString, 10, 64)

	caseDet, err := c.serv.GetCaseDetails(int(cID))

	if err == database.ErrNoCaseFound {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.GetCaseDetails()")
	ctx.JSON(http.StatusOK, caseDet)
}

func (c *Controller) GetAllMeetings(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	cIDString := params.Get("case_id")
	if cIDString == "" {
		log.Println("no Case ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no case id provided"})
		return
	}

	cID, _ := strconv.ParseInt(cIDString, 10, 64)

	meetings, err := c.serv.GetAllMeetings(int(cID))

	if err == service.ErrInvalidCaseID {
		log.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("success for controller.GetAllMeetings()")
	ctx.JSON(http.StatusOK, meetings)
}

func (c *Controller) GetMeetingDetails(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	mIDString := params.Get("meeting_id")
	if mIDString == "" {
		log.Println("no Meeting ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no meeting id provided"})
		return
	}

	mID, _ := strconv.ParseInt(mIDString, 10, 64)

	mDet, err := c.serv.GetMeetingDetails(int(mID))
	if err != nil {
		if err == database.ErrNoMeetingFound {
			log.Println("no Meeting ID given")
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid meeting id provided"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, mDet)

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

func (c *Controller) CheckLogin(ctx *gin.Context, next string, redirect string, caseID string, meetingID string) {
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
		log.Println("Missing one or more required cookies.")
		if redirect != "" {
			ctx.Redirect(http.StatusFound, redirect)

		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "not authenticated"})
		}
		return
	}

	if caseID != "" {
		lawyerIDInt, _ := strconv.ParseInt(lawyerID, 10, 64)
		caseIDInt, _ := strconv.ParseInt(caseID, 10, 64)
		Cases, _ := database.GetAllCasesForLawyer(int(lawyerIDInt))

		log.Println(lawyerIDInt)
		log.Println(caseIDInt)

		caseFound := false
		for _, caseObj := range Cases {
			log.Println("In Cases loop")
			log.Println(caseObj.ID)
			if caseObj.ID == int(caseIDInt) {
				caseFound = true
				break
			}
		}

		if !caseFound {
			log.Println("Case ID not found for the lawyer")

			if redirect != "" {
				ctx.Redirect(http.StatusFound, redirect)

			} else {
				ctx.JSON(http.StatusOK, gin.H{"error": "Unauthorized access to the case"})
			}
			return

		}
	}

	data := config.CookieKey + timestamp + lawyerID
	hash := sha256.Sum256([]byte(data))

	if hex.EncodeToString(hash[:]) == securitystring {
		// If everything matches, return lawyer id, proceed to display cases.
		data := gin.H{"message": "authenticated", "lawyer_id": lawyerID, "timestamp": timestamp, "securitystring": securitystring, "case_id": caseID, "meeting_id": meetingID}
		if next != "" {
			ctx.HTML(http.StatusOK, next, data)
			return
		}
		ctx.JSON(http.StatusOK, data)

	} else {
		if redirect != "" {
			//log.Println(redirect)
			ctx.Redirect(http.StatusFound, redirect)

		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "not authenticated"})
		}
		return
	}

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

}

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

	var foundValidExt bool = false
	for _, ext := range config.AudioFileExtensions {
		if audioType == ext {
			foundValidExt = true
			break
		}
	}

	if !foundValidExt {
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

	// generate and update gpt_summary for this case
	caseDet, err := c.serv.GetCaseDetails(int(caseIDInt))
	if err != nil {
		log.Println("error fetching case details: ", err.Error())
	}
	desc := caseDet.Description

	meets, err := c.serv.GetAllMeetings(int(caseIDInt))
	if err != nil {
		log.Println("error fetching all meetings: ", err.Error())
	}

	var meetSummaries []string
	for _, iter := range meets {
		r, err := c.serv.GetMeetingDetails(iter.ID)
		if err != nil {
			log.Println("error fetching meeting details of id: ", iter.ID, err.Error())
		}

		for _, innerIter := range r.GPTResp {
			meetSummaries = append(meetSummaries, innerIter.Summary)
		}
	}

	err = c.serv.GenAndStoreCaseSummary(int(caseIDInt), desc, meetSummaries)
	if err != nil {
		log.Println("error gen and storing case summary: ", err.Error())
	}
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

func (c *Controller) AssignCaseToLawyer(ctx *gin.Context) {
	params := ctx.Request.URL.Query()
	cIDString := params.Get("case_id")
	lIDString := params.Get("lawyer_id")
	if cIDString == "" {
		log.Println("no Case ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no case id provided"})
		return
	}
	if lIDString == "" {
		log.Println("no Lawyer ID given")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "no lawyer id provided"})
		return
	}

	cID, _ := strconv.ParseInt(cIDString, 10, 64)
	lID, _ := strconv.ParseInt(lIDString, 10, 64)

	err := c.serv.AssignCaseToLawyer(int(cID), int(lID))
	if err != nil {
		log.Println("error assigning case to lawyer: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": "lawyer assigned to case successfully"})
}
