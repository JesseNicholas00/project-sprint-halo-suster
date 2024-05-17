package medicalrecord

import "errors"

var (
	ErrDuplicateIdentityNumber = errors.New(
		"medicalRecordService: duplicate identity number found",
	)
	ErrIdentityNumberNotFound = errors.New(
		"medicalRecordService: no such identity number found",
	)
)
