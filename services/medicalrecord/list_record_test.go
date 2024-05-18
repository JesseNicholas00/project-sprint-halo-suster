package medicalrecord

import (
	"context"
	"testing"
	"time"

	"github.com/JesseNicholas00/HaloSuster/repos/auth"
	"github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListRecord(t *testing.T) {
	Convey("When called", t, func() {
		mockCtrl, service, mockedRepo, mockedAuthRepo := NewWithMockedRepo(t)
		defer mockCtrl.Finish()

		limit := 5
		offset := 0
		name := "epic"
		phoneNumber := "62838"
		createdAtSort := "asc"

		req := ListRecordReq{
			IdentityNumber:  nil,
			CreatedByNip:    nil,
			CreatedByUserId: nil,
			Limit:           &limit,
			Offset:          &offset,
			Name:            &name,
			PhoneNumber:     &phoneNumber,
			CreatedAtSort:   &createdAtSort,
		}

		var res ListRecordRes

		Convey("Should return the results given by the repo layer", func() {

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
			users := []auth.User{}
			for i := range dummyIds {
				users = append(users, auth.User{
					Id:       dummyIds[i],
					Nip:      dummyNips[i],
					Name:     "firstname lastname",
					Password: "hashedPasswordVeryScure",
					Admin:    true,
					Active:   true,
					ImageUrl: "https://bread.com/bread.png",
				})
			}

			record1 := medicalrecord.Record{
				RecordId:           123456789,
				PatientId:          987654321,
				PatientPhoneNumber: "+1234567890",
				PatientName:        "John Doe",
				PatientBirthDate:   time.Date(1990, 5, 15, 0, 0, 0, 0, time.UTC),
				PatientGender:      "Male",
				PatientImageUrl:    "https://example.com/johndoe.jpg",
				Symptoms:           "Fever, headache",
				Medications:        "Paracetamol",
				CreatedAt:          time.Date(2024, 5, 18, 10, 0, 0, 0, time.UTC),
				CreatedByUserId:    users[0].Id,
			}

			record2 := medicalrecord.Record{
				RecordId:           987654321,
				PatientId:          123456789,
				PatientPhoneNumber: "+1987654321",
				PatientName:        "Jane Smith",
				PatientBirthDate:   time.Date(1985, 10, 20, 0, 0, 0, 0, time.UTC),
				PatientGender:      "Female",
				PatientImageUrl:    "https://example.com/janesmith.jpg",
				Symptoms:           "Cough, sore throat",
				Medications:        "Cough syrup",
				CreatedAt:          time.Date(2024, 5, 17, 15, 30, 0, 0, time.UTC),
				CreatedByUserId:    users[1].Id,
			}
			records := []medicalrecord.Record{
				record1, record2,
			}

			mockedRepo.
				EXPECT().
				ListRecord(
					context.TODO(),
					medicalrecord.RecordFilter{
						IdentityNumber:  nil,
						CreatedByUserId: nil,
						Limit:           limit,
						Offset:          offset,
						CreatedAtSort:   &createdAtSort,
					},
				).
				Return(records, nil).
				Times(1)

			mockedAuthRepo.
				EXPECT().
				ListAllUsers(
					context.TODO(),
					auth.AllUsersFilter{
						UserIds: []string{record1.CreatedByUserId, record2.CreatedByUserId},
					}).Return(users, nil).
				Times(1)

			err := service.ListRecord(context.TODO(), req, &res)
			So(err, ShouldBeNil)

			So(res.Data, ShouldHaveLength, 2)
			for idx, record := range records {
				cur := res.Data[idx]
				So(
					cur.CreatedBy.UserId,
					ShouldEqual,
					record.CreatedByUserId,
				)
				So(
					cur.CreatedAt,
					ShouldEqual,
					record.CreatedAt.Format(time.RFC3339),
				)
				So(cur.Medications, ShouldEqual, record.Medications)
				So(cur.Symptoms, ShouldEqual, record.Symptoms)
				So(cur.IdentityDetail.PhoneNumber, ShouldEqual, record.PatientPhoneNumber)
				So(cur.IdentityDetail.Name, ShouldEqual, record.PatientName)
				So(cur.IdentityDetail.Gender, ShouldEqual, record.PatientGender)
				So(cur.IdentityDetail.BirthDate, ShouldEqual, record.PatientBirthDate.Format(time.DateOnly))
				So(cur.IdentityDetail.IdentityNumber, ShouldEqual, record.PatientId)
				So(cur.IdentityDetail.IdentityCardScanImg, ShouldEqual, record.PatientImageUrl)
				So(cur.IdentityDetail.IdentityCardScanImg, ShouldEqual, record.PatientImageUrl)
				So(cur.CreatedBy.UserId, ShouldEqual, record.CreatedByUserId)
				So(cur.CreatedBy.Nip, ShouldEqual, users[idx].Nip)
				So(cur.CreatedBy.Name, ShouldEqual, users[idx].Name)
			}
		})
	})
}
