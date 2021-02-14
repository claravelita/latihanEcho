package persistence

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"latihanEcho/domain/entities"
)

func NewDbConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=AdminPassword123 dbname=latihan port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entities.Buku{})
	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&entities.Anggota{})
	if err != nil {
		fmt.Println(err)
	}

	return db
}
