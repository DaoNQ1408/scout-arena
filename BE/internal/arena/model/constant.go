package model

import (
	"database/sql/driver"
	"errors"
)

type ArenaStatus string

const (
	ArenaDraft    ArenaStatus = "DRAFT"
	ArenaIncoming ArenaStatus = "INCOMING"
	ArenaActive   ArenaStatus = "ACTIVE"
	ArenaFinished ArenaStatus = "FINISHED"
	ArenaDeleted  ArenaStatus = "DELETED"
)

func (s ArenaStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ArenaStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("kiểu dữ liệu từ DB không hợp lệ cho ArenaStatus")
	}
	*s = ArenaStatus(string(bytes))
	return nil
}

func (s ArenaStatus) IsValid() bool {
	switch s {
	case ArenaDraft, ArenaIncoming, ArenaActive, ArenaFinished, ArenaDeleted:
		return true
	}
	return false
}
