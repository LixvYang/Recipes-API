package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"github.com/rs/xid"
)

type Recipe struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Tag []string `json:"tags"`
	Ingredients []string `json:"ingredients"`
	Instructions []string `json:"instructions"`
	PublishedAt time.Time `json:"publishedAt"`
}

var recipes []Recipe
func init() {
 	recipes = make([]Recipe, 0)
}

func NewRecipeHandler(c *gin.Context)  {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			return
		return 
	}
	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(http.StatusOK, recipe)
}

func main() {
	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)

	router.Run()
} 