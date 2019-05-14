package cmd

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rls/ping-api/conn/cache"
	"github.com/rls/ping-api/pkg/config"
	"github.com/rls/ping-api/router"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Long:  `Start the http-json API server of location-ping-api`,
	Run:   serve,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if err := cache.Connect(config.AppCfg().CacheType); err != nil {
			log.Fatal("Error happened while connecting to caching server, reason", err)
		}
		return nil
	},
}

func init() {
	serveCmd.PersistentFlags().IntP("p", "p", 8080, "port on which the server will listen for http")
	viper.BindPFlag("app.http_port", serveCmd.PersistentFlags().Lookup("p"))
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	appCfg := config.AppCfg()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGKILL, syscall.SIGINT, syscall.SIGQUIT)

	htsrvr := &http.Server{
		ReadTimeout:  appCfg.ReadTimeout,
		WriteTimeout: appCfg.WriteTimeout,
		IdleTimeout:  appCfg.IdleTimeout,
		Addr:         ":" + strconv.Itoa(appCfg.HTTPPort),
		Handler:      router.Route(),
	}

	go func() {
		log.Println("HTTP: Listening on port " + strconv.Itoa(appCfg.HTTPPort))
		log.Fatal(htsrvr.ListenAndServe())
	}()

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	log.Println("Shutting down server...")
	htsrvr.Shutdown(ctx)
	log.Println("Server shutdown gracefully")
	cancel()
}
