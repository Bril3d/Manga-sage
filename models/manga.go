package models

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	Title string
	Description string
}