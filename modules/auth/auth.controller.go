package auth

import (
	"go-gin/interfaces"
	"go-gin/package/database/models"
	"go-gin/package/utils"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type authController struct {
	service interfaces.AuthService
}

func NewController(service interfaces.AuthService) *authController {
	return &authController{service}
}

// SignUp godoc
// @Summary Register a new user
// @Description Register a new user with email, username, and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userData body models.SignUpRequest true "User data"
// @Success 201 
// @Failure 400 
// @Failure 409
// @Failure 500 
// @Router /api/auth/signup [post]
func (controller *authController) SignUp(ctx *gin.Context) {
	ctx.Header("Content-type", "application/json")

	var userData models.User

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	_, err = govalidator.ValidateStruct(&userData)
	if err != nil {
		ctx.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if !utils.ValidatePassword(userData.Password) {
		ctx.JSON(400, gin.H{"message": "Password length at least 8 characters, has at least 1 uppercase letter, 1 lowercase letter, 1 number, and 1 special character."})
		return
	}

	responseData, status := controller.service.SignUp(&userData)

	ctx.JSON(status, responseData)
}

// SignIn godoc
// @Summary Login as an authenticated user
// @Description Login with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param userData body models.SignInRequest true "User data"
// @Success 200 
// @Failure 401 
// @Failure 500 
// @Router /api/auth/signin [post]
func (controller *authController) SignIn(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	var userData models.User

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to parse request"})
		return
	}

	responseData, status := controller.service.SignIn(&userData)

	ctx.JSON(status, responseData)
}
