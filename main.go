package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	XMLName	xml.Name	`xml:"person"`
	FirstName string	`xml:"firstName,attr"`
	LastName string	`xml:"lastName,attr"`
}

type Recipe struct {
	Name string `json:"name"`
	Tag []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

func main() {
	router := gin.Default()

	router.GET("/", IndexHandler)

	router.Run()

} 

func IndexHandler(c *gin.Context)  {
	c.XML(200, Person{
		FirstName: "Lixv",
		LastName:	"Yang",
	})
}