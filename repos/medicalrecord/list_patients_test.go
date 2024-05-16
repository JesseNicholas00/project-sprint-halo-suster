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

func TestListPatients(t *testing.T) {
	Convey("With dummy data", t, func() {
		repo := NewWithTestDatabase(t)

		birthDate := helper.MustParseDateOnly("1998-01-01")

		names := []string{
			"rizz gyatt fubuki",
			"roti jenna tools bruh",
			"gyatt turbo joseph mewing",
		}
		identityNumbers := []int64{
			1234561234561234,
			1231231231237771,
			2893877839281939,
		}
		phoneNumbers := []string{
			"+62123123123123",
			"+62321321321321",
			"+62323232323232",
		}
		genders := []string{
			"male",
			"female",
			"male",
		}

		var patients []medicalrecord.Patient
		for i := 0; i < 3; i++ {
			patients = append(
				patients,
				medicalrecord.Patient{
					Name:           names[i],
					IdentityNumber: identityNumbers[i],
					PhoneNumber:    phoneNumbers[i],
					BirthDate:      birthDate,
					Gender:         genders[i],
					ImageUrl:       "https://bread.com/bread.png",
				},
			)
		}

		for _, patient := range patients {
			err := repo.CreatePatient(context.TODO(), patient)
			So(err, ShouldBeNil)
		}

		Convey("When querying with phoneNumber filter", func() {
			Convey("Should return the matching users only", func() {
				req := medicalrecord.PatientFilter{
					PhoneNumber: helper.ToPointer("32"),
					Limit:       5,
					Offset:      0,
				}

				res, err := repo.ListPatients(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 2)

				var returnedPhoneNumbers []string
				for _, patient := range res {
					returnedPhoneNumbers = append(
						returnedPhoneNumbers,
						patient.PhoneNumber,
					)
				}

				So(returnedPhoneNumbers, ShouldContain, phoneNumbers[1])
				So(returnedPhoneNumbers, ShouldContain, phoneNumbers[2])
			})
		})

		Convey("When querying with name filter", func() {
			Convey("Should return the matching users only", func() {
				req := medicalrecord.PatientFilter{
					Name:   helper.ToPointer("gyatt"),
					Limit:  5,
					Offset: 0,
				}

				res, err := repo.ListPatients(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 2)

				var returnedNames []string
				for _, patient := range res {
					returnedNames = append(returnedNames, patient.Name)
				}

				So(returnedNames, ShouldContain, names[0])
				So(returnedNames, ShouldContain, names[2])
			})
		})

		Convey("When querying with identityNumber filter", func() {
			Convey("When user exists", func() {
				Convey("Should successfully get only the specific user",
					func() {
						for idx, identityNumber := range identityNumbers {
							req := medicalrecord.PatientFilter{
								IdentityNumber: &identityNumber,
								Limit:          5,
								Offset:         0,
							}

							res, err := repo.ListPatients(context.TODO(), req)
							So(err, ShouldBeNil)
							So(res, ShouldHaveLength, 1)

							got := res[0]
							exp := patients[idx]

							So(got.IdentityNumber, ShouldEqual, identityNumber)

							So(got.BirthDate, ShouldEqual, exp.BirthDate)
							So(got.Gender, ShouldEqual, exp.Gender)
							So(got.ImageUrl, ShouldEqual, exp.ImageUrl)
							So(got.Name, ShouldEqual, exp.Name)
							So(got.PhoneNumber, ShouldEqual, exp.PhoneNumber)
						}
					},
				)
			})
			Convey("When user doesn't exist", func() {
				Convey("Should return an empty list", func() {
					req := medicalrecord.PatientFilter{
						IdentityNumber: helper.ToPointer(
							int64(6969696969696969),
						),
						Limit:  5,
						Offset: 0,
					}

					res, err := repo.ListPatients(context.TODO(), req)
					So(err, ShouldBeNil)
					So(res, ShouldBeEmpty)
				})
			})
		})
	})
}
