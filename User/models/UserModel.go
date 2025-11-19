package models

import "database/sql"

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validator:"require , min =3"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min =6"`
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,

		password VARCHAR(255) NOT NULL,
		role VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(query)
	return err
}
