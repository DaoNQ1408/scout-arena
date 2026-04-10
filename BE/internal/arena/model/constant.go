package model

import (
	"database/sql/driver"
	"fmt"
)

type ArenaStatus string

const (
	ArenaDraft    ArenaStatus = "DRAFT"
	ArenaIncoming ArenaStatus = "INCOMING"
	ArenaActive   ArenaStatus = "ACTIVE"
	ArenaFinished ArenaStatus = "FINISHED"
)

func (as *ArenaStatus) Scan(value interface{}) error {
	if value == nil {
		*as = ""
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	default:
		return fmt.Errorf("kiểu dữ liệu %T không hợp lệ cho ArenaStatus", value)
	}

	*as = ArenaStatus(str)
	return nil
}

func (as *ArenaStatus) Value() (driver.Value, error) {
	if as == nil {
		return nil, nil
	}
	return string(*as), nil
}

func (as *ArenaStatus) IsValid() bool {
	if as == nil {
		return false
	}
	switch *as {
	case ArenaDraft, ArenaIncoming, ArenaActive, ArenaFinished:
		return true
	}
	return false
}
