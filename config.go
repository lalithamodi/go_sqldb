package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	//d, err := gorm.Open(postgres.Open("postgres://postgres:Lalitha@7654@localhost:5432/postgres?sslmode=disable", &gorm.config))
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
