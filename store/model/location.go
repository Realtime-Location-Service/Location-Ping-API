package model

// Location ..
type Location struct {
	UserID  string  `json:"user_id,omitempty" valid:"-"`
	Lat     float64 `json:"lat" valid:"required"`
	Lon     float64 `json:"lon" valid:"required"`
	Dist    float64 `json:"distance,omitempty" valid:"-"`
	GeoHash int64   `json:"geo_hash,omitempty" valid:"-"`
}
