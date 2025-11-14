package main

import (
	databases "github.com/Sachiink/Raw_Shop/config"
	"github.com/gin-gonic/gin"
)

func main() {
databases.Connect()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": "Server running on"})
	})
	r.Run(":8080")

}
