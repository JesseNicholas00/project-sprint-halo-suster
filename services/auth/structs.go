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
