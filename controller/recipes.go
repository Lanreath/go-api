package controller

import (
	"net/http"
	"strconv"

	"github.com/Lanreath/go-api/model"

	"github.com/gin-gonic/gin"
)

// GetRecipe godoc
// @Summary Get a recipe
// @Description Get a recipe by ID
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param id path int true "Recipe ID"
// @Success 200 {object} model.Recipe
// @Failure 400 {object} string
// @Router /recipes/{id} [get]
func (c *Controller) GetRecipe(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	recipe, err := model.RecipeOne(uid)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, recipe)
}

// GetRecipes godoc
// @Summary Get all recipes
// @Description Get all recipes
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param q query string false "Search query"
// @Success 200 {array} model.Recipe
// @Failure 400 {object} string
// @Router /recipes [get]
func (c *Controller) GetRecipes(ctx *gin.Context) {
	q := ctx.Query("q")

	recipes, err := model.RecipesAll(q)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, recipes)
}

// PostRecipe godoc
// @Summary Create a recipe
// @Description Create a recipe
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param recipe body model.AddRecipe true "Recipe"
// @Success 201 {object} model.Recipe
// @Failure 400 {object} string
// @Router /recipes [post]
func (c *Controller) PostRecipe(ctx *gin.Context) {
	var newRecipe model.AddRecipe

	if err := ctx.BindJSON(&newRecipe); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := newRecipe.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	recipe := model.Recipe{
		Name:        newRecipe.Name,
		Ingredients: newRecipe.Ingredients,
		Steps:       newRecipe.Steps,
		CategoryID:  newRecipe.CategoryID,
	}
	lastID, err := recipe.InsertRecipe()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	recipe.ID = lastID
	ctx.IndentedJSON(http.StatusCreated, recipe)
}

// PutRecipe godoc
// @Summary Update a recipe
// @Description Update a recipe
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param recipe body model.UpdateRecipe true "Recipe"
// @Success 200 {object} model.Recipe
// @Failure 400 {object} string
// @Router /recipes/{id} [put]
func (c *Controller) PutRecipe(ctx *gin.Context) {
	var updateRecipe model.UpdateRecipe

	if err := ctx.BindJSON(&updateRecipe); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := updateRecipe.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	recipe := model.Recipe{
		ID:          updateRecipe.ID,
		Name:        updateRecipe.Name,
		Ingredients: updateRecipe.Ingredients,
		Steps:       updateRecipe.Steps,
		CategoryID:  updateRecipe.CategoryID,
	}
	if err := recipe.UpdateRecipe(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, recipe)
}

// DeleteRecipe godoc
// @Summary Delete a recipe
// @Description Delete a recipe
// @Tags recipes
// @Accept  json
// @Produce  json
// @Param id path int true "Recipe ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /recipes/{id} [delete]
func (c *Controller) DeleteRecipe(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := model.DeleteRecipe(uid); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Recipe deleted"})
}
