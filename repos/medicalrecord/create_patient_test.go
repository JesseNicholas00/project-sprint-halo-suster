//go:build integration
// +build integration

package medicalrecord_test

import (
	"context"
	"errors"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatePatient(t *testing.T) {
	Convey("When inserting new patient", t, func() {
		repo := NewWithTestDatabase(t)

		birthDate := helper.MustParseDateOnly("1998-01-01")

		req := medicalrecord.Patient{
			IdentityNumber: int64(1234567812345678),
			PhoneNumber:    "+621234567890",
			BirthDate:      birthDate,
			Gender:         "female",
			ImageUrl:       "https://bread.com/bread.png",
		}

		err := repo.CreatePatient(context.TODO(), req)
		Convey("Should return nil", func() {
			So(err, ShouldBeNil)
		})

		Convey("When inserting patient with duplicate identityNumber", func() {
			birthDate := helper.MustParseDateOnly("1992-11-07")

			reqDupe := medicalrecord.Patient{
				IdentityNumber: req.IdentityNumber,
				PhoneNumber:    "+620987654321",
				BirthDate:      birthDate,
				Gender:         "male",
				ImageUrl:       "https://funny.com/xd.png",
			}
			err := repo.CreatePatient(context.TODO(), reqDupe)
			Convey("Should return ErrDuplicateIdentityNumber", func() {
				So(
					errors.Is(err, medicalrecord.ErrDuplicateIdentityNumber),
					ShouldBeTrue,
				)
			})
		})
	})
}
