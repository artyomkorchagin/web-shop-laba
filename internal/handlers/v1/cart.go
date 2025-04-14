package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderCart(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)

	sum_cart := 0
	products, err := h.services.product.GetCartByUserEmail(c, email)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	for i := range products {
		products[i].Sum = int(products[i].Price) * products[i].StockQuantity
		sum_cart += products[i].Sum
	}

	c.HTML(http.StatusOK, "cart.html", gin.H{
		"CartProducts": products,
		"Sum_cart":     sum_cart,
	})
}

func (h *Handler) addToCart(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	if email == "" {
		c.Redirect(http.StatusTemporaryRedirect, "/sign-in-page")
		return
	}

	productID, err := strconv.Atoi(c.PostForm("product_id"))
	if err != nil {
		fmt.Print("prod_id   ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	amount, err := strconv.Atoi(c.PostForm("amount"))
	fmt.Print(amount)
	if err != nil {
		fmt.Print("amount   ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	h.services.product.AddToCart(c, email, productID, amount)
}
