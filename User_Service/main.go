package main

import (
	databases "github.com/Sachiink/Raw_Shop/config"
	"github.com/Sachiink/Raw_Shop/models"
	"github.com/Sachiink/Raw_Shop/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	databases.Connect()
	godotenv.Load()
	models.CreateTable(databases.DB)
	r := gin.Default()
	routes.SetupRoutes(r)
	r.Run(":8000")

}
