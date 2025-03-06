package services

import (
	"errors"

	"github.com/ericNKS/gommerce/app/models"
	"github.com/ericNKS/gommerce/app/repository"
)

type userCreateService struct {
	repository repository.UserRepositoryInterface
}

func UserCreateService(
	repo repository.UserRepositoryInterface,
) *userCreateService {

	return &userCreateService{
		repository: repo,
	}
}

func (ucs *userCreateService) Execute(user *models.User) error {
	if err := validate(user); err != nil {
		return err
	}
	return ucs.repository.Save(user)
}

func validate(user *models.User) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	return nil
}
