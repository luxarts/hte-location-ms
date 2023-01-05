package router

import (
	"fmt"
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
	"os"
)

func New() *gin.Engine {
	r := gin.Default()

	mapRoutes(r)

	return r
}

func mapRoutes(r *gin.Engine) {
	// DB connectors, rest clients, and other stuff init
	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable",
		os.Getenv(defines.EnvPostgresUser),
		os.Getenv(defines.EnvPostgresPassword),
		os.Getenv(defines.EnvPostgresHost),
		os.Getenv(defines.EnvPostgresPort),
	)
	db, err := sqlx.Open("postgres", postgresURI)
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
	r.GET(defines.EndpointGetLocationsByDeviceID, ctrl.GetLocationsByDeviceID)

	// Health check endpoint
	r.GET(defines.EndpointPing, healthCheck)
}

func healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, jsend.NewSuccess("pong"))
}
