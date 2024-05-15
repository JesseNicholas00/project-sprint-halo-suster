package medicalrecord

import "errors"

var (
	ErrDuplicateIdentityNumber = errors.New(
		"medicalRecordRepository: duplicate identity number",
	)
)
