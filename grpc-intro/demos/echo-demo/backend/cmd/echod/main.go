package main

import (
	"backend"
	"fmt"
	"log"
	"net"
	"os"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("logger creation failed: %v", err)
	}

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	l, err := net.Listen("tcp", port)
	if err != nil {
		logger.Fatal("listen failed", zap.String("port", port), zap.Error(err))
	}
	defer l.Close()

	echoServiceServer := &backend.EchoServiceServer{}

	server := backend.NewServer(logger, echoServiceServer)
	server.Serve(l)
}
