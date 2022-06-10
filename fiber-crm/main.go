package main

import (
	"github.com/wakatara/11-go-projects/fiber-crm/lead"
	"github.com/gofiber/fiber"
	)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

// func initDatabase() {
// 	var err error
// 	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
// 	if err != nil {
// 		panic("failed to connect to database")
// 	}
// 	fmt.Println("Connection opened to DB")
// 	database.DBConn.AutoMigrate(&lead.Lead{})
// 	fmt.Println("DB Migrated")
// }

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}
