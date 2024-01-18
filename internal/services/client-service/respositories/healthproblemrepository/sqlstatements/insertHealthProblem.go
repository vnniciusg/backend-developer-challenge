package sqlstatements

const InsertHealthProblem = `
	INSERT INTO tb_health_problems 
	(client_id, name, grau , created_at, updated_at) 
	VALUES ($1, $2, $3 , $4, $5) 
	RETURNING id , client_id, name, grau, created_at, updated_at
`
