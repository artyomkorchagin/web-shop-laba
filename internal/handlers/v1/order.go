package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) renderOrder(c *gin.Context) {

	email := c.Request.Context().Value("email").(string)

	products, _ := h.services.product.GetCartByUserEmail(c, email)

	c.HTML(http.StatusOK, "order.html", gin.H{
		"Products": products,
	})
}

func (h *Handler) addOrder(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)
	address := c.PostForm("address")

	h.services.order.AddOrder(c, email, address)

}
