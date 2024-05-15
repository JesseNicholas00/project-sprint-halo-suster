package medicalrecord

type Patient struct {
	IdentityNumber int64  `db:"identity_number"`
	PhoneNumber    string `db:"phone_number"`
	BirthDate      string `db:"birth_date"`
	Gender         string `db:"gender"`
	ImageUrl       string `db:"image_url"`
	CreatedAt      string `db:"created_at"`
}
