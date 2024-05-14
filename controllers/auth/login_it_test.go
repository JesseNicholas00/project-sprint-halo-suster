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

func TestLoginItValid(t *testing.T) {
	Convey("When given a valid login request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		userId := "dummyId"
		name := "namadepan namabelakang"
		nip := nip.New(nip.RoleIt, nip.GenderFemale, 2003, 1, 123)
		password := "password"
		accessToken := "token"

		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/user/it/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"nip":      nip,
				"password": password,
			}),
		)

		Convey("Should forward the request to the service layer", func() {
			expectedReq := auth.LoginReq{
				Nip:      nip,
				Password: password,
			}
			expectedRes := auth.LoginRes{
				UserId:      userId,
				Nip:         nip,
				Name:        name,
				AccessToken: accessToken,
			}

			service.
				EXPECT().
				Login(gomock.Any(), expectedReq, gomock.Any()).
				Do(
					func(
						_ context.Context,
						_ auth.LoginReq,
						res *auth.LoginRes,
					) {
						*res = expectedRes
					},
				).
				Return(nil).
				Times(1)

			unittesting.CallController(ctx, controller.loginIt)

			Convey(
				"Should return the expected response with HTTP 200",
				func() {
					So(rec.Code, ShouldEqual, http.StatusOK)

					expectedBody := helper.MustMarshalJson(
						map[string]interface{}{
							"message": "User logged in successfully",
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

func TestLoginItInvalid(t *testing.T) {
	Convey("When given an invalid login request", t, func() {
		mockCtrl, controller, service := NewControllerWithMockedService(t)
		defer mockCtrl.Finish()

		Convey("On bad request", func() {
			// wrong nip format
			nip := 19823673698126
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/it/login",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"nip":      nip,
					"password": password,
				}),
			)

			Convey("Should return HTTP code 400", func() {
				unittesting.CallController(ctx, controller.loginIt)
				So(rec.Code, ShouldEqual, http.StatusBadRequest)
			})
		})

		Convey("On not found", func() {
			// nurse nip
			nip := nip.New(nip.RoleNurse, nip.GenderFemale, 2000, 6, 999)
			password := "password"

			rec := httptest.NewRecorder()
			ctx := unittesting.CreateEchoContextFromRequest(
				http.MethodPost,
				"/v1/user/it/login",
				rec,
				unittesting.WithJsonPayload(map[string]interface{}{
					"nip":      nip,
					"password": password,
				}),
			)

			Convey("Should return HTTP code 404", func() {
				unittesting.CallController(ctx, controller.loginIt)
				So(rec.Code, ShouldEqual, http.StatusNotFound)
			})
		})

		nip := nip.New(nip.RoleIt, nip.GenderFemale, 2001, 6, 999)
		password := "password"
		rec := httptest.NewRecorder()
		ctx := unittesting.CreateEchoContextFromRequest(
			http.MethodPost,
			"/v1/user/it/login",
			rec,
			unittesting.WithJsonPayload(map[string]interface{}{
				"nip":      nip,
				"password": password,
			}),
		)

		expectedReq := auth.LoginReq{
			Nip:      nip,
			Password: password,
		}

		Convey("On user not found", func() {
			service.
				EXPECT().
				Login(gomock.Any(), expectedReq, gomock.Any()).
				Return(auth.ErrUserNotFound).
				Times(1)

			Convey(
				"Should return HTTP code 404",
				func() {
					unittesting.CallController(ctx, controller.loginIt)
					So(rec.Code, ShouldEqual, http.StatusNotFound)
				},
			)
		})

		Convey("On invalid credentials", func() {
			service.
				EXPECT().
				Login(gomock.Any(), expectedReq, gomock.Any()).
				Return(auth.ErrInvalidCredentials).
				Times(1)

			Convey(
				"Should return HTTP code 400",
				func() {
					unittesting.CallController(ctx, controller.loginIt)
					So(rec.Code, ShouldEqual, http.StatusBadRequest)
				},
			)
		})
	})
}
