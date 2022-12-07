package repository

import (
	"go-rest-template/internal/domain"
	"log"
)

type LocationRepository interface {
	Create(p *domain.Payload)
}

type locationRepository struct {
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{}
}

func (r *locationRepository) Create(p *domain.Payload) {
	log.Println(p)
}
