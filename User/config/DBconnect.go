package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("env loading failed")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Database connection error:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Cannot reach database:", err)
	}

	fmt.Println("✅ Connected to MySQL!")
}
