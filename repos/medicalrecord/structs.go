package medicalrecord

import "time"

type Patient struct {
	IdentityNumber int64     `db:"identity_number"`
	PhoneNumber    string    `db:"phone_number"`
	Name           string    `db:"name"`
	BirthDate      time.Time `db:"birth_date"`
	Gender         string    `db:"gender"`
	ImageUrl       string    `db:"image_url"`
	CreatedAt      time.Time `db:"created_at"`
}

type PatientFilter struct {
	IdentityNumber *int64
	Limit          int
	Offset         int
	Name           *string
	PhoneNumber    *string
	CreatedAtSort  *string
}

type RecordFilter struct {
	IdentityNumber  *int64
	CreatedByUserId *string
	Limit           int
	Offset          int
	CreatedAtSort   *string
}

type Record struct {
	RecordId           int64     `db:"medical_record_id"`
	PatientId          int64     `db:"patient_identity_number"`
	PatientPhoneNumber string    `db:"patient_phone_number"`
	PatientName        string    `db:"patient_name"`
	PatientBirthDate   time.Time `db:"patient_birth_date"`
	PatientGender      string    `db:"patient_gender"`
	PatientImageUrl    string    `db:"patient_image_url"`
	Symptoms           string    `db:"symptoms"`
	Medications        string    `db:"medications"`
	CreatedAt          time.Time `db:"created_at"`
	CreatedByUserId    string    `db:"created_by"`
}
