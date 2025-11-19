package routes

import (
	databases "github.com/Sachiink/Raw_Shop/config"
	"github.com/Sachiink/Raw_Shop/controller"
	"github.com/Sachiink/Raw_Shop/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	db := databases.DB

	r.POST("/signup", func(c *gin.Context) {
		controller.Singup(c, db)
	})

	r.GET("/login", func(c *gin.Context) {
		controller.SignIn(c, db)
	})

	protected := r.Group("/user")

	protected.Use(middleware.AuthMidddleware())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "This is a protected route"})
		})
	}
}
