package postgres

import (
	"context"
	"time"

	"github.com/mizmorr/ingrytech/internal/config"
	logger "github.com/mizmorr/loggerm"
)

type PostgresRepo struct {
	db     *db
	stop   chan interface{}
	config *config.Config
	log    *logger.Logger
}

func NewPostgresRepo(ctx context.Context) (*PostgresRepo, error) {
	config := config.Get()
	log := logger.GetLoggerFromContext(ctx)

	ch := make(chan interface{})

	db, err := newDB(ctx, config)
	if err != nil {
		return nil, err
	}

	return &PostgresRepo{
		stop:   ch,
		config: config,
		log:    log,
		db:     db,
	}, nil
}

func (repo *PostgresRepo) Start(ctx context.Context) error {
	err := repo.db.dial(ctx, repo.config.ConnectAttempts, repo.config.Timeout)
	if err != nil {
		return err
	}
	go repo.keepAlive(ctx)

	return nil
}

func (repo *PostgresRepo) keepAlive(ctx context.Context) {
	repo.log.Debug().Msg("Keeping database connection alive...")

	for {
		select {
		case <-repo.stop:
			repo.log.Info().Msg("Keep alive worker is stopped..")
			return
		default:
			repo.maintainConnection(ctx)
		}
	}
}

func (repo *PostgresRepo) maintainConnection(ctx context.Context) {
	time.Sleep(repo.config.KeepAliveTimeout)

	connectionLost := false

	err := repo.db.con.Ping()
	if err != nil {
		connectionLost = true
		repo.log.Debug().Msg("[keepAlive] Lost connection, is trying to reconnect...")
	}

	if connectionLost {
		err = repo.db.dial(ctx, repo.config.ConnectAttempts, repo.config.Timeout)
		if err != nil {
			repo.log.Err(err).Msg("Failed to reconnect to PostgreSQL database")
		}
	}
}

func (repo *PostgresRepo) Stop(ctx context.Context) error {
	repo.log.Info().Msg("Stopping PostgreSQL repository..")

	sqlDB, err := repo.db.DB.DB()
	if err != nil {
		repo.log.Err(err).Msg("Failed to retrieve SQL connection")
		return err
	}
	select {
	case <-ctx.Done():
		repo.log.Warn().Msg("Context cancelled before stop signal could be sent")
	default:
		repo.stop <- struct{}{}
		repo.log.Info().Msg("Stop signal sent successfully")

		if err := sqlDB.Close(); err != nil {
			repo.log.Err(err).Msg("There was a problem with pgsql stopping")
		} else {
			repo.log.Info().Msg("PostgreSQL connection closed successfully")
		}
	}

	return nil
}
