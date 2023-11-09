package app

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func LoadConfig(file string) Config {
	cfg, err := load(file)
	if err != nil {
		log.Fatalf("app: load file: %v", err)
	}
	if !cfg.isValid() {
		log.Fatalf("app: invalid config: %+v", cfg)
	}
	log.Printf("config: %+v", cfg)
	return cfg
}

func DefaultRead() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func WaitForExit(onExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c

	log.Printf("wait for exit: got signal %q. Stopping...", sig)
	onExit()
	os.Exit(0)
}
