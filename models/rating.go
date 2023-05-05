package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	UserID    int
	User      User
	MangaID   int
	Manga     Manga
	ChapterID int
	Chapter   Chapter
	Score     int
}
