package backend_test

import (
	"fmt"
	"net"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"go.uber.org/zap"
)

func assertEqual(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf(
			"\033[31m%s:%d:\n\n\texp: %#v\n\n\tact: %#v\033[39m\n\n",
			filepath.Base(file), line, exp, act,
		)
		tb.FailNow()
	}
}

func keysFor(fields ...zap.Field) []string {
	keys := make([]string, len(fields))
	for i, field := range fields {
		keys[i] = field.Key
	}
	return keys
}

type logEntry struct {
	msg  string
	keys []string
}

type FakeGrpcServer struct {
	serve        func(net.Listener) error
	gracefulStop func()
	stop         func()
}

func (m FakeGrpcServer) Serve(l net.Listener) error { return m.serve(l) }
func (m FakeGrpcServer) GracefulStop()              { m.gracefulStop() }
func (m FakeGrpcServer) Stop()                      { m.stop() }

type FakeAddr struct {
	network func() string
	stringf func() string
}

func (m FakeAddr) Network() string { return m.network() }
func (m FakeAddr) String() string  { return m.stringf() }

type FakeListener struct {
	accept func() (net.Conn, error)
	closef func() error
	addr   func() net.Addr
}

func (m FakeListener) Accept() (net.Conn, error) { return m.accept() }
func (m FakeListener) Close() error              { return m.closef() }
func (m FakeListener) Addr() net.Addr            { return m.addr() }
