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

func TestRegisterNurseValid(t *testing.T) {
	Convey("When given a valid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		userId := "dummyId"
		name := "namadepan namabelakang"
		nip := nip.New(nip.RoleNurse, nip.GenderMale, 2001, 1, 420)
		imageUrl := "https://awss3.d87801e9-fcfc-42a8-963b-fe86d895b51a.jpeg"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/user/nurse/register",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"nip":                 nip,
				"name":                name,
				"identityCardScanImg": imageUrl,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.RegisterNurseReq{
				Nip:      nip,
				Name:     name,
				ImageUrl: imageUrl,
			}
			expectedRes := auth.RegisterNurseRes{
				UserId: userId,
				Nip:    nip,
				Name:   name,
			}

			service.
				EXPECT().
				RegisterNurse(gomock.Any(), expectedReq, gomock.Any()).
				Do(func(_ context.Context, _ auth.RegisterNurseReq, res *auth.RegisterNurseRes) {
					*res = expectedRes
				}).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.registerNurse)

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

func TestRegisterNurseInvalid(t *testing.T) {
	Convey("When given an invalid register request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		name := "firstname lastname"
		imageUrl := "https://awss3.d87801e9-fcfc-42a8-963b-fe86d895b51a.jpeg"

		Convey("On invalid request", func() {
			// IT NIP
			nip := nip.New(nip.RoleIt, nip.GenderMale, 2001, 1, 429)

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/nurse/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"nip":                 nip,
					"name":                name,
					"identityCardScanImg": imageUrl,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.registerNurse)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On duplicate NIP", func() {
			nip := nip.New(nip.RoleNurse, nip.GenderMale, 2001, 1, 429)

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/nurse/register",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"nip":                 nip,
					"name":                name,
					"identityCardScanImg": imageUrl,
				}),
			)

			Convey("Should return HTTP code 409", func() {
				expectedReq := auth.RegisterNurseReq{
					Nip:      nip,
					Name:     name,
					ImageUrl: imageUrl,
				}

				service.EXPECT().
					RegisterNurse(gomock.Any(), expectedReq, gomock.Any()).
					Return(auth.ErrNipAlreadyExists).
					Times(1)

				unittesting.CallController(ctx, controller.registerNurse)
				So(rec.Code, ShouldEqual, http.StatusConflict)
			})
		})
	})
}
