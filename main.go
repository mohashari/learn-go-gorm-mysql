package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Code     string
	Name     string
	Email    string
	Password string
}

var db *gorm.DB

func main() {
	db = initDb()

	app := fiber.New()

	app.Get("/", MainPage)

	app.Listen(":3000")
}

func MainPage(c *fiber.Ctx) error {
	return c.SendString("Main Page")
}

func initDb() *gorm.DB {
	dns := "root:welcome1@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("Error")
	}
	fmt.Println("Db connected")
	db.AutoMigrate(&User{})
	return db
}
