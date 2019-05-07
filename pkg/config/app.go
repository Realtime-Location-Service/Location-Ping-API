package config

import (
	"time"

	"github.com/rls/ping-api/utils/consts"
	"github.com/spf13/viper"
)

// App holds the app configuration
type App struct {
	Name         string
	Debug        bool
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
	CacheType    consts.CacheType
}

var app = &App{}

// AppCfg returns the app configuration
func AppCfg() *App {
	return app
}

// LoadAppCfg loads app configuration
func LoadAppCfg() {
	app.Debug = viper.GetBool("app.debug")
	app.CacheType = consts.CacheType(viper.GetString("app.cache_type"))
	app.HTTPPort = viper.GetInt("app.http_port")
	app.ReadTimeout = viper.GetDuration("app.read_timeout") * time.Second
	app.WriteTimeout = viper.GetDuration("app.write_timeout") * time.Second
	app.IdleTimeout = viper.GetDuration("app.idle_timeout") * time.Second
	app.Name = viper.GetString("app.name")
}
