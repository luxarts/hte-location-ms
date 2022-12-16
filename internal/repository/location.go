package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"hte-location-ms/internal/domain"
	"log"
	"time"
)

type LocationRepository interface {
	Create(p *domain.Location)
}

type locationRepository struct {
	db *sqlx.DB
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) Create(p *domain.Location) {
	log.Println(p)
	tss := time.Unix(p.Timestamp, 0).Format(time.RFC3339)
	coords := fmt.Sprintf("(%f,%f)", *p.Coordinates.Latitude, *p.Coordinates.Longitude)
	_, err := r.db.Exec("INSERT INTO hte.locations (device_id, battery, timestamp,coordinates) VALUES ($1,$2,$3,$4)", p.DeviceID, p.Battery, tss, coords)
	if err != nil {
		log.Println(err)
		return
	}
}
