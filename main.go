// Package main takes care of the initialization of the application.
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/kinduff/firefly_iii_exporter/config"
	"github.com/kinduff/firefly_iii_exporter/internal/collector"
	"github.com/kinduff/firefly_iii_exporter/internal/metrics"
	"github.com/kinduff/firefly_iii_exporter/internal/server"

	log "github.com/sirupsen/logrus"
)

var (
	s *server.Server
)

func main() {
	cfg := config.Load()

	cfg.Show()

	metrics.Init(cfg)

	client := collector.NewCollector(cfg)
	go client.Scrape()

	initHTTPServer(cfg.HTTPPort)

	handleExitSignal()
}

func initHTTPServer(port string) {
	s = server.NewServer(port)
	go s.ListenAndServe()
}

func handleExitSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	s.Stop()
	log.Fatal("HTTP server stopped")
}
