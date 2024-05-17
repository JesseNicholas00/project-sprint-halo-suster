package auth

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	gomock "github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	Convey("When logging in", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := LoginReq{
			Nip:      nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
			Password: "password",
		}
		reqWrong := LoginReq{
			Nip:      req.Nip,
			Password: "epic bruh moment",
		}

		cryptedPw, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			service.bcryptCost,
		)
		So(err, ShouldBeNil)

		dummyTime := helper.MustParseDateOnly("2022-02-02")

		repoRes := auth.User{
			Id:        "bread",
			Nip:       req.Nip,
			Name:      "john bread",
			Password:  string(cryptedPw),
			Active:    true,
			ImageUrl:  "https://bread.com/bread.png",
			CreatedAt: dummyTime,
		}

		repoResInactive := auth.User{
			Id:        "bread",
			Nip:       req.Nip,
			Name:      "john bread",
			Password:  string(cryptedPw),
			Active:    false,
			ImageUrl:  "https://bread.com/bread.png",
			CreatedAt: dummyTime,
		}

		Convey("If the user is inactive", func() {
			mockedRepo.EXPECT().
				FindUserByNip(gomock.Any(), req.Nip).
				Return(repoResInactive, nil).
				Times(1)

			res := LoginRes{}
			err := service.Login(context.TODO(), req, &res)
			Convey("Should return ErrUserHasNoAccess", func() {
				So(
					errors.Is(err, ErrUserHasNoAccess),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the NIP is not registered", func() {
			mockedRepo.EXPECT().
				FindUserByNip(gomock.Any(), req.Nip).
				Return(auth.User{}, auth.ErrNipNotFound).
				Times(1)

			res := LoginRes{}
			err := service.Login(context.TODO(), req, &res)
			Convey("Should return ErrUserNotFound", func() {
				So(
					errors.Is(err, ErrUserNotFound),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the NIP is registered", func() {
			mockedRepo.EXPECT().
				FindUserByNip(gomock.Any(), req.Nip).
				Return(repoRes, nil).
				Times(1)

			Convey(
				"And the password is incorrect",
				func() {
					res := LoginRes{}
					err := service.Login(context.TODO(), reqWrong, &res)

					Convey("Should return ErrInvalidCredentials", func() {
						So(errors.Is(err, ErrInvalidCredentials), ShouldBeTrue)
					})
				},
			)

			Convey(
				"And the password is correct",
				func() {
					res := LoginRes{}
					err := service.Login(context.TODO(), req, &res)

					Convey(
						"Should return nil and write the correct result to res",
						func() {
							So(err, ShouldBeNil)
							So(res.UserId, ShouldEqual, repoRes.Id)
							So(res.Nip, ShouldEqual, repoRes.Nip)
							So(res.Name, ShouldEqual, repoRes.Name)
						},
					)
				},
			)
		})
	})
}
