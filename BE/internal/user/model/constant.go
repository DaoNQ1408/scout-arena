package model

import (
	"database/sql/driver"
	"errors"
)

type UserStatus string
type TeamStatus string

const (
	UserActive UserStatus = "ACTIVE"
	UserMoved  UserStatus = "MOVED"
)

const (
	TeamActive   TeamStatus = "ACTIVE"
	TeamInactive TeamStatus = "INACTIVE"
)

func (s UserStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s TeamStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *UserStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("kiểu dữ liệu từ DB không hợp lệ cho UserStatus")
	}
	*s = UserStatus(string(bytes))
	return nil
}

func (s *TeamStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("kiểu dữ liệu từ DB không hợp lệ cho TeamStatus")
	}
	*s = TeamStatus(string(bytes))
	return nil
}
