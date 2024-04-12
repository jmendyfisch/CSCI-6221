package database

// All the queries for the database

const (
	GetAllUnassignedCasesQ = `select id, created_at, client_first_name, client_last_name, type, description, phone_number, email_address, from cases where lawyer_id = 0`

	CreateLawyerQ = `insert into lawyers (lawyer_first_name, lawyer_last_name, email_address, password) values ($1, $2, $3, $4) returning email_address`

	/*CreateCaseQ = `insert into cases (created_at,client_first_name, client_last_name, type, description, phone_number, email_address, lawyer_id) values
	(timestamp 'now()', $1, $2, $3, $4, $5, $6, 1) returning id`*/

	CreateCaseQ = `insert into cases (created_at,client_first_name, client_last_name, type, description, phone_number, email_address, lawyer_id, address_street, address_city, address_state, address_zip) values 
	(timestamp 'now()', $1, $2, $3, $4, $5, $6, 1, $7, $8, $9, $10) returning id`

	LawyerLoginQ = `SELECT id, password FROM lawyers WHERE email_address = $1`

	CreateMeetingQ = `insert into meetings(created_at, case_id, lawyer_id) values (now(), $1, (select lawyer_id from cases where id = $1 limit 1)) returning id`

	CreateGPTRespQ = `insert into gpt_resp(created_at, meeting_id, questions, summary, points) values (now(), $1, $2, $3, $4) returning id`

	AddNotesToMeetingQ = `update meetings set lawyer_notes = $2 where meeting_id = $1 returning id`
)
