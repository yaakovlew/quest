package models

import "gorm.io/gorm"

type RepoQuest struct {
	gorm.Model
	Name          string `db:"name"`
	Description   string `db:"description"`
	AuthorComment string `db:"author_comment"`
	Point         string `db:"point"`
	AgeLevel      int    `db:"age"`
	Difficult     string `db:"difficult"`
	Duration      int    `db:"duration"`
	Location      string `db:"location"`
	Organizer     string `db:"organizer"`
}

type Quest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	AuthorComment string `json:"author_comment"`
	Point         string `json:"point"`
	AgeLevel      int    `json:"age_level"`
	Difficult     string `json:"difficult"`
	Duration      int    `json:"duration"`
	Location      string `json:"location"`
	Organizer     string `json:"organizer"`
}
