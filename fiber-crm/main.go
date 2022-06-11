package main

import (
	"fmt"
	"github.com/gofiber/fiber"
	"github.com/wakatara/11-go-projects/fiber-crm/database"
	"github.com/wakatara/11-go-projects/fiber-crm/lead"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()
	database.Connect()
	database.Database.AutoMigrate(&lead.Lead{})
	fmt.Println("DB Migrated")
	setupRoutes(app)
	app.Listen(3000)
}
