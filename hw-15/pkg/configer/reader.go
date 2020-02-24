package configer

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
	"github.com/spf13/pflag"
	"log"
)

func ReadConfig() *Config {
	var configPath string
	pflag.StringVarP(&configPath, "config", "c", "configs/config.json", "Config file path")
	pflag.Parse()

	loader := confita.NewLoader(
		file.NewBackend(configPath),
	)
	cfg := Config{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}
