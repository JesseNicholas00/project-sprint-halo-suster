package main

import (
	"fmt"

	"github.com/JesseNicholas00/HaloSuster/controllers"
	authCtrl "github.com/JesseNicholas00/HaloSuster/controllers/auth"
	imageCtrl "github.com/JesseNicholas00/HaloSuster/controllers/image"
	medicalRecordCtrl "github.com/JesseNicholas00/HaloSuster/controllers/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/middlewares"
	authRepo "github.com/JesseNicholas00/HaloSuster/repos/auth"
	medicalRecordRepo "github.com/JesseNicholas00/HaloSuster/repos/medicalrecord"
	authSvc "github.com/JesseNicholas00/HaloSuster/services/auth"
	medicalRecordSvc "github.com/JesseNicholas00/HaloSuster/services/medicalrecord"
	"github.com/JesseNicholas00/HaloSuster/types/nip"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
	"github.com/JesseNicholas00/HaloSuster/utils/logging"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/jmoiron/sqlx"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
	uploader *manager.Uploader,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Error(
				fmt.Sprintf("panic while initializing controllers: %s", r),
			)
			panic(r)
		}
	}()

	dbRizzer := ctxrizz.NewDbContextRizzer(db)

	// withTxMw := middlewares.NewWithTxMiddleware(dbRizzer)

	authRepo := authRepo.NewAuthRepository(dbRizzer)
	authSvc := authSvc.NewAuthService(
		authRepo,
		dbRizzer,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authMwIt := middlewares.NewAuthMiddleware(authSvc, nip.RoleIt)
	authCtrl := authCtrl.NewAuthController(authSvc, authMwIt)
	ctrls = append(ctrls, authCtrl)

	authMwEither := middlewares.NewAuthMiddleware(
		authSvc,
		nip.RoleIt,
		nip.RoleNurse,
	)

	medicalRecordRepo := medicalRecordRepo.NewMedicalRecordRepository(dbRizzer)
	medicalRecordSvc := medicalRecordSvc.NewMedicalRecordService(
		medicalRecordRepo,
		authRepo,
		dbRizzer,
	)
	medicalRecordCtrl := medicalRecordCtrl.NewMedicalRecordController(
		medicalRecordSvc,
		authMwEither,
	)
	ctrls = append(ctrls, medicalRecordCtrl)
	imageCtrl := imageCtrl.NewImageController(uploader, cfg.awsS3BucketName, authMwEither)
	ctrls = append(ctrls, imageCtrl)

	return
}
