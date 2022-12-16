package controller

import (
	"hte-location-ms/internal/domain"
	"hte-location-ms/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LocationController interface {
	Create(ctx *gin.Context)
}
type locationController struct {
	svc service.LocationService
}

func NewLocationController(svc service.LocationService) LocationController {
	return &locationController{svc: svc}
}

func (c *locationController) Create(ctx *gin.Context) {
	var p domain.Location
	err := ctx.ShouldBindJSON(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}
	if !p.IsValid() {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid payload"})
		return
	}
	l, err := c.svc.Create(&p)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, l)
}
