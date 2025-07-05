package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/p1xray/pxr-qrcode/internal/config"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	cfg := config.MustLoad()

	fmt.Printf("%+v\n", cfg)
}
