package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func GetConnection() *gorm.DB {
	connection := "host=localhost user=postgres password=postgres dbname=psm-db port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Error trying to connect with Postgres")
	}
	return db
}
