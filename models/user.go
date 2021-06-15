package models

type User struct {
	Username string `form:"user" json:"user" binding:"required,alphanum,min=5"`
	Password string `form:"password" json:"password" binding:"required,min=12"`
}
