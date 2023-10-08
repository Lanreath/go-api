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
func (c *Controller) GetComment(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	comment, err := model.CommentOne(uid)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, comment)
}

// GetComments godoc
// @Summary Get all comments
// @Description Get all comments
// @Tags comments
// @Accept  json
// @Produce  json
// @Param q query string false "Search query"
// @Success 200 {array} model.Comment
// @Failure 400 {object} string
// @Router /comments [get]
func (c *Controller) GetComments(ctx *gin.Context) {
	q := ctx.Query("q")

	comments, err := model.CommentsAll(q)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, comments)
}

// PostComment godoc
// @Summary Create a comment
// @Description Create a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body model.AddComment true "Comment"
// @Success 201 {object} model.Comment
// @Failure 400 {object} string
// @Router /comments [post]
func (c *Controller) PostComment(ctx *gin.Context) {
	var newComment model.AddComment

	if err := ctx.BindJSON(&newComment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := newComment.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	comment := model.Comment{
		Subject: newComment.Subject,
		Content: newComment.Content,
	}
	lastID, err := comment.InsertComment()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	comment.ID = lastID
	ctx.IndentedJSON(http.StatusCreated, comment)
}

// PutComment godoc
// @Summary Update a comment
// @Description Update a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param comment body model.UpdateComment true "Comment"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /comments [put]
func (c *Controller) PutComment(ctx *gin.Context) {
	var updateComment model.UpdateComment

	if err := ctx.BindJSON(&updateComment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := updateComment.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	comment := model.Comment{
		ID:      updateComment.ID,
		Subject: updateComment.Subject,
		Content: updateComment.Content,
	}
	if err := comment.UpdateComment(); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, comment)
}

// DeleteComment godoc
// @Summary Delete a comment
// @Description Delete a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param id path int true "Comment ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /comments/{id} [delete]
func (c *Controller) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	if err := model.DeleteComment(uid); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
