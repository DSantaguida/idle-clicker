package models

type User struct {
	Id       string
	Username string
	Password string //no plain text, find out best way to handle this
}

func CreateUser(username string, password string) *User {
	return &User{Username: username, Password: password}
}