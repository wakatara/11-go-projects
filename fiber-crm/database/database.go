package database

import (
	"fmt"
	"github.com/wakatara/11-go-projects/fiber-crm/lead"
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

	Database.AutoMigrate(&lead.Lead{})
	fmt.Println("DB Migrated.")
	return nil
}
