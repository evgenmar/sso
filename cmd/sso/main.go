package main

import (
	"fmt"

	"github.com/evgenmar/sso/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO init logger

	// TODO init app

	// TODO run gRPC server
}
