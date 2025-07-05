package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// QRCodeService is a service for working with QR-code.
type QRCodeService interface {
	// GenerateBase64 returns generated QR-code encoded as base64 string.
	GenerateBase64(ctx context.Context, url string) (string, error)
}

func InternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}

func InvalidArgumentError(msg string) error {
	return status.Error(codes.InvalidArgument, msg)
}
