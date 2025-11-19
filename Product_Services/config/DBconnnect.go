package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect(c *gin.Context) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("env loadin failed")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	DB , err := sql.Open("mysql" , dsn)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	if err := DB.Ping(); err != nil{
		log.Fatal("connot reach database")
	}

	c.JSON(http.StatusOK , gin.H{"msg":"DB Connected succesfully"})	

}
