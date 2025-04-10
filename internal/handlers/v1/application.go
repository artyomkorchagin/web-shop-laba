package v1

import (
	"errors"
	"fmt"
	"net/http"

	utils "socialsecurity/internal/mssql"
	"socialsecurity/internal/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) renderApplicationForm(c *gin.Context) {
	benefits, err := h.services.application.GetBenefits(c)
	if err != nil {
		fmt.Println("error with benefits", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	services, err := h.services.application.GetServices(c)
	if err != nil {
		fmt.Println("error with services", err)
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.HTML(http.StatusOK, "addapplication.html", gin.H{
		"Services": services,
		"Benefits": benefits,
	})
}

func (h *Handler) createApplication(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)

	user, err := h.services.user.GetUserByEmail(c, email)
	if err != nil {
		fmt.Println(err)
		return
	}
	var input types.CreateApplicationRequest

	input.UserID = user.UserID
	benefit := c.PostForm("benefit_id")
	if benefit != "" {
		input.BenefitID, err = uuid.Parse(benefit)
		if err != nil {
			fmt.Println(err)
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
			return
		}
	}
	service := c.PostForm("service_id")
	if service != "" {
		input.ServiceID, err = uuid.Parse(service)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, errors.New("invalid input body"))
			return
		}
	}
	input.Description = c.PostForm("description")
	fmt.Println(input)
	application, err := h.services.application.CreateApplication(c, input)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, application)
}

func (h *Handler) renderListOfApplications(c *gin.Context) {
	email := c.Request.Context().Value("email").(string)

	user, err := h.services.user.GetUserByEmail(c, email)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}

	apps, err := h.services.application.ListUsersApplications(c, user.UserID)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}

	c.HTML(http.StatusOK, "listofuserapps.html", gin.H{
		"Applications": apps,
	})
}

func (h *Handler) renderAllApps(c *gin.Context) {
	// Fetch all applications from the service layer
	apps, err := h.services.application.ListAllApplications(c)
	if err != nil {
		fmt.Println(err)
		c.Abort()
		return
	}

	// Filter the applications to include only those with status "pending"
	var pendingApps []types.Application // Replace `YourAppType` with the actual type of `apps`
	for _, app := range apps {
		if app.Status == "pending" { // Assuming `Status` is a field in your app struct
			pendingApps = append(pendingApps, *app)
		}
	}

	// Render the filtered list of applications
	c.HTML(http.StatusOK, "listofallapps.html", gin.H{
		"Applications": pendingApps,
	})
}

func (h *Handler) getApplicationDetails(c *gin.Context) {
	// Extract the application ID from the URL path
	appID := c.Param("id")

	if appID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	appID = utils.ReverseUUID(appID)

	appUUID, err := uuid.Parse(appID)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid application ID"))
		return
	}
	app, err := h.services.application.GetApplicationByID(c, appUUID)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusNotFound, errors.New("application not found"))
		return
	}

	c.HTML(http.StatusOK, "adminapplication.html", gin.H{
		"Application": app,
	})
}

func (h *Handler) workApplication(c *gin.Context) {
	appID := c.Param("id")

	var input types.WorkApplicationRequest

	if appID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	input.Status = c.PostForm("status")
	if input.Status == "rejected" {
		input.RejectionReason = c.PostForm("rejection_reason")
	}

	appID = utils.ReverseUUID(appID)

	appUUID, err := uuid.Parse(appID)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid application ID"))
		return
	}
	input.ID = appUUID
	err = h.services.application.WorkApplication(c, input)
	if err != nil {
		fmt.Println(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, input)
}
