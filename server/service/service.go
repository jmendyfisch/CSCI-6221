package service

import (
	"errors"
	"log"
	"server/database"
	"server/types"

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
		log.Println("err: ", err.Error())
		return nil, ErrInvalidLawyerID
	}

	return
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

	id, password, error := database.GetLawyerByEmail(c)
	if error != nil {
		log.Println("db err: ", error.Error())
		return false, id, err
	}

	/* Code used to debug if passwords aren't matching

	var ErrHashingPassword = errors.New("error hashing password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("bcrypt error: ", err.Error())
		return false, ErrHashingPassword
	}

	log.Println(string(hashedPassword))
	log.Println("in db:" + password)*/

	// Compare the hashed password with the stored password
	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(c.Password))
	if err != nil {
		return false, id, nil // Password does not match
	}

	return true, id, nil // Success
}
