package medicalrecord

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestListRecordValid(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		identityNumber := int64(1234123412341234)
		limit := 5
		offset := 0
		name := "epic"
		sort := "asc"
		phoneNumber := "62838"
		createdByNip := int64(1244123412341234)
		createdByUserId := "NIPNIPNIP"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodGet,
			"/v1/medical/record",
			rec,
			unittesting.WithQueryParams(map[string]string{
				"identityNumber":   fmt.Sprint(identityNumber),
				"name":             name,
				"createdAt":        sort,
				"phoneNumber":      phoneNumber,
				"createdBy.nip":    fmt.Sprint(createdByNip),
				"createdBy.userId": createdByUserId,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := medicalrecord.ListRecordReq{
				IdentityNumber:  &identityNumber,
				Name:            &name,
				PhoneNumber:     &phoneNumber,
				CreatedAtSort:   &sort,
				Limit:           &limit,
				Offset:          &offset,
				CreatedByNip:    &createdByNip,
				CreatedByUserId: &createdByUserId,
			}

			expectedRes := []medicalrecord.ListRecordResData{
				{
					IdentityDetail: medicalrecord.RecordIdentityDetail{
						IdentityNumber:      identityNumber,
						PhoneNumber:         phoneNumber,
						Name:                name,
						BirthDate:           "1969-06-09",
						Gender:              "female",
						IdentityCardScanImg: "https://someimage.com",
					},
					Symptoms:    "batuk gan",
					Medications: "obat batuk lah",
					CreatedAt:   "2024-06-09",
					CreatedBy: medicalrecord.RecordCreatedByDetail{
						Nip:    createdByNip,
						UserId: createdByUserId,
						Name:   "juan nip nip",
					},
				},
			}

			Convey("When the result data is not empty", func() {
				service.
					EXPECT().
					ListRecord(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							_ medicalrecord.ListRecordReq,
							res *medicalrecord.ListRecordRes,
						) {
							res.Data = expectedRes
						},
					).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.listRecord)

				Convey(
					"Should return HTTP 200 and the resulting array",
					func() {
						So(rec.Code, ShouldEqual, http.StatusOK)

						expectedBody := helper.MustMarshalJson(
							map[string]interface{}{
								"message": "success",
								"data":    expectedRes,
							},
						)

						So(
							rec.Body.String(),
							ShouldEqualJSON,
							string(expectedBody),
						)
					},
				)
			})

			Convey("When the result data is empty", func() {
				service.
					EXPECT().
					ListRecord(gomock.Any(), expectedReq, gomock.Any()).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.listRecord)

				Convey(
					"Should return HTTP 200 and an empty array",
					func() {
						So(rec.Code, ShouldEqual, http.StatusOK)

						expectedBody := helper.MustMarshalJson(
							map[string]interface{}{
								"message": "success",
								"data":    make([]struct{}, 0),
							},
						)

						So(
							rec.Body.String(),
							ShouldEqualJSON,
							string(expectedBody),
						)
					},
				)
			})
		})
	})
}
