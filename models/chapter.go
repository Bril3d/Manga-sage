package models

import "gorm.io/gorm"

type Chapter struct {
	gorm.Model
	Title string
	Translator string
	Page []Page
}