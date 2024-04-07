package types

import "time"

type Case struct {
	ID              int       `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	ClientFirstName string    `json:"client_first_name"`
	ClientLastName  string    `json:"client_last_name"`
	PhoneNumber     string    `json:"phone_number"`
	EmailAddress    string    `json:"email_address"`
	Type            string    `json:"type"`
	Description     string    `json:"description"`
	LawyerID        int       `json:"lawyer_id"`
}

type NewCaseResp struct {
	CaseID     int    `json:"case_id"`
}
