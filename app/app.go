package app

import (
	"bufio"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func LoadConfigFrom(file string) Config {
	cfg, err := load(file)
	if err != nil {
<<<<<<< Updated upstream
		log.Fatalf("app: load config: %v", err)
	}
	if !cfg.isValid() {
		log.Fatal("app: load config: invalid config")
=======
		log.Fatalf("app: load file: %v", err)
	}
	if !cfg.isValid() {
		log.Fatalf("app: invalid config: %+v", cfg)
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
	log.Printf("app: got signal %q. Stopping...", sig)
=======
	log.Printf("wait for exit: got signal %q. Stopping...", sig)
>>>>>>> Stashed changes
	onExit()
	os.Exit(0)
}
