package repository

import (
	"github.com/ericNKS/gommerce/app/models"
	"github.com/ericNKS/gommerce/config/database"
)

type userRepository struct {
	db *database.Database
}

type UserRepositoryInterface interface {
	Save(user *models.User) error
}

func UserRepository() (*userRepository, error) {
	db, err := database.DB()
	if err != nil {
		return nil, err
	}

	return &userRepository{
		db: db,
	}, nil
}

func (ur *userRepository) Save(user *models.User) error {
	return nil
}
