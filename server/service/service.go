/*
The Service package is responsible for implementing the actual business logic based on
the requirements from the Controller. A Controller object is always assoicated with a Service object.

It is also responsible to request and format the database query outputs and propagate any errors that arise.

Original for project, based on controller-service-repository architecture (ChatGPT assisted with development)

*/

package service

import (
	"errors"
	"log"
	"server/database"
	"server/types"

	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
}

func New() Service {
	return Service{}
}

// Gets all cases for a given lawyer ID
func (s *Service) GetAllCases(lawyerID int) (cases []types.Case, err error) {
	log.Println("called service.GetAllCases()")

	cases, err = database.GetAllCasesForLawyer(lawyerID)
	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("err: ", err.Error())
			return nil, ErrInvalidLawyerID
		}
		return nil, ErrQueryFailure
	}

	return
}

func (s *Service) GetCaseDetails(caseID int) (c types.Case, err error) {
	log.Println("called service.GetAllCases()")

	c, err = database.GetCaseDetails(caseID)
	//log.Println("case: ", c)
	c.ID = caseID

	return
}

func (s *Service) GetAllMeetings(caseID int) (m []types.Meeting, err error) {
	log.Println("called service.GetAllMeetings()")

	m, err = database.GetAllMeetings(caseID)
	if err != nil {
		log.Println("db err: ", err.Error())
		if err == pgx.ErrNoRows {

			return nil, ErrInvalidCaseID
		}
		return nil, ErrQueryFailure
	}

	return

}

func (s *Service) GetMeetingDetails(meetingID int) (r types.MeetingDetails, err error) {
	log.Println("called service.GetMeetingDetails()")

	r.Meet, err = database.GetMeetingDetails(meetingID)
	if err != nil {
		return r, err
	}

	r.GPTResp, err = database.GetGPTResponses(meetingID)
	if err != nil {
		return r, err
	}

	return r, nil

}

// Create a new case, it is assigned to a default lawyer
func (s *Service) CreateNewCase(c types.Case) (caseID int, err error) {
	log.Println("called service.CreateNewCase()")

	caseID, err = database.CreateNewCase(c)
	if err != nil {
		log.Println("db err: ", err.Error())
		return 0, ErrQueryFailure
	}

	return
}

func (s *Service) CreateNewMeeting(caseID string) (meetingID int, err error) {
	log.Println("called service.CreateNewMeeting()")

	meetingID, err = database.CreateNewMeeting(caseID)
	if err != nil {
		log.Println("db err: ", err.Error())
		return 0, ErrQueryFailure
	}

	return
}

func (s *Service) DeleteMeeting(meetingID string) (err error) {
	log.Println("called service.DeleteMeeting()")

	err = database.DeleteMeeting(meetingID)
	if err != nil {
		log.Println("db err: ", err.Error())
		return ErrQueryFailure
	}

	return
}

func (s *Service) CreateNewLawyer(c types.Lawyer) (LawyerEmail string, err error) {
	log.Println("called service.CreateNewLawyer()")

	var ErrHashingPassword = errors.New("error hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error: ", err.Error())
		return "", ErrHashingPassword
	}

	LawyerEmail, err = database.CreateNewLawyer(c, string(hashedPassword))
	if err != nil {
		log.Println("db err: ", err.Error())
		return "", ErrQueryFailure
	}

	return
}

func (s *Service) AuthenticateLawyer(c types.LawyerLogin) (Success bool, LawyerId int, err error) {

	log.Println("called service.AuthenticateLawyer()")

	id, password, err := database.GetLawyerByEmail(c)
	if err != nil {
		log.Println("db err: ", err.Error())
		return false, id, err
	}

	//log.Println("passwords: ", password, c.Password)

	//return password == c.Password, id, nil

	// Compare the hashed password with the stored password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(c.Password))
	if err != nil {
		return false, id, nil // Password does not match
	}

	return true, id, nil // Success
}

func (s *Service) ProcessInterview(caseID int, meetingID int, gptSummaries []string, interviewAudio []byte, filepath string) (gptRes types.GPTPromptOutput, err error) {
	// step 1 - convert audio to text
	log.Println("called service.ProcessInterview()")

	interviewText, err := getTextFromAudio(filepath)
	if err != nil {
		return
	}

	log.Println("interview text: ", interviewText)

	// step 2 - send text, ask to split it into lawyer and client, get summary and addtl questions to ask client.
	gptRes, err = getOutputTextFromTranscription(caseID, meetingID, interviewText, gptSummaries)
	if err != nil {
		return
	}

	// store and return result
	log.Println("meeting id from service.ProcessInterview(): ", meetingID)
	_, err = database.AddGPTResponse(meetingID, gptRes)

	return
	// optionally ask for the client's summary
}

func (s *Service) AddNotesToMeeting(meetingID string, notes string) (err error) {
	err = database.AddNotesToMeeting(meetingID, notes)
	return
}

func (s *Service) GenAndStoreCaseSummary(caseID int, caseDescription string, gptSummaries []string) (err error) {

	summary, err := getCaseSummary(caseDescription, gptSummaries)
	if err != nil {
		return err
	}

	err = database.UpdateCaseSummary(caseID, summary)
	return err
}

func (s *Service) AssignCaseToLawyer(caseID, lawyerID int) (err error) {
	return database.AssignCaseToLawyer(caseID, lawyerID)
}
