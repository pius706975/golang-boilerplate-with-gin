package interfaces

import (
	"go-gin/package/database/models"

	"github.com/gin-gonic/gin"
)

type AuthRepo interface {
	SignUp(userData *models.User) (*models.User, error)
	SignIn(email string) (*models.User, error)
}

type AuthService interface {
	SignUp(data *models.User) (gin.H, int)
	SignIn(data *models.User) (gin.H, int)
}