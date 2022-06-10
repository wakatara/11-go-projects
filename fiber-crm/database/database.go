package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "leads.db"

func Connect() error {
	var err error

	Database, err = gorm.Open(sqlite.Open(DATABASE_URI), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		panic(err)
	}

	return nil
}
