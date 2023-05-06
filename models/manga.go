package models

import "gorm.io/gorm"

type Manga struct {
	gorm.Model
	Title       string `gorm:"unique;not null"`
	Author      string
	Genre       string
	Description string
	Cover_Image string
	Chapters    []Chapter `gorm:"foreignKey:MangaID"`
}
