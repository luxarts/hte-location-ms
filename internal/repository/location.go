package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"hte-location-ms/internal/domain"
	"log"
	"time"
)

type LocationRepository interface {
	Create(p *domain.LocationEntity) (*domain.LocationEntity, error)
	GetLocationsByDeviceID(id int64, filter *domain.LocationFilters) (*[]domain.LocationEntity, error)
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

func (r *locationRepository) Create(p *domain.LocationEntity) (*domain.LocationEntity, error) {
	query, args, err := r.sqlBuilder.CreateSQL(p)
	if err != nil {
		return p, err
	}
	_, err = r.db.Exec(query, args...)
	return p, err
}

func (r *locationRepository) GetLocationsByDeviceID(id int64, filter *domain.LocationFilters) (*[]domain.LocationEntity, error) {
	query, args, err := r.sqlBuilder.GetLocationsByDeviceIDSQL(id, filter)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	var results []domain.LocationEntity
	for rows.Next() {
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		c, _ := rows.SliceScan()
		log.Println(c)
		var l domain.LocationEntity
		err = rows.StructScan(&l)
		if err != nil {
			return nil, err
		}
		results = append(results, l)
	}

	return &results, nil
}

//--------------------------

type locationsSQL struct {
	table string
}

func (s *locationsSQL) CreateSQL(l *domain.LocationEntity) (string, []interface{}, error) {
	query, args, err := squirrel.Insert(s.table).
		Columns("device_id", "battery", "timestamp", "coordinates").
		Values(l.DeviceID, l.Battery, l.Timestamp, l.Coordinates).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	return query, args, err
}

func (s *locationsSQL) GetLocationsByDeviceIDSQL(id int64, filter *domain.LocationFilters) (string, []interface{}, error) {
	and := squirrel.And{
		squirrel.Eq{"device_id": id},
	}
	if filter.From != nil {
		and = append(and, squirrel.GtOrEq{"timestamp": filter.From.Format(time.RFC3339)})
	}
	if filter.To != nil {
		and = append(and, squirrel.LtOrEq{"timestamp": filter.To.Format(time.RFC3339)})
	}
	query, args, err := squirrel.Select("*").
		From(s.table).
		Where(and).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	return query, args, err
}
