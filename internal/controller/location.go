package controller

import (
	"hte-location-ms/internal/domain"
	"hte-location-ms/internal/service"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LocationController interface {
	Create(ctx *gin.Context)
	GetLocationsByDeviceID(ctx *gin.Context)
}
type locationController struct {
	svc service.LocationService
}

func NewLocationController(svc service.LocationService) LocationController {
	return &locationController{svc: svc}
}

func (c *locationController) Create(ctx *gin.Context) {
	var p domain.LocationDTO
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

func (c *locationController) GetLocationsByDeviceID(ctx *gin.Context) {
	idString := ctx.Query("device_id")
	if idString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "missing device_id"})
		return
	}
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid device_id"})
		return
	}
	var filter domain.LocationFilters

	if fromStr, fromExist := ctx.GetQuery("from"); fromExist {
		if from, err := strconv.ParseInt(fromStr, 10, 64); err == nil {
			fromTs := time.Unix(from, 0)
			filter.From = &fromTs
		}
	}
	if toStr, toExist := ctx.GetQuery("to"); toExist {
		if to, err := strconv.ParseInt(toStr, 10, 64); err == nil {
			toTs := time.Unix(to, 0)
			filter.To = &toTs
		}
	}

	resp, err := c.svc.GetLocationsByDeviceID(id, &filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}
