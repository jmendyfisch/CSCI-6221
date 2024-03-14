package database

// All the queries for the database

const (
	GetAllCasesQ = `select id, client_name, type, description, contact, interviewed, lawyer_id from cases where lawyer_id = $1`
	CreateCaseQ  = `insert into cases (client_name, type, description, contact, interviewed, lawyer_id) values ($1, $2, $3, $4, $5, $6) returning true`
)
