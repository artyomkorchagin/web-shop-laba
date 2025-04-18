package v1

import (
	"fmt"
	"net/http"
	"strconv"

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
		"Domain":   "localhost:3000",
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

func (h *Handler) renderProduct(c *gin.Context) {
	role := c.Request.Context().Value("role").(string)
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("something wrong with product id ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	product, err := h.services.product.GetProductById(c, productID)
	if err != nil {
		fmt.Println("something wrong with product", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.services.product.IncreaseViewCount(c, productID)
	if err != nil {
		fmt.Println("something wrong with view count", err)
	}

	c.HTML(http.StatusOK, "product-detail.html", gin.H{
		"LoggedIn": role != "",
		"Role":     role,
		"Product":  product,
	})
}
