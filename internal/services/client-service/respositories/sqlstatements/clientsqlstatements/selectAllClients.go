package clientsqlstatements

const SelectAllClients = `
		SELECT
			c.id,
			c.name,
			c.birth_date,
			c.sexo,
			c.created_at,
			c.updated_at
		FROM
			tb_clients c
`
