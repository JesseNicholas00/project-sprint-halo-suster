package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	"github.com/JesseNicholas00/HaloSuster/utils/unittesting"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRegisterItValid(t *testing.T) {
	Convey("When given a valid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		userId := "dummyId"
		name := "namadepan namabelakang"
		nip := nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 429)
		password := "password"
		accessToken := "token"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/user/it/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"nip":      nip,
				"name":     name,
				"password": password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.RegisterItReq{
				Nip:      nip,
				Name:     name,
				Password: password,
			}
			expectedRes := auth.RegisterItRes{
				UserId:      userId,
				Nip:         nip,
				Name:        name,
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				RegisterIt(gomock.Any(), expectedReq, gomock.Any()).
				Do(func(_ context.Context, _ auth.RegisterItReq, res *auth.RegisterItRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerIt)

			Convey(
				"Should return the expected response with HTTP 201",
				func() {
					So(rec.Code, ShouldEqual, http.StatusCreated)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "User registered successfully",
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
	})
}

func TestRegisterInvalid(t *testing.T) {
	Convey("When given an invalid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		name := "firstname lastname"
		password := "password"

		Convey("On invalid request", func() {
			// nurse NIP
			nip := nip.New(nip.RoleNurse, nip.GenderMale, 2001, 1, 429)

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/it/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"name":     name,
					"nip":      nip,
					"password": password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.registerIt)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On duplicate NIP", func() {
			nip := nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 429)

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/it/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"name":     name,
					"nip":      nip,
					"password": password,
				}),
			)

			Convey("Should return HTTP code 409", func() {
				expectedReq := auth.RegisterItReq{
					Nip:      nip,
					Name:     name,
					Password: password,
				}

				service.EXPECT().
					RegisterIt(gomock.Any(), expectedReq, gomock.Any()).
					Return(auth.ErrNipAlreadyExists).
					Times(1)

				unittesting.CallController(ctx, controller.registerIt)
				So(rec.Code, ShouldEqual, http.StatusConflict)
			})
		})
	})
}
