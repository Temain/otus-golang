package configer

import (
	"context"
	"log"

	"github.com/heetch/confita/backend/env"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

func ReadConfig(path string) *ConfigApi {
	loader := confita.NewLoader(
		file.NewBackend(path),
		env.NewBackend(),
	)
	cfg := ConfigApi{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}
