package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	TgUserId string `json:"tg_user_id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
}
