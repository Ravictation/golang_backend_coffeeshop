package repositories

import (
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

	query := `UPDATE public.products SET product_name=:product_name, price=:price WHERE id_product = :id_product;`
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

func (r *RepoProduct) GetAllProduct(search string, page int, limit int, categories string) ([]models.Product, error) {
	var products []models.Product

	query := `SELECT product_name, price, categories, product_image FROM public.products WHERE 1=1`

	if search != "" {
		query += " AND product_name ILIKE '%' || $1 || '%'"
	}

	if categories != "" {
		query += " AND categories = $2"
	}

	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query += fmt.Sprintf(" OFFSET %d", offset)
	}

	err := r.Select(&products, query, search, categories)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
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
