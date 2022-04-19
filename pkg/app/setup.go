package app

import (
	"bufio"
	"log"
	"os"
	"strings"
	"tictactoe/pkg/config"
)

func LoadConfig() config.Config {
	file := ".env"
	cfg, err := config.Load(file)
	if err != nil {
		log.Fatalf("config: %v", err)
	}
	log.Printf("%+v", cfg)
	return cfg
}

func DefaultRead() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}
