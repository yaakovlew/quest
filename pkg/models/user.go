package models

import "gorm.io/gorm"

type RepoUser struct {
	gorm.Model
	TgUserId int    `db:"tg_user_id"`
	Name     string `db:"name"`
	Age      int    `db:"age"`
	Phone    string `db:"phone"`
}

type User struct {
	TgUserId int    `json:"tg_user_id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
}
