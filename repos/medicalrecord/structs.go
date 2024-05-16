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
