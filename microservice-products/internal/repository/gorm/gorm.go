package gorm

import (
	"errors"
	"fmt"

	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/config"
	"github.com/PIRSON21/grpc-microservices/microservice-products/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBGorm struct {
	db *gorm.DB
}

var (
	NoSuchDriverError = errors.New("no such driver")
	ErrDBConnection   = errors.New("failed to connect to the database")
	ErrDBMigration    = errors.New("failed to migrate database")
)

// NewDBGorm initializes and returns a new DBGorm instance based on the provided configuration.
func NewDBGorm(cfg *config.DBConfig) (*DBGorm, error) {
	switch cfg.Driver {
	case "postgres":
		return setupPostgres(cfg)
	default:
		return nil, errors.Join(NoSuchDriverError, fmt.Errorf("driver: %s", cfg.Driver))
	}
}

// setupPostgres sets up a PostgreSQL database connection using GORM.
func setupPostgres(cfg *config.DBConfig) (*DBGorm, error) {
	dsn := cfg.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBConnection, err)
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDBMigration, err)
	}

	return &DBGorm{db: db}, nil
}
