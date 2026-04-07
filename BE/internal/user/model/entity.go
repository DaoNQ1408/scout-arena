package model

type Rank struct {
	ID   uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" gorm:"size:50;not null"`
}

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email" gorm:"uniqueIndex"`
	Username string `json:"username" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`

	RankID uint `json:"rank_id"`
	Rank   Rank `json:"rank" gorm:"foreignKey:RankID"`
}
