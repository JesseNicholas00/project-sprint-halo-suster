package medicalrecord

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRecord(t *testing.T) {
	mockCtrl, service, mockedRepo, _ := NewWithMockedRepo(t)
	defer mockCtrl.Finish()

	identityNumber := int64(1234567812345678)
	phoneNumber := "+62123456892"
	name := "firstname lastname"
	birthDate := "1997-03-24T16:02:22.011Z"
	gender := "male"
	identityCardImg := "https://bread.com/bread.png"

	req := RegisterPatientReq{
		IdentityNumber:  identityNumber,
		PhoneNumber:     phoneNumber,
		Name:            name,
		BirthDate:       birthDate,
		Gender:          gender,
		IdentityCardImg: identityCardImg,
	}

	expectedRepoReq := medicalrecord.Patient{
		IdentityNumber: identityNumber,
		PhoneNumber:    phoneNumber,
		Name:           name,
		BirthDate:      helper.MustParse(birthDate),
		Gender:         gender,
		ImageUrl:       identityCardImg,
	}

	Convey("When patient is successfully created", t, func() {
		mockedRepo.
			EXPECT().
			CreatePatient(gomock.Any(), expectedRepoReq).
			Return(nil).
			Times(1)

		err := service.RegisterPatient(
			context.TODO(),
			req,
			&RegisterPatientRes{},
		)

		Convey("Should not return an error", func() {
			So(err, ShouldBeNil)
		})
	})

	Convey("When patient has duplicate identity number", t, func() {
		mockedRepo.
			EXPECT().
			CreatePatient(gomock.Any(), expectedRepoReq).
			Return(medicalrecord.ErrDuplicateIdentityNumber).
			Times(1)

		err := service.RegisterPatient(
			context.TODO(),
			req,
			&RegisterPatientRes{},
		)

		Convey("Should return ErrDuplicateIdentityNumber", func() {
			So(errors.Is(err, ErrDuplicateIdentityNumber), ShouldBeTrue)
		})
	})
}
