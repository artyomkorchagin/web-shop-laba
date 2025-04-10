package v1

import (
	"socialsecurity/internal/app/application"
	"socialsecurity/internal/app/user"
	"socialsecurity/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AllServices struct {
	user        *user.Service
	application *application.Service
}

type Handler struct {
	services AllServices
}

func NewAllSercivces(u *user.Service, a *application.Service) AllServices {
	return AllServices{
		user:        u,
		application: a,
	}
}

func NewHandler(services AllServices) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Static("/static", "../web/static/")

	router.LoadHTMLGlob("../web/static/html/*")

	main := router.Group("/")
	{
		main.GET("/", h.renderMain)
		main.GET("/about", h.renderAbout)
		main.GET("/sign-in-page", h.renderSignIn)
		main.GET("/sign-up-page", h.renderSignUp)
		main.GET("/application-form", h.renderApplicationForm)
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
		apiv1.GET("/application-form", h.renderApplicationForm)
		apiv1.POST("/application-apply", h.createApplication)
		apiv1.GET("/users", h.renderUsers)
		apiv1.GET("/menu", h.renderMenu)
		apiv1.POST("/sign-out", h.signOut)
		apiv1.GET("/applications", h.renderListOfApplications)
		apiv1.GET("/applications-review", h.renderAllApps)
		apiv1.GET("/applications/:id", h.getApplicationDetails)
		apiv1.PATCH("/applications/:id/action", h.workApplication)
	}

	return router
}
