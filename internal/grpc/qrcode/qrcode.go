package qrcodeserver

import (
	"context"
	server "github.com/p1xray/pxr-qrcode/internal/grpc"
	qrcodepb "github.com/p1xray/pxr-qrcode/pkg/grpc/gen/go/qrcode"
	"google.golang.org/grpc"
)

type serverAPI struct {
	qrcodepb.UnimplementedQrCodeServer
	service server.QRCodeService
}

// Register registers the implementation of the API service with the gRPC server.
func Register(gRPC *grpc.Server, service server.QRCodeService) {
	qrcodepb.RegisterQrCodeServer(gRPC, &serverAPI{service: service})
}

func (s *serverAPI) Generate(
	ctx context.Context,
	req *qrcodepb.GenerateRequest,
) (*qrcodepb.GenerateResponse, error) {
	if err := validateGenerateRequest(req); err != nil {
		return nil, err
	}

	qrCode, err := s.service.GenerateBase64(req.GetUrl())
	if err != nil {
		return nil, server.InternalError("error generating QR-code")
	}

	return &qrcodepb.GenerateResponse{QrCode: qrCode}, nil
}

func validateGenerateRequest(req *qrcodepb.GenerateRequest) error {
	if req.GetUrl() == "" {
		return server.InvalidArgumentError("URL is empty")
	}

	return nil
}
