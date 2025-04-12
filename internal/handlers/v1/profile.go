package v1

import (
	"artyomkorchagin/web-shop/internal/types"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateProfile(c *gin.Context) {
	var newUserData types.CreateUserRequest

	email := c.Request.Context().Value("email").(string)
	if err := c.ShouldBind(&newUserData); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
		return
	}

	newUserData.Email = email
	user := types.NewUser(newUserData)
	if err := h.services.user.UpdateUser(c, user); err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
