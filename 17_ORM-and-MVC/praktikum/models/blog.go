package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	
	gorm.Model

	UserID uint

	Judul string `json:"judul" form:"judul"`

	Konten string `json:"konten" form:"konten"`
}