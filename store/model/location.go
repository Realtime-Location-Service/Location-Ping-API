package model

// Location ..
type Location struct {
	UserID string  `json:"user_id" valid:"required"`
	Lat    float64 `json:"lat" valid:"required"`
	Lon    float64 `json:"lon" valid:"required"`
}
