package model

// Location ..
type Location struct {
	UserID  string  `json:"userId" valid:"required"`
	Lat     float64 `json:"lat" valid:"required"`
	Lon     float64 `json:"lon" valid:"required"`
	Dist    float64 `json:"dist" valid:"-"`
	GeoHash int64   `json:"geo_hash" valid:"-"`
}
