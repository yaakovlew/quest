package models

import "gorm.io/gorm"

type RepoQuest struct {
	gorm.Model
	Name          string `db:"name"`
	Description   string `db:"description"`
	AuthorComment string `db:"author_comment"`
	Point         string `db:"point"`
	AgeLevel      int    `db:"age"`
	difficult     string `db:"difficult"`
	duration      int    `db:"duration"`
	location      string `db:"location"`
	organizer     string `db:"organizer"`
}


type Quest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AuthorComment string `json:"author_comment"`
	Point         string `json:"point"`
	AgeLevel      int    `json:"age_level"`
	difficult     string `json:"difficult"`
	duration      int    `json:"duration"`
	location      string `json:"location"`
	organizer     string `json:"organizer"`
}
