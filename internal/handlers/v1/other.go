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
	role := c.Request.Context().Value("role").(string)

	products, err := h.services.product.GetAllProducts(c)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"LoggedIn": role != "",
		"Role":     role,
		"Products": products,
	})
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

func (h *Handler) renderAddStuff(c *gin.Context) {
	role := c.Request.Context().Value("role").(string)

	categories, err := h.services.category.GetAllCategories(c)
	if err != nil {
		fmt.Print("error with categories", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.HTML(http.StatusOK, "addstuff.html", gin.H{
		"Role":       role,
		"Categories": categories,
	})
}
