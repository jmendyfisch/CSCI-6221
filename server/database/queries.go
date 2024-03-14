package database

// All the queries for the database

const (
	GetAllUnprocessedCasesQ = `select id, client_name, type, description, contact, interviewed, lawyer_id from cases where lawyer_id = $1 and interviewed = 0`
	CreateCaseQ             = `insert into cases (client_name, type, description, contact, interviewed, lawyer_id) values 
	($1, $2, $3, $4, 0, (select lawyer_id from cases group by lawyer_id order by count(*) asc limit 1)) 
	returning id, (select lawyers.name from cases join lawyers on cases.lawyer_id = lawyers.id group by lawyer_id, lawyers.name order by count(*) asc limit 1)`
)
