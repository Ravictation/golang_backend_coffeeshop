package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoProduct struct {
	*sqlx.DB
}

func NewProduct(db *sqlx.DB) *RepoProduct {
	return &RepoProduct{db}
}

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func (r RepoProduct) CreateProduct(data *models.Product) (string, error) {
	q := `INSERT INTO public.products( 
		product_name, 
		price, 
		categories,
		product_image
	)
	VALUES( 
	:product_name,
	:price, 
	:categories,
	:product_image
	 );`
	_, err := r.NamedExec(q, data)
	if err != nil {
		return "", err
	}

	return "1 product data created", nil
}

func (r *RepoProduct) UpdateProduct(data *models.Product) (string, error) {

	query := `UPDATE public.products SET
	product_name = COALESCE(NULLIF(:product_name, ''), product_name),
	price = COALESCE(NULLIF(:price, 0), price),
	product_image = COALESCE(NULLIF(:product_image, ''), product_image),
	updated_at = now()
WHERE id_product = :id_product`
	_, er := r.NamedExec(query, data)
	if er != nil {
		fmt.Print(er)
		return "", er
	}

	return "1 data product has been updated", nil
}

// func (r *RepoProduct) GetProduct(data *models.Product) (interface{}, error) {

// 	var userModel models.User
// 	query := `SELECT * FROM public."user" WHERE id_user=$1;`
// 	fmt.Println(&userModel)
// 	err := r.Get(&userModel, query, data.Id_user)
// 	if err != nil {
// 		log.Fatal(err)
// 		return userModel, err
// 	}

// 	return userModel, nil
// }

func (r *RepoProduct) GetProduct(data *models.Product) (*models.Product, error) {
	var product models.Product

	query := `SELECT * FROM public.products WHERE id_product=$1;`

	err := r.Get(&product, query, data.Id_product)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &product, nil
}

func (r *RepoProduct) GetAllProduct(search string, page int, limit int, categories string) ([]models.Product, *Pagination, error) {
	var products []models.Product

	var (
		pgnt        = &Pagination{}
		recordcount int
	)

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = 5
	}

	offset := limit * (page - 1)

	sqltable := fmt.Sprintf("SELECT count(id_product) FROM products")

	r.QueryRow(sqltable).Scan(&recordcount)

	total := (recordcount / limit)

	remainder := (recordcount % limit)
	if remainder == 0 {
		pgnt.TotalPage = total
	} else {
		pgnt.TotalPage = total + 1
	}

	pgnt.CurrentPage = page
	pgnt.RecordPerPage = limit

	if page <= 0 {
		pgnt.Next = page + 1
	} else if page < pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = page + 1
	} else if page == pgnt.TotalPage {
		pgnt.Previous = page - 1
		pgnt.Next = 0
	}

	if search != "" && categories != "" {
		query := `SELECT * FROM products WHERE product_name ILIKE '%' || $1 || '%' AND categories = $2 LIMIT $3 OFFSET $4`
		r.Select(&products, query, search, categories, limit, offset)
	} else if search != "" {
		query := `SELECT * FROM products WHERE product_name ILIKE '%' || $1 || '%' LIMIT $2 OFFSET $3`
		r.Select(&products, query, search, limit, offset)
	} else if categories != "" {
		query := `SELECT * FROM products WHERE categories = $1 LIMIT $2 OFFSET $3`
		r.Select(&products, query, categories, limit, offset)
	} else {
		r.Select(&products, `SELECT * FROM products LIMIT $1 OFFSET $2`, limit, offset)
	}
	if len(products) == 0 {
		return nil, nil, errors.New("data not found.")
	}
	return products, pgnt, nil
}

// func (r *RepoProduct) GetAllProduct(data *models.Product) ([]models.Product, error) {
// 	var product []models.Product

// 	query := `SELECT product_name, price, categories, product_image FROM public.products`

// 	err := r.Select(&product, query)
// 	if err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return product, nil
// }

func (r *RepoProduct) DeleteProduct(data *models.Product) (string, error) {
	query := `DELETE FROM public.products WHERE id_product = :id_product;`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data product has been Deleted", nil
}
