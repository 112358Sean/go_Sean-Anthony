package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Nama	string 	`json:"nama" form:"nama"`
}
