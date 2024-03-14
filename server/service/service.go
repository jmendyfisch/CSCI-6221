package service

import (
	"log"
	"server/database"
	"server/types"
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

// Create a new case and assign a lawyer to it
func (s *Service) CreateNewCase(c types.Case) (caseID int, lawyerName string, err error) {
	log.Println("called service.CreateNewCase()")

	caseID, lawyerName, err = database.CreateNewCase(c)
	if err != nil {
		log.Println("db err: ", err.Error())
		return 0, "", ErrQueryFailure
	}

	return
}
