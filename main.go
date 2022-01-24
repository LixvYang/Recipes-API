package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"github.com/rs/xid"
	"io/ioutil"
	"encoding/json"
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
	 file, _ := ioutil.ReadFile("recipes.json")
	 _ = json.Unmarshal([]byte(file), &recipes)
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

func ListRecipesHandler(c *gin.Context)  {
	c.JSON(http.StatusOK, recipes)
}

func main() {
	router := gin.Default()

	router.POST("/recipes", NewRecipeHandler)
	router.GET("/recipes", ListRecipesHandler)
	router.Run()
} 