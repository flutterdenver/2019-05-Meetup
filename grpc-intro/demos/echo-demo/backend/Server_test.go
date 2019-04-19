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
			fatals            []logEntry
			infos             []logEntry
			nowCalls          int
			sinceCalls        int
			notifyCalls       int
			serveCalls        int
			gracefulStopCalls int
		}{
			fatals:            []logEntry{},
			infos:             []logEntry{},
			nowCalls:          0,
			sinceCalls:        0,
			notifyCalls:       0,
			serveCalls:        0,
			gracefulStopCalls: 0,
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
				observed.sinceCalls++
				return time.Minute
			},
			Notify: func(ch chan<- os.Signal, sigs ...os.Signal) {
				observed.notifyCalls++
				ch <- syscall.SIGINT
			},
			GrpcServer: MockGrpcServer{
				serve: func(net.Listener) error {
					observed.serveCalls++
					return nil
				},
				gracefulStop: func() {
					observed.gracefulStopCalls++
				},
			},
		}

		l := MockListener{
			addr: func() net.Addr {
				return MockAddr{stringf: func() string { return ":9000" }}
			},
		}

		s.Serve(l)

		expected := struct {
			fatals            []logEntry
			infos             []logEntry
			nowCalls          int
			sinceCalls        int
			notifyCalls       int
			serveCalls        int
			gracefulStopCalls int
		}{
			fatals: []logEntry{},
			infos: []logEntry{
				{msg: "starting", keys: []string{"addr"}},
				{msg: "started", keys: []string{}},
				{msg: "stopping", keys: []string{"sig"}},
				{msg: "stopped", keys: []string{"uptime"}},
			},
			nowCalls:          1,
			sinceCalls:        1,
			notifyCalls:       1,
			serveCalls:        1,
			gracefulStopCalls: 1,
		}

		assertEqual(t, expected, observed)
	})
}
