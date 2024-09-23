package main

import (
	"github.com/gofiber/fiber"
)

// create route
func setupRoutes() {
	app.Get(GetLeads)
	app.Post(NewLead)
	app.Delete(deleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
}
