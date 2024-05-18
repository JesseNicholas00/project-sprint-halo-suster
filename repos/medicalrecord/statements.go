package medicalrecord

import (
	"github.com/JesseNicholas00/HaloSuster/utils/statementutil"
	"github.com/jmoiron/sqlx"
)

type statements struct {
	createPatient *sqlx.NamedStmt
	createRecord  *sqlx.NamedStmt
}

func prepareStatements() statements {
	return statements{
		createPatient: statementutil.MustPrepareNamed(`
			INSERT INTO patients(
				identity_number,
				phone_number,
				name,
				birth_date,
				gender,
				image_url
			) VALUES (
				:identity_number,
				:phone_number,
				:name,
				:birth_date,
				:gender,
				:image_url
			)
		`),
		createRecord: statementutil.MustPrepareNamed(`
			INSERT INTO medical_records(
				patient_identity_number,
				patient_phone_number,
				patient_name,
				patient_birth_date,
				patient_gender,
				patient_image_url,
				symptoms,
				medications,
				created_by
			) VALUES (
				:patient_identity_number,
				:patient_phone_number,
				:patient_name,
				:patient_birth_date,
				:patient_gender,
				:patient_image_url,
				:symptoms,
				:medications,
				:created_by
			)
		`),
	}
}
