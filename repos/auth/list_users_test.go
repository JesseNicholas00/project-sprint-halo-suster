//go:build integration
// +build integration

package auth_test

import (
	"context"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListUsersNipFilter(t *testing.T) {
	Convey(
		"When database contains users with different nips",
		t,
		func() {
			repo := NewWithTestDatabase(t)

			dummyIds := []string{
				"id1",
				"id2",
				"id3",
				"id4",
			}
			dummyNips := []int64{
				nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
				nip.New(nip.RoleIt, nip.GenderMale, 2012, 8, 36112),
				nip.New(nip.RoleIt, nip.GenderFemale, 1999, 12, 3612),
				nip.New(nip.RoleNurse, nip.GenderFemale, 2001, 2, 69),
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
				"Should return the users with a matching NIP prefix",
				func() {
					res, err := repo.ListUsers(
						context.TODO(),
						auth.UserFilter{
							Limit:  10,
							Offset: 0,
							Nip:    helper.ToPointer(int64(615120)),
						},
					)
					So(err, ShouldBeNil)
					So(res, ShouldHaveLength, 2)

					ids := []string{}
					for _, user := range res {
						ids = append(ids, user.Id)
					}

					So(ids, ShouldContain, dummyIds[0])
					So(ids, ShouldContain, dummyIds[1])
				},
			)
		},
	)
}
