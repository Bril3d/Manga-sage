package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	User User
	Manga Manga
	Chapter Chapter
	Score int
}