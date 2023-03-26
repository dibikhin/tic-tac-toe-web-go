package app

import (
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	GameServer
	Server
}

type GameServer struct {
	Port uint16 `env:"GAME_SERVER_SERVER_PORT"`
}

type Server struct {
	Address   string        `env:"GAME_CLIENT_SERVER_ADDRESS"`
	Timeout   time.Duration `env:"GAME_CLIENT_SERVER_TIMEOUT"`
	LoopDelay time.Duration `env:"GAME_CLIENT_LOOP_DELAY"`
}

func load(path string) (Config, error) {
	if err := godotenv.Load(path); err != nil {
		return Config{}, err
	}
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
