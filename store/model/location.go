package model

// Location ..
type Location struct {
	UserID             string  `json:"user_id,omitempty" valid:"required"`
	Lat                float64 `json:"lat" valid:"required,latitude"`
	Lon                float64 `json:"lon" valid:"required,longitude"`
	Dist               float64 `json:"distance,omitempty" valid:"-"`
	GeoHash            int64   `json:"geo_hash,omitempty" valid:"-"`
	ClientTimestampUTC int64   `json:"client_timestamp_utc,omitempty" valid:"required"`
	ServerTimestampUTC int64   `json:"server_timestamp_utc,omitempty"`
	Domain             string  `json:"domain,omitempty"`
}
