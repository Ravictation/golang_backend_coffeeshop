package repositories

import (
	"fmt"
	"log"

	"github.com/Ravictation/golang_backend_coffeeshop/internal/models"

	"github.com/jmoiron/sqlx"
)

type RepoUser struct {
	*sqlx.DB
}

func NewUser(db *sqlx.DB) *RepoUser {
	return &RepoUser{db}
}

func (r *RepoUser) CreateUser(data *models.User) (string, error) {
	query := `INSERT INTO coffeshop."user" ( 
				email, 
				pass, 
				phone_number) 
				VALUES(
					:email,
					:pass, 
					:phone_number
				);`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data user created", nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {

	query := `UPDATE coffeshop.user SET pass=:pass, phone_number=:phone_number WHERE id_user = :id_user;`
	_, er := r.NamedExec(query, data)
	if er != nil {
		fmt.Print("ini errornya", er)
		return "", er
	}

	return "1 data has been updated", nil
}

func (r *RepoUser) GetUser(data *models.User) (interface{}, error) {

	fmt.Println(data)
	var userModel models.User
	query := `SELECT * FROM coffeshop."user" WHERE id_user=$1;`
	fmt.Println(&userModel)
	err := r.Get(&userModel, query, data.Id_user)
	if err != nil {
		log.Fatal(err)
		return userModel, err
	}

	return userModel, nil
}

func (r *RepoUser) GetAllUser(data *models.User) ([]models.User, error) {

	//var user models.User
	// query := `SELECT email, pass, phone_number, created_at, updated_at FROM coffeshop."user";`
	// err := r.Get(&user, query)
	//fmt.Printf("%v\n", user)
	//fmt.Println(err)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return user, err
	// }

	// return user, nil
	var users []models.User
	query := "SELECT * FROM coffeshop.user"
	err := r.Select(&users, query)

	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM coffeshop."user" WHERE id_user = :id_user;`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data has been Deleted", nil
}
