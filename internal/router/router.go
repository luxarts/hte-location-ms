package router

import (
	"go-rest-template/internal/controller"
	"go-rest-template/internal/defines"
	"go-rest-template/internal/repository"
	"go-rest-template/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luxarts/jsend-go"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init

	// Repositories init
	repo := repository.NewLocationRepository()

	// Services init
	svc := service.NewLocationService(repo)

	// Controllers init
	ctrl := controller.NewLocationController(svc)

	// Endpoints
	r.POST(defines.EndpointCreateLocation, ctrl.Create)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}
