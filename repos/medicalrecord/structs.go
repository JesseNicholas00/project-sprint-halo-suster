package medicalrecord

type Patient struct {
	IdentityNumber int64  `db:"identity_number"`
	PhoneNumber    string `db:"phone_number"`
	Name           string `db:"name"`
	BirthDate      string `db:"birth_date"`
	Gender         string `db:"gender"`
	ImageUrl       string `db:"image_url"`
	CreatedAt      string `db:"created_at"`
}

type PatientFilter struct {
	IdentityNumber *int64
	Limit          int
	Offset         int
	Name           *string
	PhoneNumber    *string
	CreatedAt      *string
}
