package consts

// CacheType ...
type CacheType string

const (
	// Redis ...
	Redis CacheType = "redis"
	// Memcached ...
	Memcached CacheType = "memcached"
)

// Radius units
const (
	// Meter unit
	Meter string = "m"
	// KiloMeter unit
	KiloMeter string = "km"
	// Feet unit
	Feet string = "ft"
	// Mile unit
	Mile string = "mi"
)

const (
	// DefaultLocationsLimit will be used if limit is missing
	DefaultLocationsLimit = 100
)

// ValidRadiusUnits ....
var ValidRadiusUnits = map[string]bool{
	Meter:     true,
	KiloMeter: true,
	Feet:      true,
	Mile:      true,
}
