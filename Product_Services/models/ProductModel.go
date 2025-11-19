package models

import "database/sql"

type ProductModel struct {
	Id         int     `json:"id"`
	Name       string  `json:"name" validate:"require"`
	Type       string  `json:"type" validate:"require , oneof=fish chicken"`
	SupplierID string  `json:"supplier_id" validate:"require"`
	Price      float64 `json:"price" validate:"require , gt=0"`
	Stock      int     `json:"stock" validate:"require , gte=0"`
	CreatedAt  string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}

func CreateTable(db *sql.DB) error {
	query := `
		 id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    type ENUM('fish', 'chicken') NOT NULL,
    supplier_id VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    stock INT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`

	_, err := db.Exec(query)
	return err
}
