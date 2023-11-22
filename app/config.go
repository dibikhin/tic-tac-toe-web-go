package app

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	GameServer GameServerConfig
	Server     ServerConfig
}

type GameServerConfig struct {
	Port uint16 `env:"GAME_SERVER_SERVER_PORT"`
}

type ServerConfig struct {
	Address   string        `env:"GAME_CLIENT_SERVER_ADDRESS"`
	Timeout   time.Duration `env:"GAME_CLIENT_SERVER_TIMEOUT"`
	LoopDelay time.Duration `env:"GAME_CLIENT_LOOP_DELAY"`
}

func (cfg *Config) isValid() bool {
	allEmpty := cfg.GameServer == GameServerConfig{} && cfg.Server == ServerConfig{}
	return !allEmpty
}

func load(path string) (*Config, error) {
	if err := godotenv.Load(path); err != nil {
		return nil, err
	}
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
