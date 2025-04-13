package v1

import (
	"net/http"

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
