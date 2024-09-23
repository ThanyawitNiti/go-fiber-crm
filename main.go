package main

import (
	"fmt"

	"github.com/ThanyawitNiti/go-fiber-crm/database"
	"github.com/ThanyawitNiti/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// create route
func setupRoutes(app *fiber.App) {
	app.Get(lead.GetLeads)
	app.Get(lead.GetLead)
	app.Post(lead.NewLead)
	app.Delete(lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("mysql", "leads.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)

	defer database.DBConn.Close() //Close connection when finish

}
