package models

import "time"


type User struct {
	Id_user      string     `db:"id_user" form:"id_user" json:"id_user"`
	Email_user   string     `db:"email_user" form:"email_user" json:"email_user"`
	Password     string     `db:"password" form:"password" json:"password"`
	Phone_number string     `db:"phone_number" form:"phone_number" json:"phone_number"`
	Image_user   *string    `db:"image_user" from:"image_user" json:"image_user"`
	Created_at   *time.Time `db:"created_at" json:"created_at"`
	Updated_at   *time.Time `db:"updated_at" json:"updated_at"`
}