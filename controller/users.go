package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lanreath/go-api/model"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Get a user
// @Description Get a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} string
// @Router /users/{id} [get]
func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	user, err := model.UserOne(uid)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Param q query string false "Search query"
// @Success 200 {array} model.User
// @Failure 400 {object} string
// @Router /users [get]
func (c *Controller) GetUsers(ctx *gin.Context) {
	q := ctx.Query("q")

	users, err := model.UsersAll(q)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// PostUser godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.AddUser true "User"
// @Success 201 {object} model.User
// @Failure 400 {object} string
// @Router /users [post]
func (c *Controller) PostUser(ctx *gin.Context) {
	var newUser model.AddUser

	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := newUser.Validation(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	user := model.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Password: newUser.Password,
	}
	lastID, err := user.InsertUser()
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	user.ID = lastID
	ctx.IndentedJSON(http.StatusCreated, user)
}

// PutUser godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body model.UpdateUser true "User"
// @Success 200 {object} model.User
// @Failure 400 {object} string
// @Router /users/{id} [put]
func (c *Controller) PutUser(ctx *gin.Context) {
	var updateUser model.UpdateUser

	if err := ctx.BindJSON(&updateUser); err != nil {
		fmt.Println("error")
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := updateUser.Validation(); err != nil {
		fmt.Println("here")
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	user := model.User{
		ID:       updateUser.ID,
		Name:     updateUser.Name,
		Email:    updateUser.Email,
		Password: updateUser.Password,
	}
	if err := user.UpdateUser(); err != nil {
		fmt.Println("there")
		fmt.Println(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, user)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /users/{id} [delete]
func (c *Controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	uid, err := strconv.Atoi(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	if err := model.DeleteUser(uid); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
}
