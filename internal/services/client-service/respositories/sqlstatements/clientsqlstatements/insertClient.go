package clientsqlstatements

const InsertClient = `
	INSERT INTO tb_clients (
		name,
		birth_date,
		sexo,
		created_at,
		updated_at
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING id, name, birth_date, sexo , created_at, updated_at;
`
