package interfaces

import (
	"go-gin/package/database/models"

	"github.com/gin-gonic/gin"
)

type UserRepo interface {
	GetUsers() (*models.Users, error)
	GetUserById(id string) (*models.User, error)
}

type UserService interface {
	GetUsers() (gin.H, int)
	GetUserById(id string) (gin.H, int)
}
