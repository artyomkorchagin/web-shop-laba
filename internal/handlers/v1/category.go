package v1

import (
	"artyomkorchagin/web-shop/internal/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) addCategory(c *gin.Context) {
	name := c.PostForm("name")

	category := &types.CreateCategoryRequest{
		Name: name,
	}

	if err := h.services.category.AddCategory(c, category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push product to db"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
