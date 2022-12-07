package repository

import (
	"go-rest-template/internal/domain"
	"log"
)

type LocationRepository interface {
	Create(location domain.Payload)
}

type locationRepository struct {
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{}
}

func (r *locationRepository) Create(location domain.Payload) {
	log.Println(location)
}
