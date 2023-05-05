package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User      User
	UserID    int
	MangaID   int
	Manga     Manga
	ChapterID int
	Chapter   Chapter
	Content   string
}
