package grpcapp

import (
	"fmt"
	"log"
	"log/slog"
	"net"

	"github.com/evgenmar/sso/internal/grpc/authgrpc"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(
	log *slog.Logger,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	authgrpc.Register(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}

// Run runs gRPC server.
func (a *App) Run() error {
	const op = "grpcapp.Run"

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	a.log.Info("gRPC server started",
		slog.String("op", op),
		slog.String("addr", l.Addr().String()),
	)

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Stop stops gRPC server.
func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.Info("stopping gRPC server", slog.String("op", op))

	a.gRPCServer.GracefulStop()
}
