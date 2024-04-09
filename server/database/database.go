package database

import (
	"context"
	"log"
	"server/config"
	"server/types"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	conn *pgxpool.Pool
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

	rows, err := conn.Query(context.Background(), GetAllUnassignedCasesQ, lawyerID)
	if err != nil {
		log.Println("error with db: ", err.Error())
		return
	}

	for rows.Next() {
		err = rows.Scan(&c.ID, &c.ClientFirstName, &c.ClientLastName, &c.Type, &c.Description, &c.PhoneNumber, &c.EmailAddress, &c.LawyerID)
		if err != nil {
			return nil, err
		}
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

	log.Println(c.LawyerFirstName)
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
