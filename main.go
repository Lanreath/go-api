package main

import (
	"fmt"

	_ "github.com/Lanreath/go-api/docs"

	"github.com/Lanreath/go-api/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Foodconnect API
//	@version		1
//	@description	This is a sample server Foodconnect server.

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

func main() {
	fmt.Println("Running...")
	fmt.Println("Initializing router...")
	router := gin.Default()
	fmt.Println("Initializing controller...")
	c := controller.NewController()

	v1 := router.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET("", c.GetUsers)
			users.GET("/:id", c.GetUser)
			users.POST("/users", c.PostUser)
			users.PUT("/:id", c.PutUser)
			users.DELETE("/:id", c.DeleteUser)
		}
		recipes := v1.Group("/recipes")
		{
			recipes.GET("", c.GetRecipes)
			recipes.GET("/:id", c.GetRecipe)
			recipes.POST("/recipes", c.PostRecipe)
			recipes.PUT("/:id", c.PutRecipe)
			recipes.DELETE("/:id", c.DeleteRecipe)
		}
		comments := v1.Group("/comments")
		{
			comments.GET("", c.GetComments)
			comments.GET("/:id", c.GetComment)
			comments.POST("/comments", c.PostComment)
			comments.PUT("/:id", c.PutComment)
			comments.DELETE("/:id", c.DeleteComment)
		}
		categories := v1.Group("/categories")
		{
			categories.GET("", c.GetCategories)
			categories.GET("/:id", c.GetCategory)
			categories.POST("/categories", c.PostCategory)
			categories.PUT("/:id", c.PutCategory)
			categories.DELETE("/:id", c.DeleteCategory)
		}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		router.Run(":8080")
	}
}
