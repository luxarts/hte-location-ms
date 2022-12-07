package domain

type Coordinates struct {
	Latitude  *float64 `json:"lat"`
	Longitude *float64 `json:"lon"`
}

type Location struct {
	DeviceID    string      `json:"device_id"`
	Timestamp   int64       `json:"ts"`
	Battery     int64       `json:"bat"`
	Coordinates Coordinates `json:"coords"`
}

func (p *Location) IsValid() bool {
	return p.DeviceID != "" && p.Timestamp > 0 && p.Battery >= 0 && p.Coordinates.Latitude != nil && p.Coordinates.Longitude != nil
}
