package db

type User struct {
	ID       uint
	UserName string `json:"username"`
	Password string `json:"password"`
}
