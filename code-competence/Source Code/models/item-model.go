package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Nama		string 	`json:"nama" form:"nama"`
	Deskripsi	string 	`json:"deskripsi" form:"deskripsi"`
	Jumlah		int64 	`json:"jumlah" form:"jumlah"`	
	Harga		int64 	`json:"harga" form:"harga"`
	ID_Kategori	int64 	`json:"id_kategori" form:"id_kategori"`
}
