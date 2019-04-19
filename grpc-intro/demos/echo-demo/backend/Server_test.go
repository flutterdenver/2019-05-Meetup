package backend_test

import (
	"backend"
	"net"
	"os"
	"syscall"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestServer(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		observed := struct {
			fatals   []logEntry
			infos    []logEntry
			nowCalls int
		}{
			fatals:   []logEntry{},
			infos:    []logEntry{},
			nowCalls: 0,
		}

		s := backend.Server{
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
			Now: func() time.Time {
				observed.nowCalls++
				return time.Now()
			},
			Since: func(time.Time) time.Duration {
				return time.Minute
			},
			Notify: func(ch chan<- os.Signal, sigs ...os.Signal) {
				ch <- syscall.SIGINT
			},
			GrpcServer: MockGrpcServer{
				serve: func(net.Listener) error {
					return nil
				},
				gracefulStop: func() {
				},
			},
		}

		l := MockListener{
			addr: func() net.Addr {
				return MockAddr{
					stringf: func() string { return ":9000" },
				}
			},
		}

		s.Serve(l)

		expected := struct {
			fatals   []logEntry
			infos    []logEntry
			nowCalls int
		}{
			fatals: []logEntry{},
			infos: []logEntry{
				{msg: "starting", keys: []string{"addr"}},
				{msg: "started", keys: []string{}},
				{msg: "stopping", keys: []string{"sig"}},
				{msg: "stopped", keys: []string{"uptime"}},
			},
			nowCalls: 1,
		}

		assertEqual(t, expected.fatals, observed.fatals)
		assertEqual(t, expected.infos, observed.infos)
		assertEqual(t, expected.nowCalls, observed.nowCalls)
	})
}
