package v1

import (
	"artyomkorchagin/web-shop/internal/middleware"
	"artyomkorchagin/web-shop/internal/services/cart"
	"artyomkorchagin/web-shop/internal/services/product"
	"artyomkorchagin/web-shop/internal/services/user"

	"github.com/gin-gonic/gin"
)

type AllServices struct {
	user    *user.Service
	product *product.Service
	cart    *cart.Service
}

type Handler struct {
	services *AllServices
}

func NewAllServices(u *user.Service, p *product.Service, c *cart.Service) *AllServices {
	return &AllServices{
		user:    u,
		product: p,
		cart:    c,
	}
}

func NewHandler(services *AllServices) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/static", "../web/static/")

	router.Static("/uploads", "./uploads")

	router.LoadHTMLGlob("../web/static/html/*")

	main := router.Group("/")
	{
		main.GET("/", h.renderMain)
		main.GET("/about", h.renderAbout)
		main.GET("/sign-in-page", h.renderSignIn)
		main.GET("/sign-up-page", h.renderSignUp)
	}

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	apiv1 := router.Group("/api/v1")
	apiv1.Use(middleware.AuthMiddleware())
	{
		apiv1.GET("/profile", h.renderProfile)
		apiv1.PATCH("/update-profile", h.updateProfile)
		apiv1.GET("/users", h.renderUsers)
		apiv1.GET("/menu", h.renderMenu)
		apiv1.POST("/sign-out", h.signOut)
		apiv1.POST("/add-product", h.addProduct)
	}

	return router
}
