package healthproblemsqlstatements

const SelectHealthProblemByClientId = `
		SELECT
			hp.id,
			hp.name,
			hp.client_id,
			hp.grau,
			hp.created_at,
			hp.updated_at
		FROM
			tb_health_problems hp
		WHERE
			hp.client_id = $1
`
