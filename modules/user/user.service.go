package user

import (
	"go-gin/interfaces"

	"github.com/gin-gonic/gin"
)

type userService struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *userService {
	return &userService{repo}
}

func (service *userService) GetUsers() (gin.H, int) {
	users, err := service.repo.GetUsers()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"status": 200, "message": "All users fetched successfully", "data": users}, 200
}

func (service *userService) GetUserById(id string) (gin.H, int) {
	user, err := service.repo.GetUserById(id)

	if err != nil {
		return gin.H{"status": 500, "message": "Failed to retrieve user data"}, 500
	}

	return gin.H{"status": 200, "data": user}, 200
}