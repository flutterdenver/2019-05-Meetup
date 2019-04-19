package backend

import (
	"net"
	"os"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type Server struct {
	Fatal      func(string, ...zap.Field)           // zap.Logger.Fatal
	Info       func(string, ...zap.Field)           // zap.Logger.Info
	Now        func() time.Time                     // time.Now
	Since      func(time.Time) time.Duration        // time.Since
	Notify     func(chan<- os.Signal, ...os.Signal) // signal.Notify
	GrpcServer interface {                          // *grpc.Server
		Serve(net.Listener) error
		GracefulStop()
	}
}

func (s *Server) Serve(l net.Listener) {
	defer func(t time.Time) {
		s.Info("stopped", zap.Duration("uptime", s.Since(t)))
	}(s.Now())

	s.Info("starting", zap.String("addr", l.Addr().String()))
	go func() {
		if err := s.GrpcServer.Serve(l); err != nil {
			s.Fatal("serve failed", zap.Error(err))
		}
	}()

	c := make(chan os.Signal, 1)
	s.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	s.Info("started")
	sig := <-c
	s.Info("stopping", zap.Any("sig", sig))
	s.GrpcServer.GracefulStop()
}
