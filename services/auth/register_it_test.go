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
)

func TestRegisterIt(t *testing.T) {
	Convey("When registering staff", t, func() {
		mockCtrl, service, mockedRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		req := RegisterItReq{
			Name:     "firstname lastname",
			Nip:      nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 420),
			Password: "password",
		}
		dummyTime := helper.MustParseDateOnly("2022-02-02")

		repoReq := auth.User{
			Id:       "bread",
			Nip:      req.Nip,
			Name:     req.Name,
			Active:   true,
			Password: req.Password,
		}
		repoRes := auth.User{
			Id:        repoReq.Id,
			Nip:       repoReq.Nip,
			Name:      repoReq.Name,
			Password:  repoReq.Password,
			Active:    repoReq.Active,
			ImageUrl:  repoReq.ImageUrl,
			CreatedAt: dummyTime,
		}

		Convey("If the NIP is already registered", func() {
			mockedRepo.EXPECT().
				FindUserByNip(gomock.Any(), req.Nip).
				Return(repoRes, nil).
				Times(1)

			res := RegisterItRes{}
			err := service.RegisterIt(context.TODO(), req, &res)
			Convey("Should return ErrPhoneNumberAlreadyRegistered", func() {
				So(
					errors.Is(err, ErrNipAlreadyExists),
					ShouldBeTrue,
				)
			})
		})

		Convey("If the NIP is unique", func() {
			mockedRepo.EXPECT().
				FindUserByNip(gomock.Any(), req.Nip).
				Return(auth.User{}, auth.ErrNipNotFound).
				Times(1)
			mockedRepo.EXPECT().
				CreateUser(gomock.Any(), gomock.Any()).
				Do(func(_ context.Context, reqFromSvc auth.User) {
					So(reqFromSvc.Nip, ShouldEqual, req.Nip)
					So(reqFromSvc.Name, ShouldEqual, req.Name)
					So(reqFromSvc.Active, ShouldBeTrue)
				}).
				Return(repoRes, nil).
				Times(1)

			res := RegisterItRes{}
			err := service.RegisterIt(context.TODO(), req, &res)
			Convey(
				"Should return nil and write the correct result to res",
				func() {
					So(err, ShouldBeNil)
					So(res.UserId, ShouldEqual, repoRes.Id)
					So(res.Nip, ShouldEqual, req.Nip)
					So(res.Name, ShouldEqual, req.Name)
				},
			)
		})
	})
}
