package controller

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/Sachiink/Raw_Shop/models"
	"github.com/Sachiink/Raw_Shop/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Singup(c *gin.Context, db *sql.DB) {

	var input models.User

	// Bind JSON
	if err := c.ShouldBindJSON(&input); err != nil {

		var verrs validator.ValidationErrors
		if errors.As(err, &verrs) {
			for _, v := range verrs {
				switch v.Field() {
				case "Email":
					c.JSON(http.StatusBadRequest, gin.H{"error": "Email must be in a valid format like test@gmail.com"})
					return

				case "Password":
					c.JSON(http.StatusBadRequest, gin.H{"error": "Pasword must be at least 6 character"})
					return 
				}
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check missing fields
	missing := []string{}

	if input.Username == "" {
		missing = append(missing, "username")
	}
	if input.Email == "" {
		missing = append(missing, "email")
	}
	if input.Role == "" {
		missing = append(missing, "role")
	}
	if input.Password == "" {
		missing = append(missing, "password")
	}

	// If ANY missing, show which ones
	if len(missing) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":         "Missing required fields",
			"missingFields": missing,
		})
		return
	}

	// Hash password
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Insert Query
	query := "INSERT INTO users (username, email, role, password) VALUES (?,?,?,?)"
	res, err := db.Exec(query, input.Username, input.Email, input.Role, string(hashpassword))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email might already exist"})
		return
	}

	// Last Insert ID
	id, _ := res.LastInsertId()

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user_id": id,
	})
}

func SignIn(c *gin.Context, db *sql.DB) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User

	err := db.QueryRow("SELECT id, username , email, role ,password FROM users WHERE email = ?", input.Email).
		Scan(&user.Id, &user.Username, &user.Email, &user.Role , &user.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := utils.GenerateToken(user.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"user": user,
		"token": token})
}
