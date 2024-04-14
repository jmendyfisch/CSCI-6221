/*
The Database package is responsible for connecting to the Postgres server
and handling all database queries for the Service package.
Original for project, based on controller-service-repository architecture (ChatGPT assisted with development)
Utilizes PGX library for Postgres database connection.
*/

package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"server/config"
	"server/types"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	conn *pgxpool.Pool
)

var (
	ErrMeetingInsert  error = errors.New("could not add a new meeting")
	ErrInvalidLawyer  error = errors.New("lawyer does not exist")
	ErrNoMeetingFound error = errors.New("no meeting found with given id")
	ErrNoCaseFound    error = errors.New("no case found with given id")
)

func Init() {
	log.Println("Connecting to database")

	dbConfig, err := pgxpool.ParseConfig("postgresql://" + config.DBUser + ":" + config.DBPassword + "@" + config.DBHost + ":" + config.DBPort + `/` + config.DBName)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}
	conn, err = pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Println("err connecting with db: ", err.Error())
		panic(err)
	}
}

func GetAllCasesForLawyer(lawyerID int) (cases []types.Case, err error) {
	log.Println("inside database.GetAllCases()")

	var c types.Case

	rows, err := conn.Query(context.Background(), GetAllCasesQ, lawyerID)
	if err != nil {
		log.Println("error with db: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.CreatedAt, &c.ClientFirstName, &c.ClientLastName, &c.Type, &c.Description, &c.PhoneNumber, &c.EmailAddress, &c.LawyerID)
		if err != nil {
			fmt.Println("err: ", err)
			return nil, err
		}
		//fmt.Println(c)

		cases = append(cases, c)
	}

	return
}

func CreateNewCase(c types.Case) (caseID int, err error) {
	log.Println("inside database.CreateNewCase()")

	row := conn.QueryRow(context.Background(), CreateCaseQ, c.ClientFirstName, c.ClientLastName, c.Type, c.Description, c.PhoneNumber, c.EmailAddress, c.AddressStreet, c.AddressCity, c.AddressState, c.AddressZip)
	err = row.Scan(&c.ID)

	return
}

func CreateNewLawyer(c types.Lawyer, hashedPassword string) (LawyerEmail string, err error) {
	log.Println("inside database.CreateNewLawyer()")
	var email = strings.ToLower(c.EmailAddress)

	//log.Println(c.LawyerFirstName)
	row := conn.QueryRow(context.Background(), CreateLawyerQ, c.LawyerFirstName, c.LawyerLastName, email, hashedPassword)
	err = row.Scan(&email)

	return
}

func GetLawyerByEmail(c types.LawyerLogin) (int, string, error) {
	log.Println("inside database.GetLawyerByEmail()")
	emailAddress := strings.ToLower(c.EmailAddress)
	var password string
	var id int
	row := conn.QueryRow(context.Background(), LawyerLoginQ, emailAddress)
	err := row.Scan(&id, &password)

	if err != nil {
		log.Printf("Failed to get lawyer by email: %v\n", err)
		return 0, "", err // Return 0, an empty string and the error
	}

	return id, password, nil // Return the password and nil as the error
}

func CreateNewMeeting(caseID string) (meetingID int, err error) {
	log.Println("inside database.CreateNewMeeting()")

	row := conn.QueryRow(context.Background(), CreateMeetingQ, caseID)
	err = row.Scan(&meetingID)

	return
}

func DeleteMeeting(meetingID string) (err error) {
	log.Println("inside database.DeleteMeeting()")

	var success bool
	row1 := conn.QueryRow(context.Background(), DeleteMeetingGPTRespQ, meetingID)
	err = row1.Scan(&success)
	if err != nil {
		log.Println("Deleting meeting "+meetingID+". For deleting associated GPT responses, get error", err)
	}
	row2 := conn.QueryRow(context.Background(), DeleteMeetingQ, meetingID)
	err = row2.Scan(&success)
	return
}

func AddGPTResponse(meetingID int, gptResp types.GPTPromptOutput) (gptResptID int, err error) { //Rename this function to AddGPTResponse
	log.Println("inside database.AddGPTResponse()")

	questions, points := "", ""
	for _, iter := range gptResp.Questions {
		questions += iter + "\n"
	}

	for _, iter := range gptResp.AdditionalPoints {
		points += iter + "\n"
	}

	log.Println("Meeting id from database.AddGPTResponse()", meetingID)
	row := conn.QueryRow(context.Background(), CreateGPTRespQ, meetingID, questions, gptResp.Summary, points)
	err = row.Scan(&gptResptID)
	if err != nil {
		log.Println("db error: ", err)
		return 0, ErrMeetingInsert
	}

	return gptResptID, nil
}

func AddNotesToMeeting(meetingID string, notes string) error {
	log.Println("inside database.AddNotesToMeeting()")

	row := conn.QueryRow(context.Background(), AddNotesToMeetingQ, meetingID, notes)
	err := row.Scan(&meetingID)

	if err != nil {
		log.Println("db error: ", err)
		if err == pgx.ErrNoRows {
			return ErrNoMeetingFound
		}
	}

	return err
}

func GetCaseDetails(caseID int) (c types.Case, err error) {
	log.Println("inside database.GetCaseDetails()")

	row := conn.QueryRow(context.Background(), GetCaseDetailsQ, caseID)
	err = row.Scan(&c.CreatedAt, &c.ClientFirstName, &c.ClientLastName, &c.PhoneNumber, &c.EmailAddress, &c.Type, &c.Description, &c.LawyerID, &c.AddressStreet, &c.AddressCity, &c.AddressState, &c.AddressZip, &c.GPTSummary)

	if err != nil {
		log.Println("db error: ", err)
		if err == pgx.ErrNoRows {
			return c, ErrNoCaseFound
		}
	}

	return
}

func GetAllMeetings(caseID int) (meetings []types.Meeting, err error) {
	log.Println("inside database.GetAllMeetings()")

	var m types.Meeting

	rows, err := conn.Query(context.Background(), GetAllMeetingsQ, caseID)
	if err != nil {
		log.Println("error with db: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&m.ID, &m.CreatedAt, &m.CaseID, &m.LawyerID, &m.LawyerNotes)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, m)
	}

	return
}

func GetMeetingDetails(meetingID int) (m types.Meeting, err error) {
	log.Println("inside database.GetMeetingDetails()")

	row := conn.QueryRow(context.Background(), GetMeetingDetailsQ, meetingID)
	err = row.Scan(&m.ID, &m.CreatedAt, &m.CaseID, &m.LawyerID, &m.LawyerNotes)

	if err != nil {
		log.Println("db error: ", err)
		if err == pgx.ErrNoRows {
			return m, ErrNoMeetingFound
		}
	}

	return
}

func GetGPTResponses(meetingID int) (resps []types.GPTRespInternal, err error) {
	log.Println("inside database.GetGPTResponses()")

	var m types.GPTRespInternal

	rows, err := conn.Query(context.Background(), GetGPTRespsQ, meetingID)
	if err != nil {
		log.Println("error with db: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&m.ID, &m.CreatedAt, &m.MeetingID, &m.Questions, &m.Summary, &m.AdditionalPoints)
		if err != nil {
			return nil, err
		}
		resps = append(resps, m)
	}

	return
}

func UpdateCaseSummary(caseID int, summary string) (err error) {
	log.Println("inside database.UpdateCaseSummary()")

	var success bool

	row := conn.QueryRow(context.Background(), UpdateCaseSummaryQ, caseID, summary)
	err = row.Scan(&success)

	if err != nil {
		log.Println("db error: ", err)
		if err == pgx.ErrNoRows {
			return ErrNoCaseFound
		}
	}

	return
}

func AssignCaseToLawyer(caseID, lawyerID int) (err error) {
	log.Println("inside database.AssignCaseToLawyer()")

	var success bool

	row := conn.QueryRow(context.Background(), AssignCaseToLawyerInCasesQ, caseID, lawyerID)
	err = row.Scan(&success)

	log.Println(AssignCaseToLawyerInCasesQ)
	log.Println("caseID " + strconv.Itoa(caseID) + " lawyerID " + strconv.Itoa(lawyerID))

	if err != nil {
		log.Println("db error: ", err)
		if err == pgx.ErrNoRows {
			return ErrNoCaseFound
		}
	}

	row = conn.QueryRow(context.Background(), AssignCaseToLawyerInMeetingsQ, caseID, lawyerID)
	err = row.Scan(&success)

	//log.Println(AssignCaseToLawyerInMeetingsQ)

	if err != nil {
		if err == pgx.ErrNoRows {
			log.Println("Note: This case has no meetings.")
			return nil
		}
	}

	return
}
