package domain

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Coordinates struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lon"`
}

type Location struct {
	ID          int64
	DeviceID    string
	Timestamp   time.Time
	Battery     int64
	Coordinates Coordinates
}

type LocationDTO struct {
	ID          int64       `json:"id"`
	DeviceID    string      `json:"device_id"`
	Timestamp   int64       `json:"ts"`
	Battery     int64       `json:"bat"`
	Coordinates Coordinates `json:"coords"`
}

type LocationEntity struct {
	ID          int64   `db:"id"`
	DeviceID    string  `db:"device_id"`
	Timestamp   string  `db:"timestamp"`
	Battery     int64   `db:"battery"`
	Coordinates []uint8 `db:"coordinates"`
}

type LocationFilters struct {
	DeviceID string
	From     *time.Time
	To       *time.Time
	Limit    *uint64
	Offset   *uint64
}

func (l *Location) ToEntity() *LocationEntity {
	return &LocationEntity{
		ID:          l.ID,
		DeviceID:    l.DeviceID,
		Timestamp:   l.Timestamp.Format(time.RFC3339),
		Battery:     l.Battery,
		Coordinates: []uint8(fmt.Sprintf("(%f,%f)", *l.Coordinates.Latitude, *l.Coordinates.Longitude)),
	}
}

func (l *LocationEntity) ToLocation() *Location {
	ts, _ := time.Parse(time.RFC3339, l.Timestamp)
	coords := string(l.Coordinates)
	coords = coords[1 : len(coords)-2]
	coordsSplit := strings.Split(coords, ",")
	latitude, _ := strconv.ParseFloat(coordsSplit[0], 64)
	longitude, _ := strconv.ParseFloat(coordsSplit[1], 64)

	return &Location{
		ID:          l.ID,
		DeviceID:    l.DeviceID,
		Timestamp:   ts,
		Battery:     l.Battery,
		Coordinates: Coordinates{Latitude: &latitude, Longitude: &longitude},
	}
}

func (l *Location) ToDTO() *LocationDTO {
	return &LocationDTO{
		ID:          l.ID,
		DeviceID:    l.DeviceID,
		Timestamp:   l.Timestamp.Unix(),
		Battery:     l.Battery,
		Coordinates: l.Coordinates,
	}
}

func (l *LocationDTO) ToLocation() *Location {
	return &Location{
		ID:          l.ID,
		DeviceID:    l.DeviceID,
		Timestamp:   time.Unix(l.Timestamp, 0),
		Battery:     l.Battery,
		Coordinates: l.Coordinates,
	}
}

func (p *LocationDTO) IsValid() bool {
	return p.DeviceID != "" && p.Timestamp > 0 && p.Battery >= 0 && p.Coordinates.Latitude != nil && p.Coordinates.Longitude != nil
}
