package service

import (
	"user_management_ms/config"
	"user_management_ms/model"
	"user_management_ms/repository"
)

func GetUsersService() []model.User {
	return []model.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}
}

func SaveUser(user *model.User) (model.User, error) {

	repo := repository.NewUserRepository(config.DB)
	err := repo.Save(user)
	if err != nil {
		return model.User{}, err
	}

	return *user, nil

}
