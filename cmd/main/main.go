package main

import (
	"lg/config"
	_ "lg/docs"
	"lg/internal/app"
	"log"
)

// @title Платформа для продвижения иновационных идей
// @version 1.0
// @description Хакатон "Лига Цифровой Трансформации"

// @host localhost:9000
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}
	app.Run(cfg)
}
