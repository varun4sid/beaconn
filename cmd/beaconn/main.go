package main

import (
	"github.com/varun4sid/beaconn-backend/internal/server"
	"github.com/varun4sid/beaconn-backend/internal/config"
)

func main() {
	config.Load()

	s := server.NewServer()
	s.Start()
}