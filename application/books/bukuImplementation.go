package books

import "gorm.io/gorm"

type BukuServices struct {
	database *gorm.DB
}

func NewBukuService(gorm *gorm.DB) *BukuServices {
	return &BukuServices{
		database: gorm,
	}
}
