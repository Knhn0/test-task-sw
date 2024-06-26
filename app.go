package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"test-task-sw/api"
	"test-task-sw/config"
	"test-task-sw/database/department"
	"test-task-sw/database/employee"
	"test-task-sw/database/passport"
	"test-task-sw/lib/tctx"
	"test-task-sw/lib/tpostgres"
	"test-task-sw/service"
)

type App struct {
	logger          *zap.SugaredLogger
	config          config.Config
	Server          *api.Server
	contextProvider tctx.DefaultContextProviderFunc
	notify          chan struct{}

	//db
	pgDb *tpostgres.Postgres

	//repos
	passportRepo   passport.Repo
	departmentRepo department.Repo
	employeeRepo   employee.Repo

	//services
	employeeService *service.EmployeeService
}

func NewApp(logger *zap.SugaredLogger, cfg config.Config, contextProvider tctx.DefaultContextProviderFunc) (*App, error) {
	app := &App{
		logger:          logger,
		config:          cfg,
		contextProvider: contextProvider,
		notify:          make(chan struct{}, 1),
	}

	if err := app.initDatabases(); err != nil {
		return app, err
	}

	if err := app.initServices(); err != nil {
		return app, err
	}

	return app, nil
}

func (a *App) initDatabases() error {
	var err error

	// Postgres
	if a.pgDb, err = tpostgres.NewWithMigration(
		a.contextProvider(),
		a.config.Databases.Postgres.ConnectionString,
		a.config.Databases.Postgres.MigrationPath,
	); err != nil {
		return fmt.Errorf("database: %v", err)
	}

	// Repos
	a.passportRepo = passport.NewRepository(a.pgDb)
	a.departmentRepo = department.NewRepository(a.pgDb)
	a.employeeRepo = employee.NewRepository(a.pgDb, a.passportRepo, a.departmentRepo)

	return nil
}

func (a *App) initServices() error {

	a.employeeService = service.NewUserService(a.employeeRepo, a.passportRepo, a.departmentRepo)

	a.Server = api.NewServer(
		a.config.Port,
		a.logger,
		a.contextProvider,
		a.employeeService,
	)

	return nil
}

func (a *App) Run() {
	a.Server.Start()

	go func() {
		select {
		case err := <-a.Server.Notify():
			if !errors.Is(err, http.ErrServerClosed) {
				a.logger.Error(err.Error())
			}
		}
		a.notify <- struct{}{}
	}()
}

func (a *App) Notify() <-chan struct{} {
	return a.notify
}

func (a *App) Stop(ctx context.Context) {
	// Services
	if a.Server != nil {
		if err := a.Server.Shutdown(ctx); err != nil {
			a.logger.Errorf("Server: %v", err)
		}
	}

	// DB
	if a.pgDb != nil {
		if err := a.pgDb.Close(); err != nil {
			a.logger.Errorf("pg database: %v", err)
		}
	}
}
