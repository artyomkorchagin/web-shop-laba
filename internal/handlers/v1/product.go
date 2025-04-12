package v1

import (
	"artyomkorchagin/web-shop/internal/types"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h Handler) addProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	price, err := strconv.Atoi(c.PostForm("price"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bad price format"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image upload failed"})
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create upload directory"})
		return
	}

	fileName := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveUploadedFile(file, fileName); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save uploaded file"})
		return
	}

	//imageURL := "/" + fileName

	product := &types.CreateProductRequest{
		Name:        name,
		Description: description,
		Price:       price,
	}

	if err := h.services.product.AddProduct(c, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to push product to db"})
		return
	}

	c.Redirect(http.StatusSeeOther, "/admin")
}
