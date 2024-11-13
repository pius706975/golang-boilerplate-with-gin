package auth

import (
	"go-gin/interfaces"
	"go-gin/middlewares"
	"go-gin/package/database/models"
	"go-gin/package/utils"

	"github.com/gin-gonic/gin"
)

type authService struct {
	repo interfaces.AuthRepo
}

func NewService(repo interfaces.AuthRepo) *authService {
	return &authService{repo}
}

type tokenResponse struct {
	Token string `json:"token"`
}

func (service *authService) SignUp(userData *models.User) (gin.H, int) {
	hashedPassword, err := utils.HashPassword(userData.Password)
	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	userData.Username = utils.GenerateUsername(userData.Email)
	userData.Password = hashedPassword

	newData, err := service.repo.SignUp(userData)
	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_users_email\" (SQLSTATE 23505)" {
			return gin.H{"status": 409, "message": "Email is already used"}, 409
		}
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"data": newData}, 201
}

func (service *authService) SignIn(userData *models.User) (gin.H, int) {
	user, err := service.repo.SignIn(userData.Email)

	if err != nil {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
	}

	if !utils.CheckPassword(user.Password, userData.Password) {
		return gin.H{"status": 401, "message": "Email or password is incorrect"}, 401
	}

	jwt := middlewares.NewToken(user.ID)

	token, err := jwt.CreateToken()

	if err != nil {
		return gin.H{"status": 500, "message": err.Error()}, 500
	}

	return gin.H{"data": user, "accessToken": tokenResponse{Token: token}}, 200
}
