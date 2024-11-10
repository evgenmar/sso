package app

import (
	"log/slog"
	"time"

	"github.com/evgenmar/sso/internal/app/grpcapp"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	//TODO: init storage

	//TODO: init auth service (auth)

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
