package backend_test

import (
	"backend"
	"errors"
	"net"
	"os"
	"syscall"
	"testing"

	"go.uber.org/zap"
)

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		type notify []os.Signal

		observed := struct {
			fatals        []logEntry
			infos         []logEntry
			exits         []int
			notifies      []notify
			serves        int
			gracefulStops int
			stops         int
		}{
			fatals:        []logEntry{},
			infos:         []logEntry{},
			exits:         []int{},
			notifies:      []notify{},
			serves:        0,
			gracefulStops: 0,
			stops:         0,
		}

		s := backend.Server{
			Exit: func(code int) {
				observed.exits = append(observed.exits, code)
			},
			Fatal: func(msg string, fields ...zap.Field) {
				observed.fatals = append(observed.fatals, logEntry{
					msg:  msg,
					keys: keysFor(fields...),
				})
			},
			Info: func(msg string, fields ...zap.Field) {
				observed.infos = append(observed.infos, logEntry{
					msg:  msg,
					keys: keysFor(fields...),
				})
			},
			Notify: func(ch chan<- os.Signal, sigs ...os.Signal) {
				observed.notifies = append(observed.notifies, sigs)
			},
			GrpcServer: FakeGrpcServer{
				serve: func(net.Listener) error {
					observed.serves++
					return errors.New("no dice")
				},
				gracefulStop: func() {
					observed.gracefulStops++
				},
				stop: func() {
					observed.stops++
				},
			},
		}

		s.Serve(FakeListener{
			addr: func() net.Addr {
				return FakeAddr{stringf: func() string { return ":9000" }}
			},
		})

		expected := struct {
			fatals        []logEntry
			infos         []logEntry
			exits         []int
			notifies      []notify
			serves        int
			gracefulStops int
			stops         int
		}{
			fatals: []logEntry{
				{msg: "serve failed", keys: []string{"error"}},
			},
			infos: []logEntry{
				{msg: "starting", keys: []string{"addr"}},
				{msg: "started", keys: []string{}},
				{msg: "stopped", keys: []string{"uptime"}},
			},
			exits: []int{
				// none because Fatal should be called
			},
			notifies: []notify{
				{syscall.SIGINT},
				{syscall.SIGTERM},
			},
			serves:        1,
			gracefulStops: 0,
			stops:         0,
		}

		assertEqual(t, expected.fatals, observed.fatals)
		assertEqual(t, expected.infos, observed.infos)
		assertEqual(t, expected.exits, observed.exits)
		assertEqual(t, expected.notifies, observed.notifies)

		assertEqual(t, expected, observed)
	})
}
