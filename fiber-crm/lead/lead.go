package lead

import (
	"github.com/wakatara/11-go-projects/fiber-crm/database"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"github.com/gofiber/fiber"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")


}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}

