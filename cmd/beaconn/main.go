package main

import (
	"log"

	"github.com/varun4sid/beaconn-backend/internal/config"
	"github.com/varun4sid/beaconn-backend/internal/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	s := server.NewServer(cfg)
	s.Start()
}
