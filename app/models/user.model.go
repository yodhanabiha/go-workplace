package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Age      int    `json:"Age"`
}