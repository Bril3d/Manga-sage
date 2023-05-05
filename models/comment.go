package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Manga   Manga
	Chapter Chapter
	Content string
}
