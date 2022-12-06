package controller

import (
	"go-rest-template/internal/domain"
	"go-rest-template/internal/service"
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
	var p domain.Payload
	if !p.IsValid() {
		ctx.String(http.StatusBadRequest, "Data incorrecta")
		return
	}
	l := c.svc.Create(p)
	ctx.JSON(http.StatusCreated, l)
}