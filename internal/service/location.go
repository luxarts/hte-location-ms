package service

import (
	"go-rest-template/internal/domain"
	"go-rest-template/internal/repository"
)

type LocationService interface {
	Create(location domain.Payload) domain.Payload
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (r *locationService) Create(location domain.Payload) domain.Payload {
	r.repo.Create(location)
	return location
}
