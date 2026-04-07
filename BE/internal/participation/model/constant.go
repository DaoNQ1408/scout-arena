package model

import (
	"database/sql/driver"
	"errors"
)

type ParticipationStatus string

const (
	Draft    ParticipationStatus = "DRAFT"
	Absent   ParticipationStatus = "ABSENT"
	Review   ParticipationStatus = "REVIEW"
	Approved ParticipationStatus = "APPROVED"
	Denied   ParticipationStatus = "DENIED"
)

func (s ParticipationStatus) Value() (driver.Value, error) {
	return string(s), nil
}

func (s *ParticipationStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("kiểu dữ liệu từ DB không hợp lệ cho ParticipationStatus")
	}
	*s = ParticipationStatus(string(bytes))
	return nil
}
