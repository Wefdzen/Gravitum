package db

type UserRepository interface {
	CreateUser(user *User)
	GetUserInfo(UserName string) User
	UpdateUserInfo(UserName, newNameUser, newPasswordUser string)
}

func CreateUser(repo UserRepository, user *User) {
	repo.CreateUser(user)
}

func GetUserInfo(repo UserRepository, UserName string) User {
	return repo.GetUserInfo(UserName)
}

func UpdateUserInfo(repo UserRepository, UserName, newNameUser, newPasswordUser string) {
	repo.UpdateUserInfo(UserName, newNameUser, newPasswordUser)
}
