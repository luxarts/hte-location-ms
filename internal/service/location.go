package service

import (
	"go-rest-template/internal/domain"
	"go-rest-template/internal/repository"
)

type LocationService interface {
	Create(p *domain.Payload) *domain.Payload
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (r *locationService) Create(p *domain.Payload) *domain.Payload {
	r.repo.Create(p)
	return p
}
