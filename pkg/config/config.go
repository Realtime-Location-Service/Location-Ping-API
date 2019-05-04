package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	// this package is necessary to read config from remote consul
	_ "github.com/spf13/viper/remote"
)

// Init initiates of config load
func Init() {
	viper.SetEnvPrefix("ping_api")
	viper.BindEnv("env")
	viper.BindEnv("consul_url")
	viper.BindEnv("consul_path")

	consulURL := viper.GetString("consul_url")
	consulPath := viper.GetString("consul_path")
	if consulURL == "" {
		log.Fatal("CONSUL_URL missing")
	}
	if consulPath == "" {
		log.Fatal("CONSUL_PATH missing")
	}

	viper.AddRemoteProvider("consul", consulURL, consulPath)
	viper.SetConfigType("yml")
	err := viper.ReadRemoteConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("%s named \"%s\"", err.Error(), consulPath))
	}

	LoadApp()
}
