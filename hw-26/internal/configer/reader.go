package configer

import (
	"context"
	"log"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/file"
)

func ReadConfigApi(path string) *ConfigApi {
	loader := confita.NewLoader(
		file.NewBackend(path),
	)
	cfg := ConfigApi{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}

func ReadConfigScheduler(path string) *ConfigScheduler {
	loader := confita.NewLoader(
		file.NewBackend(path),
	)
	cfg := ConfigScheduler{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}

func ReadConfigSender(path string) *ConfigSender {
	loader := confita.NewLoader(
		file.NewBackend(path),
	)
	cfg := ConfigSender{}
	err := loader.Load(context.Background(), &cfg)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}

	return &cfg
}
