//go:build integration
// +build integration

package medicalrecord_test

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListRecord(t *testing.T) {
	Convey("With dummy data", t, func() {
		repo := NewWithTestDatabase(t)

		record1 := medicalrecord.Record{
			RecordId:           1,
			PatientId:          123456789,
			PatientPhoneNumber: "+1234567890",
			PatientName:        "John Doe",
			PatientBirthDate:   time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC),
			PatientGender:      "Male",
			PatientImageUrl:    "http://example.com/image1.jpg",
			Symptoms:           "Fever, cough",
			Medications:        "Paracetamol",
			CreatedAt:          time.Now(),
			CreatedByUserId:    "user123",
		}

		record2 := medicalrecord.Record{
			RecordId:           2,
			PatientId:          987654321,
			PatientPhoneNumber: "+9876543210",
			PatientName:        "Jane Smith",
			PatientBirthDate:   time.Date(1990, time.February, 15, 0, 0, 0, 0, time.UTC),
			PatientGender:      "Female",
			PatientImageUrl:    "http://example.com/image2.jpg",
			Symptoms:           "Headache, fatigue",
			Medications:        "Aspirin",
			CreatedAt:          time.Now(),
			CreatedByUserId:    "user456",
		}

		record3 := medicalrecord.Record{
			RecordId:           3,
			PatientId:          555555555,
			PatientPhoneNumber: "+5555555555",
			PatientName:        "Alex Johnson",
			PatientBirthDate:   time.Date(1975, time.March, 10, 0, 0, 0, 0, time.UTC),
			PatientGender:      "Male",
			PatientImageUrl:    "http://example.com/image3.jpg",
			Symptoms:           "Nausea, vomiting",
			Medications:        "Antiemetic",
			CreatedAt:          time.Now(),
			CreatedByUserId:    "user123",
		}

		records := []medicalrecord.Record{record1, record2, record3}

		for _, record := range records {
			err := repo.CreateRecord(context.TODO(), record)
			So(err, ShouldBeNil)
		}

		Convey("When querying with phoneNumber filter", func() {
			Convey("Should return the matching users only", func() {
				req := medicalrecord.RecordFilter{
					IdentityNumber: helper.ToPointer(int64(555555555)),
					Limit:          5,
					Offset:         0,
				}

				res, err := repo.ListRecord(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 1)

				var returnedPhoneNumbers []string
				for _, patient := range res {
					returnedPhoneNumbers = append(
						returnedPhoneNumbers,
						patient.PatientPhoneNumber,
					)
				}

				So(returnedPhoneNumbers, ShouldContain, record3.PatientPhoneNumber)
			})
		})

		Convey("When querying with name filter", func() {
			Convey("Should return the matching users only", func() {
				req := medicalrecord.RecordFilter{
					CreatedByUserId: helper.ToPointer("user123"),
					Limit:           5,
					Offset:          0,
				}

				res, err := repo.ListRecord(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldHaveLength, 2)

				var returnedNames []string
				for _, patient := range res {
					returnedNames = append(returnedNames, patient.PatientName)
				}

				So(returnedNames, ShouldContain, record1.PatientName)
				So(returnedNames, ShouldContain, record3.PatientName)
			})
		})

		Convey("When user doesn't exist", func() {
			Convey("Should return an empty list", func() {
				req := medicalrecord.RecordFilter{
					IdentityNumber: helper.ToPointer(
						int64(6969696969696969),
					),
					Limit:  5,
					Offset: 0,
				}

				res, err := repo.ListRecord(context.TODO(), req)
				So(err, ShouldBeNil)
				So(res, ShouldBeEmpty)
			})
		})
	})
}
