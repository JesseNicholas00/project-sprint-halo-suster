package medicalrecord

type RegisterPatientReq struct {
	IdentityNumber  int64  `json:"identityNumber"      validate:"required,intlen=16"`
	PhoneNumber     string `json:"phoneNumber"         validate:"required,phoneNum"`
	Name            string `json:"name"                validate:"required,min=3,max=30"`
	BirthDate       string `json:"birthDate"           validate:"required,iso8601"`
	Gender          string `json:"gender"              validate:"required,oneof=male female"`
	IdentityCardImg string `json:"identityCardScanImg" validate:"required,url,imageExt"`
}

type RegisterPatientRes struct {
}

type ListPatientsReq struct {
	IdentityNumber *int64  `query:"identityNumber"`
	Limit          *int    `query:"limit"`
	Offset         *int    `query:"offset"`
	Name           *string `query:"name"`
	PhoneNumber    *string `query:"phoneNumber"`
	CreatedAtSort  *string `query:"createdAt"`
}

type ListPatientsRes struct {
	Data []ListPatientsResData
}

type ListPatientsResData struct {
	IdentityNumber int64  `json:"identityNumber"`
	PhoneNumber    string `json:"phoneNumber"`
	Name           string `json:"name"`
	BirthDate      string `json:"birthDate"`
	Gender         string `json:"gender"`
	CreatedAt      string `json:"createdAt"`
}
