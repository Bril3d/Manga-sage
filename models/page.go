package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	Image     string
	ChapterID uint
}
