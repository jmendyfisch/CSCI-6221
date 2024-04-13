package types

import (
	"database/sql"
	"time"
)

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
	GPTSummary      string    `json:"gpt_summary"`
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

type MeetingDetails struct {
	Meet    Meeting           `json:"meeting"`
	GPTResp []GPTRespInternal `json:"gpt_resp"`
}

type GPTRespInternal struct {
	ID               int       `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	MeetingID        int       `json:"meeting_id"`
	Questions        string    `json:"questions"`
	Summary          string    `json:"summary"`
	AdditionalPoints string    `json:"points"`
}

type Notes struct {
	MeetingID int    `json:"meeting_id"`
	Notes     string `json:"notes"`
}

type Meeting struct {
	ID          int            `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	CaseID      int            `json:"case_id"`
	LawyerID    int            `json:"lawyer_id"`
	LawyerNotes sql.NullString `json:"lawyer_notes"`
}
