package lead

import (
	"github.com/ThanyawitNiti/go-fiber-crm/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"` //make go understand JSON
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

//db := database.DBConn (connect to database)

func GetLeads(c *fiber.Ctx) { //c is working with data that's coming from user , e.g. find id in c.param , c will have all the data from the user
	db := database.DBConn
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead) //send result to user

}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with given ID")
		return
	}
	db.Delete(&lead) // Flag delete
	c.Send("Lead successfully deleted")
}
