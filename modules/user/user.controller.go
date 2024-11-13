package user

import (
	"go-gin/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service interfaces.UserService
}

func NewController(service interfaces.UserService) *userController {
	return &userController{service}
}

// GetUsers godoc
// @Summary Get all users
// @Description Fetch all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 
// @Failure 500 
// @Router /api/user/ [get]
func (controller *userController) GetUsers(ctx *gin.Context) {
	responseData, status := controller.service.GetUsers()
	ctx.JSON(status, responseData)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Fetch the user details based on the ID provided
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param Authorization header string true "Authorization token"
// @Success 200 
// @Failure 401 
// @Failure 500 
// @Router /api/user/{id} [get]
func (controller *userController) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	responseData, status := controller.service.GetUserById(id)
	
	ctx.JSON(status, responseData)
}

// GetProfile godoc
// @Summary Get user profile
// @Description Fetch the user profile based on the decoded ID from access token
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 
// @Failure 401 
// @Failure 500 
// @Router /api/user/profile [get]
func (controller *userController) GetProfile(ctx *gin.Context) {
	userID, exists := ctx.Get("id")
	
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	responseData, status := controller.service.GetUserById(userID.(string))
	
	ctx.JSON(status, responseData)
}