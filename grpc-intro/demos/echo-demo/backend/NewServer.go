package backend

import (
	"backend/api"
	"os"

	"os/signal"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func NewServer(
	logger *zap.Logger,
	echoServiceServer api.EchoServiceServer,
) *Server {
	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_validator.UnaryServerInterceptor(),
			grpc_zap.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_validator.StreamServerInterceptor(),
			grpc_zap.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(),
		),
	)

	api.RegisterEchoServiceServer(grpcServer, echoServiceServer)

	return &Server{
		Exit:       os.Exit,
		Fatal:      logger.Fatal,
		Info:       logger.Info,
		GrpcServer: grpcServer,
		Notify:     signal.Notify,
	}
}
