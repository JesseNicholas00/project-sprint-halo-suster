package medicalrecord

type CreatePatientReq struct {
	IdentityNumber  int64  `json:"identityNumber"      validate:"required,intlen=16"`
	PhoneNumber     string `json:"phoneNumber"         validate:"required,phoneNum"`
	Name            string `json:"name"                validate:"required,min=3,max=30"`
	BirthDate       string `json:"birthDate"           validate:"required,iso8601"`
	Gender          string `json:"gender"              validate:"required,oneof=male female"`
	IdentityCardImg string `json:"identityCardScanImg" validate:"required,url,imageExt"`
}

type CreatePatientRes struct {
}
