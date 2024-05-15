package medicalrecord

import "errors"

var (
	ErrDuplicateIdentityNumber = errors.New(
		"medicalRecordService: duplicate identity number found",
	)
)
