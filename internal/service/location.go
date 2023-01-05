package service

import (
	"hte-location-ms/internal/domain"
	"hte-location-ms/internal/repository"
)

type LocationService interface {
	Create(dto *domain.LocationDTO) (*domain.Location, error)
	GetLocationsByDeviceID(id int64, filter *domain.LocationFilters) (*[]domain.Location, error)
}

type locationService struct {
	repo repository.LocationRepository
}

func NewLocationService(repo repository.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (r *locationService) Create(dto *domain.LocationDTO) (*domain.Location, error) {
	l := dto.ToLocation()

	le, err := r.repo.Create(l.ToEntity())
	l = le.ToLocation()

	return l, err
}

func (r *locationService) GetLocationsByDeviceID(id int64, filter *domain.LocationFilters) (*[]domain.Location, error) {
	les, err := r.repo.GetLocationsByDeviceID(id, filter)
	var ls []domain.Location
	for _, le := range *les {
		ls = append(ls, *le.ToLocation())
	}
	return &ls, err
}
