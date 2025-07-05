package qrcode

import (
	"encoding/base64"
	"errors"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
)

var (
	ErrGenerateQRCode = errors.New("error generating QR-code")
)

// Service is a service for working with QR-code.
type Service struct {
	recoveryLevel qrcode.RecoveryLevel
	size          int
}

// New create a new instance of service for working with QR-code.
func New() *Service {
	return &Service{
		recoveryLevel: qrcode.Low,
		size:          256,
	}
}

// GenerateBase64 returns generated QR-code encoded as base64 string.
func (s *Service) GenerateBase64(url string) (string, error) {
	const op = "qrcode.GenerateBase64"

	qrCode, err := qrcode.Encode(url, s.recoveryLevel, s.size)
	if err != nil {
		return "", fmt.Errorf("%s: %w: %w", op, ErrGenerateQRCode, err)
	}

	qrCodeBase64Str := base64.StdEncoding.EncodeToString(qrCode)
	return qrCodeBase64Str, nil
}
