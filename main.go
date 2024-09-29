package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ThanyawitNiti/go-fiber-crm/database"
	"github.com/ThanyawitNiti/go-fiber-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// create route
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	pwd, err_ := os.Getwd() //current directory of project
	if err_ != nil {
		panic(err_)
	}

	err_ = godotenv.Load(filepath.Join(pwd, "/config/.env"))
	if err_ != nil {
		fmt.Println("Error loading .env file")
		log.Fatalf("Error loading .env file: %v", err_)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	database.DBConn, err = gorm.Open("mysql", dsn)
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
