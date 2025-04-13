package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderCart(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)

	products, err := h.services.product.GetCartByUserEmail(c, email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.HTML(http.StatusOK, "cart.html", gin.H{
		"CartProducts": products,
	})
}

func (h *Handler) addToCart(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	if email == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/sign-in-page")
		return
	}

	productID, err := strconv.Atoi(c.PostForm("productID"))
	if err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	amount, err := strconv.Atoi(c.PostForm("amount"))
	if err != nil {
		fmt.Print(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	h.services.product.AddToCart(c, email, productID, amount)
}
