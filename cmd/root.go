package cmd

import (
	"github.com/rls/ping-api/pkg/config"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd is the root command of ping-api
var RootCmd = &cobra.Command{
	Use:   "ping-api",
	Short: "ping-api is a realtime location tracker service",
	Long:  "ping-api is a realtime location tracker service",
}

// Execute executes the root command
func Execute() {
	config.Init()

	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yml", "config file")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		return
	}
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
}
