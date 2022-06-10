package lead

import (
	"github.com/gofiber/fiber"
	"github.com/wakatara/11-go-projects/fiber-crm/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	var leads []Lead
	database.Database.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	var lead Lead
	database.Database.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	database.Database.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	var lead Lead
	database.Database.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No Lead found with ID delete!")
		return
	}
	database.Database.Delete(&lead)
	c.Send("Lead deleted.")
}
