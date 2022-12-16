package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/luxarts/jsend-go"
	"hte-location-ms/internal/controller"
	"hte-location-ms/internal/defines"
	"hte-location-ms/internal/repository"
	"hte-location-ms/internal/service"
	"log"
	"net/http"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init
	db, err := sqlx.Open("postgres", "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	// Repositories init
	repo := repository.NewLocationRepository(db)

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
