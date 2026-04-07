package model

type Season struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:50;not null"`
}

type Round struct {
}

type Challenge struct {
}
