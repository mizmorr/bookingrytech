package postgres

import (
	"context"
	"database/sql"
	"sync"
	"time"

	"github.com/mizmorr/ingrytech/internal/config"
	logger "github.com/mizmorr/loggerm"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type db struct {
	con *connector
	*gorm.DB
}

var (
	pgInstance *db
	once       sync.Once
)

func newDB(ctx context.Context, config *config.Config) (*db, error) {
	connector, err := newDBConnector(ctx, config)
	if err != nil {
		return nil, err
	}
	pgInstance = &db{
		con: connector,
	}

	return pgInstance, nil
}

func (db *db) dial(ctx context.Context, attempts int, timeout time.Duration) error {
	err := db.establishConnect(ctx, db.con.DB, attempts, timeout)
	return err
}

func (db *db) establishConnect(ctx context.Context, connection *sql.DB, attempts int, timeout time.Duration) error {
	log := logger.GetLoggerFromContext(ctx)

	var (
		err    error
		dbGorm *gorm.DB
	)
	for attempts > 0 {
		dbGorm, err = gorm.Open(
			postgres.New(postgres.Config{
				Conn: pgInstance.con.DB,
			}), &gorm.Config{})
		if err == nil {

			log.Info().Msg("Connect to postgres is established")
			db.DB = dbGorm

			break
		}
		log.Error().Err(err).Msg("Failed to connect to pg, retrying...")

		time.Sleep(timeout)

		attempts--
	}
	if err != nil {
		return errors.Wrap(err, "Cannot connect to PostgreSQL database")
	}
	return nil
}
