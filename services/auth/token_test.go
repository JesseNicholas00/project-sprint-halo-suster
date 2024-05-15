package auth

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGenerateToken(t *testing.T) {
	Convey("When generating token from user", t, func() {
		mockCtrl, service, _ := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		user := auth.User{
			Id:  "bread",
			Nip: nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
		}

		token, err := service.generateToken(user)
		Convey(
			"Should return a token containing user data without errors",
			func() {
				So(err, ShouldBeNil)

				claims := jwtClaims{}
				_, err := jwt.ParseWithClaims(
					token,
					&claims,
					func(t *jwt.Token) (interface{}, error) {
						return service.jwtSecret, nil
					},
				)

				So(err, ShouldBeNil)
				So(claims.Data.UserId, ShouldEqual, user.Id)
				So(claims.Data.Nip, ShouldEqual, user.Nip)
			},
		)
	})
}
