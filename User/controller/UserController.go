package controller

import (
	"database/sql"
	"net/http"

	"github.com/Sachiink/Raw_Shop/models"
	"github.com/Sachiink/Raw_Shop/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Singup(c *gin.Context , db *sql.DB) {

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		return
	}

	hashpassword,_ := bcrypt.GenerateFromPassword([] byte(input.Password) , bcrypt.DefaultCost)

	query := "INSEERT into users(username , email, password) values(?,?,?)"

	res , err := db.Exec(query , input.Username , input.Email , string(hashpassword))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email might already exist"})
		return
	}

	id , _ := res.LastInsertId()

	c.JSON(http.StatusOK , gin.H{"message": "User created", "user_id": id})

}

func SignIn(c *gin.Context , db *sql.DB){
	var input models.User
if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest , gin.H{"error":err.Error()})
		return
	}
var user models.User


	err := db.QueryRow("SELECT id, password FROM users WHERE email = ?", input.Email).Scan(&user.Id, &user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token , _ := utils.GenerateToken(user.Id)

	c.JSON(http.StatusOK, gin.H{"token": token})
}