package database

import (
	"gorm.io/gorm"
)

type IDatabase interface {
	GetDB() *gorm.DB
	Close() error
}

type Config struct {
	Driver string
	DSN    string
}

func NewDatabase(cfg Config) (IDatabase, error) {
	switch cfg.Driver {
	case "postgres":
		return NewPostgresDB(cfg.DSN)
	// case "mysql":
	//    return NewMySQLDB(cfg.DSN)
	default:
		return nil, nil
	}
}
