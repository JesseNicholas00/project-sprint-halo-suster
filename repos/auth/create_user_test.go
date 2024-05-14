//go:build integration
// +build integration

package auth_test

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateUser(t *testing.T) {
	Convey("When inserting new user from parameter", t, func() {
		repo := NewWithTestDatabase(t)

		reqUser := auth.User{
			Id:       "testId",
			Nip:      nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
			Name:     "firstname lastname",
			Password: "hashedPasswordVerySecure",
			Active:   true,
			ImageUrl: "https://bread.com/bread.png",
		}

		resUser, err := repo.CreateUser(context.TODO(), reqUser)
		Convey("Should return the created user with the same data", func() {
			So(err, ShouldBeNil)
			So(resUser.Id, ShouldEqual, reqUser.Id)
			So(resUser.Nip, ShouldEqual, reqUser.Nip)
			So(resUser.Name, ShouldEqual, reqUser.Name)
			So(resUser.Password, ShouldEqual, reqUser.Password)
			So(resUser.Active, ShouldEqual, reqUser.Active)
			So(resUser.ImageUrl, ShouldEqual, reqUser.ImageUrl)
		})

		Convey("When inserting duplicate user", func() {
			reqUser.Nip = nip.New(nip.RoleNurse, nip.GenderFemale, 2001, 1, 420)
			_, err := repo.CreateUser(context.TODO(), reqUser)
			Convey("Should return duplicate error", func() {
				So(errors.Is(err, auth.ErrDuplicateUser), ShouldBeTrue)
			})
		})
	})
}
