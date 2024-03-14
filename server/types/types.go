package types

type Case struct {
	ID          int    `json:"id"`
	ClientName  string `json:"client_name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Contact     string `json:"contact"`
	Interviewed int    `json:"interviewed"`
	LawyerID    int    `json:"lawyer_id"`
}
