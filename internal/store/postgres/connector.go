package postgres

import (
	"context"
	"database/sql"

	"github.com/mizmorr/ingrytech/internal/config"
	logger "github.com/mizmorr/loggerm"
	"github.com/pkg/errors"
)

type connector struct {
	*sql.DB
}

func newDBConnector(ctx context.Context, config *config.Config) (*connector, error) {
	log := logger.GetLoggerFromContext(ctx)

	log.Debug().Msg("Database config parsing...")

	sqlConnection, err := getConnection(config)
	if err != nil {
		log.Err(err).Msg("Error occured while parsing the config")

		return nil, errors.Wrap(err, "Failed to parse database config")
	}
	return &connector{sqlConnection}, nil
}

func getConnection(config *config.Config) (*sql.DB, error) {
	sqlDB, err := sql.Open("pgx", config.Postgres.URL)
	if err != nil {
		return nil, err
	}

	sqlDB.SetConnMaxIdleTime(config.MaxIdleTime)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	return sqlDB, nil
}
