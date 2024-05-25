package app

import (
	"context"
	"embed"
	"log"
	"os"
	"strings"

	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	AppAddress string `env:"APP_ADDRESS"`
	Postgres   struct {
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		Host     string `env:"POSTGRES_HOSTNAME"`
		Port     string `env:"POSTGRES_PORT"`
		DB       string `env:"POSTGRES_DB"`
	}
	RedisURI string `env:"REDIS_URI"`
}

// NewConfig is to setup app's config
func NewConfig(embedFs embed.FS) *AppConfig {
	// load file from embed
	envs, err := embedFs.ReadFile("configs/local.env")
	if err != nil {
		log.Fatalf("failed to load .env file from embed.Fs. err=%v", err)
	}

	// load envs to runtime(?) line by line
	lines := strings.Split(string(envs), "\n")
	for _, line := range lines {
		if line != "" {
			splits := strings.SplitN(line, "=", 2)

			// skip env if its defined already
			if os.Getenv(splits[0]) != "" {
				continue
			}
			if err := os.Setenv(splits[0], splits[1]); err != nil {
				log.Fatalf("failed to inject .env values. env=%s. value%s. err=%v", splits[0], splits[1], err)
			}
		}
	}

	// read all required env values
	var cfg AppConfig
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.Fatalf("failed to read environment variables. err=%v", err)
	}

	return &cfg
}

var Config *AppConfig
