// Package config takes care of loading and exposing user configuration.
package config

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"

	log "github.com/sirupsen/logrus"
)

// Config is the exporter configuration.
type Config struct {
	BaseURL        string        `config:"base_url,required,short=b,description=Base URL of Firefly III instance"`
	APIKey         string        `config:"api_key,required,short=a,description=Firefly III API key"`
	HTTPPort       string        `config:"http_port,short=p,description=HTTP port to expose the exporter"`
	ScrapeInterval time.Duration `config:"scrape_interval,short=i,description=scrape interval in seconds"`
}

func getDefaultConfig() *Config {
	return &Config{
		HTTPPort:       "4001",
		ScrapeInterval: 30 * time.Second,
	}
}

// Load method loads the configuration by using environment variables.
func Load() *Config {
	loaders := []backend.Backend{
		env.NewBackend(),
		flags.NewBackend(),
	}

	loader := confita.NewLoader(loaders...)

	cfg := getDefaultConfig()
	err := loader.Load(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

// Show method displays all the load configuration
func (c Config) Show() {
	log.Println("=============================================")
	log.Println("      Firefly III Exporter Configuration     ")
	log.Println("=============================================")

	val := reflect.ValueOf(&c).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		value := fmt.Sprintf("%v", valueField.Interface())

		if typeField.Name == "APIKey" {
			value = "[FILTERED]"
		}

		if value != "" {
			log.Printf("%s: %s", typeField.Name, value)
		}
	}

	log.Println("=============================================")
}
