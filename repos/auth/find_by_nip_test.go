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

func TestFindUserByNip(t *testing.T) {
	Convey(
		"When database contains users with different nips",
		t,
		func() {
			repo := NewWithTestDatabase(t)

			dummyIds := []string{
				"id1",
				"id2",
				"id3",
			}
			dummyNips := []int64{
				nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
				nip.New(nip.RoleNurse, nip.GenderFemale, 2001, 2, 69),
				nip.New(nip.RoleIt, nip.GenderFemale, 1999, 12, 361),
			}
			for i := range dummyIds {
				curReqUser := auth.User{
					Id:       dummyIds[i],
					Nip:      dummyNips[i],
					Name:     "firstname lastname",
					Password: "hashedPasswordVeryScure",
					Admin:    true,
					Active:   true,
					ImageUrl: "https://bread.com/bread.png",
				}
				_, err := repo.CreateUser(context.TODO(), curReqUser)
				So(err, ShouldBeNil)
			}

			Convey(
				"Should return the staff with the requested phone number if one exists",
				func() {
					for _, expectedNip := range dummyNips {
						resStaff, err := repo.FindUserByNip(
							context.TODO(),
							expectedNip,
						)
						So(err, ShouldBeNil)
						So(resStaff.Nip, ShouldEqual, expectedNip)
					}
				},
			)

			Convey(
				"Should return ErrPhoneNumberNotFound when phone number doesn't exist",
				func() {
					_, err := repo.FindUserByNip(
						context.TODO(),
						nip.New(nip.RoleNurse, nip.GenderMale, 1968, 3, 3),
					)
					So(
						errors.Is(err, auth.ErrNipNotFound),
						ShouldBeTrue,
					)
				},
			)
		},
	)
}
