package handler

import (
	// "fmt"
	"net/http"
	// "time"
	"github.com/gin-gonic/gin"
	"recipes/models"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type RecipesHandler struct {
	collection *mongo.Collection
	ctx context.Context
 }

 func NewRecipesHandler(ctx context.Context, collection *mongo. Collection) *RecipesHandler {
	 return &RecipesHandler{
	 collection: collection,
	 ctx: ctx,
	 }
}

// swagger:operation GET /recipes recipes listRecipes
// Returns list of recipes
// ---
// produces:
// - application/json
// responses:
//     '200':
//         description: Successful operation
func (handler *RecipesHandler) ListRecipesHandler(c *gin.
	Context) {
	 cur, err := handler.collection.Find(handler.ctx, bson.M{})
	 if err != nil {
		c.JSON(http.StatusInternalServerError,
		gin.H{"error": err.Error()})
		return
	 }
	 defer cur.Close(handler.ctx)
	 recipes := make([]models.Recipe, 0)
	for cur.Next(handler.ctx) {
		var recipe models.Recipe
		cur.Decode(&recipe)
		recipes = append(recipes, recipe)
	}
	c.JSON(http.StatusOK, recipes)
}


// // swagger:operation POST /recipes recipes newRecipe
// // Create a new recipe
// // ---
// // produces:
// // - application/json
// // responses:
// //     '200':
// //         description: Successful operation
// //     '400':
// //         description: Invalid input
// func NewRecipeHandler(c *gin.Context) {
// 	var recipe Recipe
// 	if err := c.ShouldBindJSON(&recipe); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	recipe.ID = primitive.NewObjectID()
// 	recipe.PublishedAt = time.Now()
	
// 	_, err = collection.InsertOne(ctx, recipe)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError,
// 		gin.H{"error": "Error while inserting a new recipe"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, recipe)
// }

// // swagger:operation PUT /recipes/{id} recipes updateRecipe
// // Update an existing recipe
// // ---
// // parameters:
// // - name: id
// //   in: path
// //   description: ID of the recipe
// //   required: true
// //   type: string
// // produces:
// // - application/json
// // responses:
// //     '200':
// //         description: Successful operation
// //     '400':
// //         description: Invalid input
// //     '404':
// //         description: Invalid recipe ID
// func UpdateRecipeHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	var recipe Recipe
// 	if err := c.ShouldBindJSON(&recipe); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	objectId, _ := primitive.ObjectIDFromHex(id)
// 	_, err = collection.UpdateOne(ctx, bson.M{
// 			"_id": objectId,
// 	}, bson.D{{"$set", bson.D{
// 			{"name", recipe.Name},
// 			{"instructions", recipe.Instructions},
// 			{"ingredients", recipe.Ingredients},
// 			{"tags", recipe.Tags},
// 	}}})
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError,
// 			gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been updated"})
// }

// // swagger:operation DELETE /recipes/{id} recipes deleteRecipe
// // Delete an existing recipe
// // ---
// // produces:
// // - application/json
// // parameters:
// //   - name: id
// //     in: path
// //     description: ID of the recipe
// //     required: true
// //     type: string
// // responses:
// //     '200':
// //         description: Successful operation
// //     '404':
// //         description: Invalid recipe ID
// func DeleteRecipeHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	index := -1
// 	for i := 0; i < len(recipes); i++ {
// 		if recipes[i].ID == id {
// 			index = i
// 			break
// 		}
// 	}

// 	if index == -1 {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
// 		return
// 	}

// 	recipes = append(recipes[:index], recipes[index+1:]...)

// 	c.JSON(http.StatusOK, gin.H{"message": "Recipe has been deleted"})
// }

// // swagger:operation GET /recipes/search recipes findRecipe
// // Search recipes based on tags
// // ---
// // produces:
// // - application/json
// // parameters:
// //   - name: tag
// //     in: query
// //     description: recipe tag
// //     required: true
// //     type: string
// // responses:
// //     '200':
// //         description: Successful operation
// func SearchRecipesHandler(c *gin.Context) {
// 	tag := c.Query("tag")
// 	listOfRecipes := make([]Recipe, 0)

// 	for i := 0; i < len(recipes); i++ {
// 		found := false
// 		for _, t := range recipes[i].Tags {
// 			if strings.EqualFold(t, tag) {
// 				found = true
// 			}
// 		}
// 		if found {
// 			listOfRecipes = append(listOfRecipes, recipes[i])
// 		}
// 	}

// 	c.JSON(http.StatusOK, listOfRecipes)
// }

// // swagger:operation GET /recipes/{id} recipes oneRecipe
// // Get one recipe
// // ---
// // produces:
// // - application/json
// // parameters:
// //   - name: id
// //     in: path
// //     description: ID of the recipe
// //     required: true
// //     type: string
// // responses:
// //     '200':
// //         description: Successful operation
// //     '404':
// //         description: Invalid recipe ID
// func GetRecipeHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	for i := 0; i < len(recipes); i++ {
// 		if recipes[i].ID == id {
// 			c.JSON(http.StatusOK, recipes[i])
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"error": "Recipe not found"})
// }