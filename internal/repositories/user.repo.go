package repositories

import (
	"errors"

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
	query := `INSERT INTO public.user ( 
				username,
				email_user, 
				password, 
				phone_number,
				image_user,
				role) 
				VALUES(
					:username,
					:email_user,
					:password, 
					:phone_number,
					:image_user,
					:role
				);`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data user created", nil
}

func (r *RepoUser) UpdateUser(data *models.User) (string, error) {
	query := `UPDATE public.user SET
				password = COALESCE(NULLIF(:password, ''), password),
				image_user = COALESCE(NULLIF(:image_user, ''), image_user),
				phone_number = COALESCE(NULLIF(:phone_number, ''), phone_number),
				updated_at = now()
			WHERE username = :username`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data has been updated", nil
}

func (r *RepoUser) GetUser(data *models.User) (*models.User, error) {
	query := `SELECT email_user, phone_number, image_user FROM public.user WHERE id_user=$1;`
	var userModel models.User
	err := r.Get(&userModel, query, data.Id_user)
	if err != nil {
		return nil, err
	}
	return &userModel, nil
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
	query := "SELECT * FROM public.user"
	err := r.Select(&users, query)

	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *RepoUser) DeleteUser(data *models.User) (string, error) {
	query := `DELETE FROM public.user WHERE id_user = :id_user;`

	_, err := r.NamedExec(query, data)
	if err != nil {
		return "", err
	}

	return "1 data has been Deleted", nil
}

func (r *RepoUser) GetAuthData(user string) (*models.User, error) {
	var result models.User
	q := `SELECT id_user, username, "role", "password" FROM public."user" WHERE username = ?`

	if err := r.Get(&result, r.Rebind(q), user); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("username not found")
		}

		return nil, err
	}

	return &result, nil
}
