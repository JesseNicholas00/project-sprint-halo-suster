package main

import (
	"fmt"

	"github.com/JesseNicholas00/HaloSuster/middlewares"
	"github.com/JesseNicholas00/HaloSuster/utils/logging"
	"github.com/JesseNicholas00/HaloSuster/utils/migration"
	"github.com/JesseNicholas00/HaloSuster/utils/validation"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func main() {
	mainInitLogger := logging.GetLogger("main", "init")

	cfg, err := loadConfig()
	if err != nil {
		mainInitLogger.Error("%s", err)
	}

	logging.SetLogLevel(cfg.logLevel)

	mainInitLogger.Debug(fmt.Sprintf("%+v", cfg))

	if cfg.migrateDownOnStart {
		if err := migration.MigrateDown(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Error("failed to migrate down db: %s", err)
			return
		}
	}
	if cfg.migrateUpOnStart {
		if err := migration.MigrateUp(cfg.dbString, "migrations"); err != nil {
			mainInitLogger.Error("failed to migrate up db: %s", err)
			return
		}
	}

	db, err := sqlx.Connect("postgres", cfg.dbString)
	if err != nil {
		mainInitLogger.Error("%s", err)
		return
	}

	db.SetMaxOpenConns(cfg.dbMaxOpenConns)
	db.SetMaxIdleConns(cfg.dbMaxIdleConns)
	db.SetConnMaxLifetime(cfg.dbMaxConnLifetime)

	defer db.Close()

	controllers := initControllers(cfg, db)

	server := echo.New()

	if cfg.traceSlowEndpoints {
		slowLogger := middlewares.NewSlowTracerMiddleware(cfg.slowThreshold)
		server.Use(slowLogger.Process)
	}

	errorHandler := middlewares.NewLoggingErrorHandlerMiddleware()
	server.Use(errorHandler.Process)

	for idx, controller := range controllers {
		if err := controller.Register(server); err != nil {
			msg := fmt.Sprintf(
				"failed during controller registration (%d/%d): %s",
				idx+1,
				len(controllers),
				err,
			)
			mainInitLogger.Error(msg)
			return
		}
	}

	server.Validator = validation.NewEchoValidator()
	server.HideBanner = true

	server.Logger.Fatal(
		server.Start(
			fmt.Sprintf(
				"%s:%d",
				cfg.serverHost,
				cfg.serverPort,
			),
		),
	)
}
