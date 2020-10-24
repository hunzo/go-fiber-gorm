package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/hunzo/go-fiber-gorm/database"
	"github.com/hunzo/go-fiber-gorm/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type PostBody struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

func initdatabase() { //in main function
	var err error
	database.DBcon, err = gorm.Open(sqlite.Open("./database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to database")

	database.DBcon.AutoMigrate(&models.Token{})
	fmt.Println("Database Migrated!!")
}

func main() {
	r := fiber.New()
	initdatabase()

	r.Get("/", home)
	r.Get("/user/:v1/:v2", param)
	r.Get("/query", query)
	r.Post("/post", postbody)
	r.Get("/redirect", redirect)
	r.Get("/read", read)
	r.Get("/create/:a/:r", create)
	r.Get("/update/:id/:a/:r", update)
	r.Get("/delete/:id", delete)

	r.Listen(8080)
}

func delete(c *fiber.Ctx) {
	db := database.DBcon
	db.Delete(&models.Token{}, c.Params("id"))

	c.JSON(fiber.Map{
		"info": "delete id " + c.Params("id"),
	})

}

func update(c *fiber.Ctx) {
	db := database.DBcon
	db.Model(&models.Token{}).Where(
		"ID = ?", c.Params("id")).Updates(map[string]interface{}{
		"AccessToken":  c.Params("a"),
		"RefreshToken": c.Params("r"),
	})
	c.JSON(fiber.Map{
		"ID": c.Params("id"),
		"a":  c.Params("a"),
		"r":  c.Params("r"),
	})

}

func create(c *fiber.Ctx) {
	a := c.Params("a")
	r := c.Params("r")

	db := database.DBcon
	db.Create(&models.Token{AccessToken: a, RefreshToken: r})

	c.JSON(fiber.Map{
		"update": "data",
	})
}

func read(c *fiber.Ctx) {
	db := database.DBcon
	var t []models.Token
	db.Find(&t)
	c.JSON(t)
}

func redirect(c *fiber.Ctx) {
	c.Redirect("http://www.nida.ac.th", 302)

}

func postbody(c *fiber.Ctx) {
	pbody := new(PostBody)
	if e := c.BodyParser(pbody); e != nil {
		c.Status(503).JSON(fiber.Map{
			"error": "don't have  body",
		})
		return
	}
	c.JSON(pbody)

}

func query(c *fiber.Ctx) {
	user := c.Query("user")
	c.JSON(fiber.Map{
		"query": user,
	})

}

func home(c *fiber.Ctx) {
	c.JSON(fiber.Map{
		"message": "home",
	})
}

func param(c *fiber.Ctx) {
	v1 := c.Params("v1")
	v2 := c.Params("v2")
	c.JSON(fiber.Map{
		"v1": v1,
		"v2": v2,
	})

}
