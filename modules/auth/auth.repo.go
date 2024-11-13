package auth

import (
	"errors"
	"go-gin/package/database/models"

	"gorm.io/gorm"
)

type authRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *authRepo {
	return &authRepo{db}
}

func (repo *authRepo) SignUp(data *models.User) (*models.User, error) {

	if err := repo.db.Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil

}


func (repo *authRepo) SignIn(email string) (*models.User, error) {
	var data models.User

	result := repo.db.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("get data failed")
	}

	return &data, nil
}