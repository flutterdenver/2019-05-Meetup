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

type MockGrpcServer struct {
	serve        func(net.Listener) error
	gracefulStop func()
}

func (m MockGrpcServer) Serve(l net.Listener) error { return m.serve(l) }
func (m MockGrpcServer) GracefulStop()              { m.gracefulStop() }

type MockAddr struct {
	network func() string
	stringf func() string
}

func (m MockAddr) Network() string { return m.network() }
func (m MockAddr) String() string  { return m.stringf() }

type MockListener struct {
	accept func() (net.Conn, error)
	closef func() error
	addr   func() net.Addr
}

func (m MockListener) Accept() (net.Conn, error) { return m.accept() }
func (m MockListener) Close() error              { return m.closef() }
func (m MockListener) Addr() net.Addr            { return m.addr() }
