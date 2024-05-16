package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createPatient *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		createPatient: statementutil.MustPrepareNamed(`
			INSERT INTO patients(
				identity_number,
				phone_number,
				birth_date,
				gender,
				image_url
			) VALUES (
				:identity_number,
				:phone_number,
				:birth_date,
				:gender,
				:image_url
			)
		`),
	}
}
