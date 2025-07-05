package qrcodeserver

import (
	"context"
	server "github.com/p1xray/pxr-qrcode/internal/grpc"
	qrcodepb "github.com/p1xray/pxr-qrcode/pkg/grpc/gen/go/qrcode"
	"google.golang.org/grpc"
)

type serverAPI struct {
	qrcodepb.UnimplementedQrCodeServer
}

// Register registers the implementation of the API service with the gRPC server.
func Register(gRPC *grpc.Server) {
	qrcodepb.RegisterQrCodeServer(gRPC, &serverAPI{})
}

func (s *serverAPI) Generate(
	ctx context.Context,
	req *qrcodepb.GenerateRequest,
) (*qrcodepb.GenerateResponse, error) {
	return nil, server.InternalError("internal server error")
}
