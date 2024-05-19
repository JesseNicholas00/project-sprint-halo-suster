package medicalrecord

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRecordValid(t *testing.T) {
	Convey("When given a valid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		session := auth.GetSessionFromTokenRes{
			UserId: "bread",
			Nip:    nip.New(nip.RoleNurse, nip.GenderMale, 2024, 12, 420),
		}
		identityNumber := int64(1234567812345678)
		symptoms := "not enough rizz"
		medications := "fanum tax"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/medical/record",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"identityNumber": identityNumber,
				"symptoms":       symptoms,
				"medications":    medications,
			}),
			unittesting.WithContextData("session", session),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := medicalrecord.CreateRecordReq{
				CreatedById:    session.UserId,
				IdentityNumber: identityNumber,
				Symptoms:       symptoms,
				Medications:    medications,
			}

			service.
				EXPECT().
				CreateRecord(gomock.Any(), expectedReq, gomock.Any()).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.createRecord)

			Convey("Should return HTTP 201", func() {
				So(rec.Code, ShouldEqual, http.StatusCreated)
			})
		})
	})
}

func TestCreateRecordInvalid(t *testing.T) {
	Convey("When given an invalid request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		session := auth.GetSessionFromTokenRes{
			UserId: "bread",
			Nip:    nip.New(nip.RoleNurse, nip.GenderMale, 2024, 12, 420),
		}
		identityNumber := int64(1234567812345678)
		symptoms := "not enough rizz"
		medications := "fanum tax"

		Convey("On invalid request", func() {
			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/medical/record",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"identityNumber": identityNumber,
					"symptoms":       "", // empty symptom
					"medications":    medications,
				}),
				unittesting.WithContextData("session", session),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.createRecord)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On missing patient identity number", func() {
			expectedReq := medicalrecord.CreateRecordReq{
				CreatedById:    session.UserId,
				IdentityNumber: identityNumber,
				Symptoms:       symptoms,
				Medications:    medications,
			}

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/medical/record",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"identityNumber": identityNumber,
					"symptoms":       symptoms,
					"medications":    medications,
				}),
				unittesting.WithContextData("session", session),
			)

			service.
				EXPECT().
				CreateRecord(gomock.Any(), expectedReq, gomock.Any()).
				Return(medicalrecord.ErrIdentityNumberNotFound).
				Times(1)

			Convey("Should return HTTP code 404", func() {
				unittesting.CallController(ctx, controller.createRecord)
				So(rec.Code, ShouldEqual, http.StatusNotFound)
			})
		})
	})
}
