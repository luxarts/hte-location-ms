package service

import (
	"hte-location-ms/internal/domain"
	"hte-location-ms/internal/repository"
)

type LocationService interface {
	Create(p *domain.Location) (*domain.Location, error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (r *locationService) Create(p *domain.Location) (*domain.Location, error) {
	err := r.repo.Create(p)
	return p, err
}
