package db

import (
	"log"

	"gorm.io/gorm"
)

// GormUserRepository represents a repository for managing users in the database using GORM.
type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository creates a new instance of GormUserRepository and establishes a connection to the database.
func NewGormUserRepository() *GormUserRepository {
	db, err := Connect()
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return &GormUserRepository{db: db}
}

// CreateUser adds a new user to the database.
func (r *GormUserRepository) CreateUser(user *User) {
	r.db.Create(&User{UserName: user.UserName, Password: user.Password})
}

// GetUserInfo retrieves user information from the database based on the provided UserName.
func (r *GormUserRepository) GetUserInfo(UserName string) User {
	var res User
	r.db.Where("user_name = ?", UserName).First(&res)
	return res
}

// UpdateUserInfo updates the information of an existing user in the database.
func (r *GormUserRepository) UpdateUserInfo(UserName, newNameUser, newPasswordUser string) {
	newInfo := User{
		UserName: newNameUser,
		Password: newPasswordUser,
	}
	r.db.Model(&User{}).Where("user_name = ?", UserName).Updates(newInfo)
}
