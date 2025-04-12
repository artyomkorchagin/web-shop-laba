package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderSignIn(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}

func (h *Handler) renderSignUp(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func (h *Handler) renderMain(c *gin.Context) {
	emailValue := c.Request.Context().Value("email")

	email, ok := emailValue.(string)
	if !ok || email == "" {
		email = ""
	}

	if email != "" {
		c.Redirect(http.StatusSeeOther, "/api/v1/menu")
	}
	c.HTML(http.StatusOK, "index.html", nil)
}

func (h *Handler) renderAbout(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", nil)
}

func (h *Handler) renderProfile(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	role := c.Request.Context().Value("role").(string)

	user, err := h.services.user.GetUserByEmail(c, email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"LoggedIn": email != "",
		"Role":     role,
		"User":     user,
	})
}

func (h *Handler) renderMenu(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	role := c.Request.Context().Value("role").(string)
	c.HTML(http.StatusOK, "userpanel.html", gin.H{
		"LoggedIn": email != "",
		"Role":     role,
	})
}

func (h *Handler) renderUsers(c *gin.Context) {

	users, err := h.services.user.GetAllUsers(c)
	if err != nil {
		fmt.Println("error with users", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "listofusers.html", gin.H{
		"Users": users,
	})
}
