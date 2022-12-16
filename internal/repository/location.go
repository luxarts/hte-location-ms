package repository

import (
	"hte-location-ms/internal/domain"
	"log"
)

type LocationRepository interface {
	Create(p *domain.Location)
}

type locationRepository struct {
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{}
}

func (r *locationRepository) Create(p *domain.Location) {
	log.Println(p)
}
