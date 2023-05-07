package models

import (
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	Number     string
	Translator string
	MangaID    uint
	Pages      []Page `gorm:"foreignKey:ChapterID"`
}
