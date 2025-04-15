package v1

import (
	"artyomkorchagin/web-shop/internal/types"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type signInInput struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func (h *Handler) signUp(c *gin.Context) {
	var input types.CreateUserRequest

	if err := c.ShouldBind(&input); err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	err := h.services.user.AddUser(c, input)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.ShouldBind(&input); err != nil {
		c.Redirect(http.StatusSeeOther, "/")
		return
	}
	token, err := h.services.user.GenerateToken(c, input.Email, input.Password)
	if err != nil {
		fmt.Println(err)
		c.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) signOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, false)
	c.Redirect(http.StatusSeeOther, "/")
}
