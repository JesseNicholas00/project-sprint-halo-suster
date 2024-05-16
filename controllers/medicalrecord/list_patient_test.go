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

func TestListPatientValid(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		identityNumber := int64(1234123412341234)
		limit := 5
		offset := 0
		name := "epic"
		phoneNumber := "62838"
		createdAtSort := "wrong"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodGet,
			"/v1/medical/patient",
			rec,
			unittesting.WithQueryParams(map[string]string{
				"identityNumber": fmt.Sprint(identityNumber),
				"name":           name,
				"phoneNumber":    phoneNumber,
				"createdAt":      createdAtSort,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := medicalrecord.ListPatientsReq{
				IdentityNumber: &identityNumber,
				Name:           &name,
				PhoneNumber:    &phoneNumber,
				CreatedAtSort:  nil,
				Limit:          &limit,
				Offset:         &offset,
			}

			expectedRes := []medicalrecord.ListPatientsResData{
				{
					IdentityNumber: identityNumber,
					Name:           name,
					PhoneNumber:    phoneNumber,
					BirthDate:      "1969-06-09",
					Gender:         "female",
					CreatedAt:      "now",
				},
			}

			Convey("When the result data is not empty", func() {
				service.
					EXPECT().
					ListPatients(gomock.Any(), expectedReq, gomock.Any()).
					Do(
						func(
							_ context.Context,
							_ medicalrecord.ListPatientsReq,
							res *medicalrecord.ListPatientsRes,
						) {
							res.Data = expectedRes
						},
					).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.listPatients)

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
					ListPatients(gomock.Any(), expectedReq, gomock.Any()).
					Return(nil).
					Times(1)

				unittesting.CallController(ctx, controller.listPatients)

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
