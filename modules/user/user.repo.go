package user

import (
	"errors"
	"go-gin/package/database/models"

	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (repo *userRepo) GetUsers() (*models.Users, error) {
	var data models.Users

	if err := repo.db.
		Select("id, name, username, email, created_at, updated_at").
		Order("created_at DESC").
		Find(&data).Error; err != nil {

		return nil, errors.New("failed to get data")
	}

	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil
}

func (repo *userRepo) GetUserById(id string) (*models.User, error) {
	var data models.User

	if err := repo.db.
		Select("id, name, username, email, created_at, updated_at").
		Find(&data, "id = ?", id).Error; err != nil {
		return nil, errors.New("failed to get data")
	}

	return &data, nil
}