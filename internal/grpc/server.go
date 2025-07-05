package server

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalError(msg string) error {
	return status.Error(codes.Internal, msg)
}
