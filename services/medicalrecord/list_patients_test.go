package medicalrecord

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListPatients(t *testing.T) {
	Convey("When called", t, func() {
		mockCtrl, service, mockedRepo, _ := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		limit := 5
		offset := 0
		name := "epic"
		phoneNumber := "62838"
		createdAtSort := "asc"

		req := ListPatientsReq{
			IdentityNumber: nil,
			Limit:          &limit,
			Offset:         &offset,
			Name:           &name,
			PhoneNumber:    &phoneNumber,
			CreatedAtSort:  &createdAtSort,
		}

		var res ListPatientsRes

		Convey("Should return the results given by the repo layer", func() {
			dummyTime := helper.MustParseDateOnly("2022-02-02")
			patients := []medicalrecord.Patient{
				{
					IdentityNumber: int64(1231231231231234),
					PhoneNumber:    phoneNumber + "123",
					Name:           "prefix" + name,
					BirthDate:      dummyTime,
					Gender:         "male",
					ImageUrl:       "https://bro.com/bro.png",
					CreatedAt:      dummyTime,
				},
				{
					IdentityNumber: int64(1231231231231234),
					PhoneNumber:    "123" + phoneNumber,
					Name:           name + "suffix",
					BirthDate:      dummyTime,
					Gender:         "female",
					ImageUrl:       "https://bro.com/bro.png",
					CreatedAt:      dummyTime,
				},
			}

			mockedRepo.
				EXPECT().
				ListPatients(
					context.TODO(),
					medicalrecord.PatientFilter{
						IdentityNumber: nil,
						Limit:          limit,
						Offset:         offset,
						Name:           &name,
						PhoneNumber:    &phoneNumber,
						CreatedAtSort:  &createdAtSort,
					},
				).
				Return(patients, nil).
				Times(1)

			err := service.ListPatients(context.TODO(), req, &res)
			So(err, ShouldBeNil)

			So(res.Data, ShouldHaveLength, 2)
			for idx, patient := range patients {
				cur := res.Data[idx]
				So(
					helper.MustParseDateOnly(cur.BirthDate),
					ShouldEqual,
					patient.BirthDate,
				)
				So(
					cur.CreatedAt,
					ShouldEqual,
					patient.CreatedAt.Format(time.RFC3339),
				)
				So(cur.Gender, ShouldEqual, patient.Gender)
				So(cur.IdentityNumber, ShouldEqual, patient.IdentityNumber)
				So(cur.Name, ShouldEqual, patient.Name)
				So(cur.PhoneNumber, ShouldEqual, patient.PhoneNumber)
			}
		})
	})
}
