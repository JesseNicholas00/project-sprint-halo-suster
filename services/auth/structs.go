package auth

import "github.com/golang-jwt/jwt/v4"

type RegisterItReq struct {
	Nip      int64  `json:"nip"      validate:"required,nip"`
	Name     string `json:"name"     validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type RegisterItRes struct {
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type RegisterNurseReq struct {
	Nip      int64  `json:"nip"                 validate:"required,nip"`
	Name     string `json:"name"                validate:"required,min=5,max=50"`
	ImageUrl string `json:"identityCardScanImg" validate:"required,imageExt"`
}

type RegisterNurseRes struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
	Name   string `json:"name"`
}

type GrantAccessNurseReq struct {
	UserId   string `param:"userId" validate:"required"`
	Password string `               validate:"required,min=5,max=33" json:"password"`
}

type GrantAccessNurseRes struct {
}

type UpdateNurseReq struct {
	UserId string `param:"userId" validate:"required"`
	Nip    int64  `               validate:"required,nip"          json:"nip"`
	Name   string `               validate:"required,min=5,max=50" json:"name"`
}

type UpdateNurseRes struct {
}

type DeleteNurseReq struct {
	UserId string `param:"userId" validate:"required"`
}

type DeleteNurseRes struct {
}

type LoginReq struct {
	Nip      int64  `json:"nip"      validate:"required,nip"`
	Password string `json:"password" validate:"required,min=5,max=33"`
}

type LoginRes struct {
	UserId      string `json:"userId"`
	Nip         int64  `json:"nip"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}

type ListUsersReq struct {
	UserId        *string `query:"userId"`
	Limit         *int    `query:"limit"`
	Offset        *int    `query:"offset"`
	Name          *string `query:"name"`
	Nip           *int64  `query:"nip"`
	Role          *string `query:"role"`
	CreatedAtSort *string `query:"createdAt"`
}

type ListUsersRes struct {
	Data []ListUsersResData
}

type ListUsersResData struct {
	UserId    string `json:"userId"`
	Nip       int64  `json:"nip"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type GetSessionFromTokenReq struct {
	AccessToken string
}

type GetSessionFromTokenRes struct {
	UserId string
	Nip    int64
}

type jwtSubClaims struct {
	UserId string `json:"userId"`
	Nip    int64  `json:"nip"`
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Data jwtSubClaims `json:"data"`
}
