package gorm

import (
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/config"
	"github.com/PIRSON21/grpc-microservices/microservice-warehouses/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DbGorm struct is a wrapper around the gorm.DB instance to interact with the database.
type DbGorm struct {
	db *gorm.DB
}

// NewDbGorm initializes a new DbGorm instance with the provided database configuration.
func NewDbGorm(cfg *config.DBConfig) *DbGorm {
	return &DbGorm{
		db: initDB(cfg),
	}
}

func initDB(cfg *config.DBConfig) *gorm.DB {
	switch cfg.Driver {
	case config.PostgresDriver:
		db := initPostgresDB(cfg)
		if err := db.AutoMigrate(&models.Warehouse{}); err != nil {
			panic("failed to migrate PostgreSQL database: " + err.Error())
		}
		return db
	default:
		panic("unsupported database driver: " + cfg.Driver)
	}
}

func initPostgresDB(cfg *config.DBConfig) *gorm.DB {
	dsn := cfg.DSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("failed to connect to PostgreSQL database: " + err.Error())
	}
	sql, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB from gorm.DB: " + err.Error())
	}
	err = sql.Ping()
	if err != nil {
		panic("failed to ping PostgreSQL database: " + err.Error())
	}
	return db
}

func (d *DbGorm) Close() error {
	sqlDB, err := d.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
