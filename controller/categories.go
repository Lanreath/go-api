package controller

import (
	"net/http"
	"strconv"

	"github.com/Lanreath/go-api/model"

	"github.com/gin-gonic/gin"
)

// GetCategory godoc
// @Summary Get a category
// @Description Get a category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} model.Category
// @Failure 400 {object} string
// @Router /categories/{id} [get]
func (c *Controller) GetCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	category, err := model.CategoryOne(uid)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
}

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Param q query string false "Search query"
// @Success 200 {array} model.Category
// @Failure 400 {object} string
// @Router /categories [get]
func (c *Controller) GetCategories(ctx *gin.Context) {
	q := ctx.Query("q")

	categorys, err := model.CategoriesAll(q)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, categorys)
}

// PostCategory godoc
// @Summary Create a category
// @Description Create a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body model.AddCategory true "Category"
// @Success 201 {object} string
// @Failure 400 {object} string
// @Router /categories [post]
func (c *Controller) PostCategory(ctx *gin.Context) {
	var newCategory model.AddCategory

	if err := ctx.BindJSON(&newCategory); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := newCategory.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	category := model.Category{
		Name: newCategory.Name,
	}
	lastID, err := category.InsertCategory()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	category.ID = lastID
	ctx.IndentedJSON(http.StatusCreated, category)
}

// PutCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body model.UpdateCategory true "Category"
// @Success 200 {object} model.Category
// @Failure 400 {object} string
// @Router /categories/{id} [put]
func (c *Controller) PutCategory(ctx *gin.Context) {
	var updateCategory model.UpdateCategory

	if err := ctx.BindJSON(&updateCategory); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := updateCategory.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	category := model.Category{
		ID:   updateCategory.ID,
		Name: updateCategory.Name,
	}
	if err := category.UpdateCategory(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /categories/{id} [delete]
func (c *Controller) DeleteCategory(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := model.DeleteCategory(uid); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
