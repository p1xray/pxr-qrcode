package app

import (
	grpcapp "github.com/p1xray/pxr-qrcode/internal/app/grpc"
	"log/slog"
)

// App is an application.
type App struct {
	GRPCServer *grpcapp.App
}

// New creates a new application.
func New(
	log *slog.Logger,
	grpcPort int,
) *App {
	// TODO: create service

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
