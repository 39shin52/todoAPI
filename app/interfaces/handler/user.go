package handler

import (
	"net/http"

	"github.com/39shin52/todoAPI/app/domain/entity"
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

func (uh UserHandler) GetUser(c *gin.Context) {
	var userName string
	if err := c.Bind(&userName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

		return
	}

	user, err := uh.userUsecase.SelectUser(userName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	userResponse := entity.User{
		ID:   user.ID,
		Mail: user.Mail,
		Work: user.Work,
	}

	c.JSON(http.StatusOK, gin.H{
		"response": userResponse,
	})
}
