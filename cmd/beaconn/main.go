package main

import (
	"fmt"

	"github.com/varun4sid/beaconn-backend/internal/config"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
}