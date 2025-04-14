package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderOrder(c *gin.Context) {

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

	c.HTML(http.StatusOK, "order.html", gin.H{
		"CartProducts": products,
		"Sum_cart":     sum_cart,
	})
}

func (h *Handler) addOrder(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	address := c.PostForm("address")

	h.services.order.AddOrder(c, email, address)

}
