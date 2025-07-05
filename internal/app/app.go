package app

import (
	grpcapp "github.com/p1xray/pxr-qrcode/internal/app/grpc"
	"github.com/p1xray/pxr-qrcode/internal/service/qrcode"
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
	qrCodeService := qrcode.New()

	grpcApp := grpcapp.New(log, grpcPort, qrCodeService)

	return &App{
		GRPCServer: grpcApp,
	}
}
