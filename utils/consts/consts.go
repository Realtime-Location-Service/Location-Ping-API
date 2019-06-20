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

// LoggerType ...
type LoggerType string

const (
	// KitLogger ...
	KitLogger LoggerType = "kitlogger"
	// Logrus ...
	Logrus LoggerType = "logrus"
)

// SupportedLogger ...
var SupportedLogger = map[LoggerType]bool{
	KitLogger: true,
}

// QueueType ...
type QueueType string

// RabbitMQ ...
const (
	RabbitMQ QueueType = "rabbitmq"
)

// content types
const (
	JSONContent = "application/json; utf8"
)
