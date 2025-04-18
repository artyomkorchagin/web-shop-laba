package v1

import (
	"artyomkorchagin/web-shop/internal/types"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h Handler) addProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")

	price, err := strconv.ParseFloat(c.PostForm("price"), 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad price format"})
		return
	}
	quantity, err := strconv.Atoi(c.PostForm("quantity"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad quantity format"})
		return
	}

	category, err := strconv.Atoi(c.PostForm("category"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad category format"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	uploadDir := "../uploads/products"

	fileName := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	imageURL := strings.ReplaceAll(fileName[2:], "\\", "/")

	productreq := &types.CreateProductRequest{
		Name:          name,
		Description:   description,
		Price:         price,
		StockQuantity: quantity,
		Category:      category,
		ImageURL:      imageURL,
	}

	if err := h.services.product.AddProduct(c, productreq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push product to db"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/render-auth/add-stuff")
}

func (h *Handler) getProductsAnalytics(c *gin.Context) {
	role := c.Request.Context().Value("role").(string)
	products, err := h.services.product.GetProductsAnalytics(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products from db"})
		return
	}
	c.HTML(http.StatusOK, "products-analytics.html", gin.H{
		"Products": products,
		"Role":     role,
	})

}
