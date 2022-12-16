package repository

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"hte-location-ms/internal/domain"
	"log"
	"time"
)

type LocationRepository interface {
	Create(p *domain.Location) error
}

type locationRepository struct {
	db         *sqlx.DB
	sqlBuilder locationsSQL
}

func NewLocationRepository(db *sqlx.DB) LocationRepository {
	return &locationRepository{
		db: db,
		sqlBuilder: locationsSQL{
			table: "hte.locations",
		},
	}
}

func (r *locationRepository) Create(p *domain.Location) error {
	log.Println(p)
	query, args, err := r.sqlBuilder.CreateSQL(p)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(query, args...)
	return err
}

type locationsSQL struct {
	table string
}

func (s *locationsSQL) CreateSQL(l *domain.Location) (string, []interface{}, error) {
	tss := time.Unix(l.Timestamp, 0).Format(time.RFC3339)
	coords := fmt.Sprintf("(%f,%f)", *l.Coordinates.Latitude, *l.Coordinates.Longitude)
	query, args, err := squirrel.Insert(s.table).
		Columns("device_id", "battery", "timestamp", "coordinates").
		Values(l.DeviceID, l.Battery, tss, coords).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	return query, args, err
}
