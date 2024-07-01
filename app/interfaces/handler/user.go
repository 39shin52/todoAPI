package handler

import (
	"net/http"

	"github.com/39shin52/todoAPI/app/interfaces/response"
	"github.com/39shin52/todoAPI/app/usecase"
	"github.com/gin-gonic/gin"
)

// gin使う予定

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (uh *UserHandler) GetUser(c *gin.Context) {
	userName := c.Query("userName")

	user, err := uh.userUsecase.SelectUser(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	userResponse := response.SelectUserResponse{
		ID:   user.ID,
		Mail: user.Mail,
		Work: user.Work,
	}

	c.JSON(http.StatusOK, gin.H{
		"response": userResponse,
	})
}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	users, err := uh.userUsecase.SelectUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": response.UsersResponse(users),
	})
}
