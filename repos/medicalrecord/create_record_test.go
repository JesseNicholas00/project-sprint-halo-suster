//go:build integration
// +build integration

package medicalrecord_test

import (
	"context"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRecord(t *testing.T) {
	Convey("When inserting new record", t, func() {
		repo := NewWithTestDatabase(t)

		birthDate := helper.MustParseDateOnly("1998-01-01")

		req := medicalrecord.Record{
			PatientId:          int64(1231231212312312),
			PatientPhoneNumber: "+628312312312",
			PatientName:        "sussy baka",
			PatientBirthDate:   birthDate,
			PatientGender:      "female",
			Symptoms:           "no rizz",
			Medications:        "mewing 5x per week",
			CreatedByUserId:    "doctor gyatt",
		}

		err := repo.CreateRecord(context.TODO(), req)
		Convey("Should return nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("When inserting the same record again", func() {
			err := repo.CreateRecord(context.TODO(), req)
			Convey("Should still return nil (allow duplicates)", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}
