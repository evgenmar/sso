package app

import (
	"log"
	"log/slog"
	"time"

	"github.com/evgenmar/sso/internal/app/grpcapp"
	"github.com/evgenmar/sso/internal/services/auth"
	"github.com/evgenmar/sso/internal/storage/sqlite"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	l *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := sqlite.New(storagePath)
	if err != nil {
		log.Fatal(err)
	}

	authService := auth.New(l, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(l, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
