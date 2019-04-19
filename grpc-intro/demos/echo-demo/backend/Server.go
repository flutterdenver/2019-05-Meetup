package backend

import (
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"go.uber.org/zap"
)

// NOTES on signaling:
// - SIGKILL and SIGSTOP may not be caught by a Go program
// - Use SIGINT to stop this server gracefully.
// - Use SIGTERM to stop this server immediately.
// - docker kill is your friend; eg:
// - docker kill --signal=SIGINT CONTAINER
// - docker kill --signal=SIGTERM CONTAINER

/*
Scenario A: Stopping gracefully with SIGINT

Terminal A:
$ docker run flutterdenver/echod:latest
{"level":"info","ts":1555693681.773813,"caller":"backend/Server.go:35","msg":"starting","cmd":"echod","pid":1,"addr":"[::]:9000"}
{"level":"info","ts":1555693681.7747886,"caller":"backend/Server.go:73","msg":"started","cmd":"echod","pid":1}
{"level":"info","ts":1555693703.7016788,"caller":"backend/Server.go:56","msg":"SIGINT received: stopping gracefully","cmd":"echod","pid":1}
{"level":"info","ts":1555693703.7022948,"caller":"backend/Server.go:39","msg":"stopped","cmd":"echod","pid":1,"uptime":21.9285335}

Terminal B:
$ docker ps
CONTAINER ID        IMAGE                        COMMAND             CREATED             STATUS              PORTS                    NAMES
d2a19b033f6c        flutterdenver/echod:latest   "/echod"            7 seconds ago       Up 6 seconds        0.0.0.0:9000->9000/tcp   ecstatic_ptolemy

$ docker kill --signal=SIGINT d2a19b033f6c
d2a19b033f6c

*/

/*
Scenario B: Stopping immediately with SIGTERM

Terminal A:
$ docker run flutterdenver/echod:latest
{"level":"info","ts":1555693719.9574184,"caller":"backend/Server.go:35","msg":"starting","cmd":"echod","pid":1,"addr":"[::]:9000"}
{"level":"info","ts":1555693719.9581058,"caller":"backend/Server.go:73","msg":"started","cmd":"echod","pid":1}
{"level":"info","ts":1555693737.1364677,"caller":"backend/Server.go:66","msg":"SIGTERM received: stopping immediately","cmd":"echod","pid":1}
{"level":"info","ts":1555693737.1367035,"caller":"backend/Server.go:39","msg":"stopped","cmd":"echod","pid":1,"uptime":17.1792851}

Terminal B:
$ docker ps
CONTAINER ID        IMAGE                        COMMAND             CREATED             STATUS              PORTS                    NAMES
cff7e36e479f        flutterdenver/echod:latest   "/echod"            6 seconds ago       Up 5 seconds        0.0.0.0:9000->9000/tcp   lucid_babbage

$ docker kill --signal=SIGTERM cff7e36e479f
*/

type Server struct {
	Exit       func(int)                            // os.Exit
	Fatal      func(string, ...zap.Field)           // zap.Logger.Fatal
	Info       func(string, ...zap.Field)           // zap.Logger.Info
	Notify     func(chan<- os.Signal, ...os.Signal) // signal.Notify
	GrpcServer interface {                          // *grpc.Server
		Serve(net.Listener) error
		GracefulStop()
		Stop()
	}
}

func (s *Server) Serve(l net.Listener) {
	startedAt := time.Now()

	s.Info("starting", zap.String("addr", l.Addr().String()))

	// last thing we do before exiting; not safe to defer
	logStopped := func() {
		s.Info("stopped", zap.Duration("uptime", time.Since(startedAt)))
	}

	var wg sync.WaitGroup

	// Use SIGINT to stop gracefully.
	sigInt := make(chan os.Signal, 1)
	s.Notify(sigInt, syscall.SIGINT)

	// Use SIGTERM stop immediately.
	sigTerm := make(chan os.Signal, 1)
	s.Notify(sigTerm, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		wg.Done()
		<-sigInt
		s.Info("SIGINT received: stopping gracefully")
		s.GrpcServer.GracefulStop()
		logStopped()
		s.Exit(0)
	}()

	wg.Add(1)
	go func() {
		wg.Done()
		<-sigTerm
		s.Info("SIGTERM received: stopping immediately")
		s.GrpcServer.Stop()
		logStopped()
		s.Exit(0)
	}()

	wg.Wait() // make sure we're receiving signals, then start the server
	s.Info("started")

	if err := s.GrpcServer.Serve(l); err != nil {
		logStopped()
		s.Fatal("serve failed", zap.Error(err))
	} else {
		logStopped()
		s.Exit(0)
	}
}
