package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var migrate = make([]func(), 0)

func Add(mF func()) {
	migrate = append(migrate, mF)
}

func DB() *gorm.DB {
	return database
}

func Connect() {
	dsn := "host=localhost " +
		"user=postgres " +
		"password=postgresql " +
		"dbname=postgres " +
		"port=5432"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}

	database = db
}

func Migrate() {
	for _, f := range migrate {
		f()
	}
}
