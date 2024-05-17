package medicalrecord

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterPatientValid(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		identityNumber := int64(1234567812345678)
		phoneNumber := "+62123456892"
		name := "firstname lastname"
		birthDate := "1997-03-24T16:02:22.011Z"
		gender := "male"
		identityCardImg := "https://bread.com/bread.png"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/medical/patient",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"identityNumber":      identityNumber,
				"phoneNumber":         phoneNumber,
				"name":                name,
				"birthDate":           birthDate,
				"gender":              gender,
				"identityCardScanImg": identityCardImg,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := medicalrecord.RegisterPatientReq{
				IdentityNumber:  identityNumber,
				PhoneNumber:     phoneNumber,
				Name:            name,
				BirthDate:       birthDate,
				Gender:          gender,
				IdentityCardImg: identityCardImg,
			}

			service.
				EXPECT().
				RegisterPatient(gomock.Any(), expectedReq, gomock.Any()).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerPatient)

			Convey("Should return HTTP 201", func() {
				So(rec.Code, ShouldEqual, http.StatusCreated)
			})
		})
	})
}

func TestRegisterPatientInvalid(t *testing.T) {
	Convey("When given an invalid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		identityNumber := int64(1234567812345678)
		phoneNumber := "+62123456892"
		name := "firstname lastname"
		birthDate := "1986-01-01T00:00:00Z"
		gender := "male"
		identityCardImg := "https://bread.com/bread.png"

		Convey("On invalid request", func() {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/medical/patient",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					// id too long
					"identityNumber":      identityNumber*10 + 2,
					"phoneNumber":         phoneNumber,
					"name":                name,
					"birthDate":           birthDate,
					"gender":              gender,
					"identityCardScanImg": identityCardImg,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.registerPatient)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On duplicate identity number", func() {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/medical/patient",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"identityNumber":      identityNumber,
					"phoneNumber":         phoneNumber,
					"name":                name,
					"birthDate":           birthDate,
					"gender":              gender,
					"identityCardScanImg": identityCardImg,
				}),
			)

			expectedReq := medicalrecord.RegisterPatientReq{
				IdentityNumber:  identityNumber,
				PhoneNumber:     phoneNumber,
				Name:            name,
				BirthDate:       birthDate,
				Gender:          gender,
				IdentityCardImg: identityCardImg,
			}

			service.
				EXPECT().
				RegisterPatient(gomock.Any(), expectedReq, gomock.Any()).
				Return(medicalrecord.ErrDuplicateIdentityNumber).
				Times(1)

			unittesting.CallController(ctx, controller.registerPatient)

			Convey("Should return HTTP 409", func() {
				So(rec.Code, ShouldEqual, http.StatusConflict)
			})
		})
	})
}
