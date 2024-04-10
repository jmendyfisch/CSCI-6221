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
	AddressStreet   string    `json:"address_street"`
	AddressCity     string    `json:"address_city"`
	AddressState    string    `json:"address_state"`
	AddressZip      string    `json:"address_zip"`
}

type NewCaseResp struct {
	CaseID int `json:"case_id"`
}

type LawyerLogin struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

type Lawyer struct {
	LawyerFirstName string `json:"lawyer_first_name"`
	LawyerLastName  string `json:"lawyer_last_name"`
	EmailAddress    string `json:"email_address"`
	Password        string `json:"password"`
}

type NewLawyerResp struct {
	LawyerEmail string `json:"email_address"`
}

type GPTPromptOutput struct {
	MeetingID        int      `json:"meeting_id"`
	Questions        []string `json:"questions" binding:"required"`
	Summary          string   `json:"summary" binding:"required"`
	AdditionalPoints []string `json:"points" binding:"required"`
}

type Notes struct {
	MeetingID int    `json:"meeting_id"`
	Notes     string `json:"notes"`
}
