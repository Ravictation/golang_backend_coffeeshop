package models

import "time"

type Product struct {
	Id_product    string     `db:"id_product" form:"id_product" json:"id_product"`
	Product_name  string     `db:"product_name" form:"product_name" json:"product_name"`
	Price         float64    `db:"price" form:"price" json:"price"`
	Categories    string     `db:"categories" form:"categories" json:"categories"`
	Product_image string     `db:"product_image" json:"product_image"`
	Created_at    *time.Time `db:"created_at" json:"created_at"`
	Updated_at    *time.Time `db:"updated_at" json:"updated_at"`
}
