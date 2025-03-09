package models

type User struct {
	Id       int    `json:"id" primaryKey:"true"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
