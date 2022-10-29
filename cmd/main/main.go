package main

import (
	"lg/config"
	"lg/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}
	app.Run(cfg)
}
