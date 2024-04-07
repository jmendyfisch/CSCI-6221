package database

// All the queries for the database

const (
	GetAllUnassignedCasesQ = `select id, created_at, client_first_name, client_last_name, type, description, phone_number, email_address, from cases where lawyer_id = 0`
	CreateCaseQ            = `insert into cases (created_at,client_first_name, client_last_name, type, description, phone_number, email_address, lawyer_id) values 
	(timestamp 'now()', $1, $2, $3, $4, $5, $6, 1) returning id`
)
