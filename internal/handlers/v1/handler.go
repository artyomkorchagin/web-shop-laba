package v1

import (
	"artyomkorchagin/web-shop/internal/middleware"
	"artyomkorchagin/web-shop/internal/services/category"
	"artyomkorchagin/web-shop/internal/services/order"
	"artyomkorchagin/web-shop/internal/services/product"
	"artyomkorchagin/web-shop/internal/services/user"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type AllServices struct {
	user     *user.Service
	product  *product.Service
	category *category.Service
	order    *order.Service
}

type Handler struct {
	services *AllServices
}

func NewAllServices(u *user.Service, p *product.Service, c *category.Service, o *order.Service) *AllServices {
	return &AllServices{
		user:     u,
		product:  p,
		category: c,
		order:    o,
	}
}

func NewHandler(services *AllServices) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	thisdir, _ := os.Getwd()

	router.Static("/static", path.Join(thisdir, "/web/static/"))

	router.Static("/uploads", path.Join(thisdir, "/uploads"))

	router.LoadHTMLGlob("/static/html/*")

	router.Use(middleware.PassUserData())

	main := router.Group("/")
	{
		main.GET("/", h.renderMain)
		main.GET("/about", h.renderAbout)
		main.GET("/sign-in-page", h.renderSignIn)
		main.GET("/sign-up-page", h.renderSignUp)
	}
	products := router.Group("/products")
	{
		products.GET("/:id", h.renderProduct)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	renderAuth := router.Group("/render-auth")
	renderAuth.Use(middleware.AuthMiddleware())
	{
		renderAuth.GET("/profile", h.renderProfile)
		renderAuth.GET("/users", h.renderUsers)
		renderAuth.GET("/add-stuff", h.renderAddStuff)
		renderAuth.GET("/cart", h.renderCart)
		renderAuth.GET("/order", h.renderOrder)
	}
	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.AuthMiddleware())
	{
		apiv1.PATCH("/update-profile", h.updateProfile)
		apiv1.GET("/sign-out", h.signOut)
		apiv1.POST("/add-product", h.addProduct)
		apiv1.POST("/add-category", h.addCategory)
		apiv1.POST("/add-order", h.addOrder)
		apiv1.POST("/add-to-cart", h.addToCart)
	}

	return router
}
