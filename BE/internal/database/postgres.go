package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	conn *gorm.DB
}

func NewPostgresDB(dsn string) (IDatabase, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &postgresDB{conn: db}, nil
}

func (p *postgresDB) GetDB() *gorm.DB {
	return p.conn
}

func (p *postgresDB) Close() error {
	dbSQL, _ := p.conn.DB()
	return dbSQL.Close()
}
