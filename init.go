package main

import (
	"github.com/JesseNicholas00/HaloSuster/controllers"
	authCtrl "github.com/JesseNicholas00/HaloSuster/controllers/auth"
	authRepo "github.com/JesseNicholas00/HaloSuster/repos/auth"
	authSvc "github.com/JesseNicholas00/HaloSuster/services/auth"
	"github.com/JesseNicholas00/HaloSuster/utils/ctxrizz"
	"github.com/JesseNicholas00/HaloSuster/utils/logging"
	"github.com/jmoiron/sqlx"
)

func initControllers(
	cfg ServerConfig,
	db *sqlx.DB,
) (ctrls []controllers.Controller) {
	ctrlInitLogger := logging.GetLogger("main", "init", "controllers")
	defer func() {
		if r := recover(); r != nil {
			// add extra context to help debug potential panic
			ctrlInitLogger.Error("panic while initializing controllers: %s", r)
			panic(r)
		}
	}()

	dbRizzer := ctxrizz.NewDbContextRizzer(db)

	// withTxMw := middlewares.NewWithTxMiddleware(dbRizzer)

	authRepo := authRepo.NewAuthRepository(dbRizzer)
	authSvc := authSvc.NewAuthService(
		authRepo,
		cfg.jwtSecretKey,
		cfg.bcryptSaltCost,
	)
	authCtrl := authCtrl.NewAuthController(authSvc)
	// authCtrl := authCtrl.NewAuthController(authSvc, withTxMw)
	// authMw := middlewares.NewAuthMiddleware(authSvc)

	ctrls = append(ctrls, authCtrl)

	return
}
