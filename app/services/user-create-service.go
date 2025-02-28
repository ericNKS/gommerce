package services

import (
	"errors"

	"github.com/ericNKS/gommerce/app/models"
	"github.com/ericNKS/gommerce/app/repository"
)

type UserCreateService struct {
	repository repository.UserRepositoryInterface
}

func (ucs *UserCreateService) Execute(user *models.User) error {
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
