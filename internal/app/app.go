package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/mizmorr/ingrytech/internal/config"
	"github.com/mizmorr/ingrytech/internal/delivery"
	"github.com/mizmorr/ingrytech/internal/service"
	"github.com/mizmorr/ingrytech/internal/store"
	"github.com/mizmorr/ingrytech/pkg/lifecycle"
	"github.com/mizmorr/ingrytech/pkg/server"
	logger "github.com/mizmorr/loggerm"
)

type component struct {
	Name        string
	ServiceTask lifecycle.Lifecycle
}
type App struct {
	log    *logger.Logger
	comps  []component
	config *config.Config
}

func New() *App {
	var (
		config = config.Get()
		log    = logger.Get(config.PathFile, config.Level)
	)
	return &App{
		log:    log,
		config: config,
	}
}

func (a *App) Start(ctx context.Context) error {
	if _, ok := ctx.Value("logger").(*logger.Logger); !ok {
		ctx = context.WithValue(ctx, "logger", a.log)
	}

	if err := a.setUp(ctx); err != nil {
		return err
	}

	okCh, errCh := make(chan interface{}), make(chan error)

	go func() {
		for _, comp := range a.comps {
			err := comp.ServiceTask.Start(ctx)
			if err != nil {
				errCh <- err
				return
			}
		}
		okCh <- struct{}{}
	}()
	select {
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info().Msg("Book service started")
		return nil
	}
}

func (a *App) Stop(ctx context.Context) error {
	if _, ok := ctx.Value("logger").(*logger.Logger); !ok {
		ctx = context.WithValue(ctx, "logger", a.log)
	}

	a.log.Info().Msg("Graceful shutdown is running..")

	okCh, errCh := make(chan interface{}), make(chan error)

	go func() {
		for _, comp := range a.comps {
			err := comp.ServiceTask.Stop(ctx)
			if err != nil {
				errCh <- err
			}
		}
		okCh <- struct{}{}
	}()
	select {
	case err := <-errCh:
		return err
	case <-okCh:
		a.log.Info().Msg("Book service is stopped")
		return nil
	}
}

func (a *App) setUp(ctx context.Context) error {
	repo := store.NewInMemoryRepo()

	service := service.NewBookService(repo)

	bookController := delivery.NewBookController(service)

	handler := echo.New()

	delivery.NewRouter(handler, bookController)

	httpServer := server.New(handler, a.config.HttpHost, a.config.HttpPort, a.config.ShutdownTimeout)

	a.comps = append(a.comps, component{Name: "server", ServiceTask: httpServer})

	return nil
}
