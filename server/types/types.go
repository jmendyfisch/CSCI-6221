package types

type Case struct {
	ID          int    `json:"id"`
	ClientName  string `json:"client_name" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
	Contact     string `json:"contact" binding:"required"`
	Interviewed int    `json:"interviewed"`
	LawyerID    int    `json:"lawyer_id"`
}

type NewCaseResp struct {
	CaseID     int    `json:"case_id"`
	LawyerName string `json:"lawyer_name"`
}
