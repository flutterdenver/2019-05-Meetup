package main

import (
	"backend/api"
	"context"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestEchoServiceServer(t *testing.T) {
	scenarios := []struct {
		name   string
		inMsg  *api.EchoMessage // input message
		expMsg *api.EchoMessage // expected message
		expErr error            // expected error
	}{
		{
			name:   "nil input message",
			inMsg:  nil,
			expMsg: nil,
			expErr: status.Error(codes.InvalidArgument, "echo.EchoMessage is required"),
		}, {
			name:   "normal input message",
			inMsg:  &api.EchoMessage{Value: "hello!", Reverse: false},
			expMsg: &api.EchoMessage{Value: "hello!", Reverse: false},
			expErr: nil,
		}, {
			name:   "normal input message, reversed",
			inMsg:  &api.EchoMessage{Value: "hello!", Reverse: true},
			expMsg: &api.EchoMessage{Value: "!olleh", Reverse: true},
			expErr: nil,
		}, {
			name:   "error input; request: Aborted",
			inMsg:  &api.EchoMessage{Value: "Respond with Aborted"},
			expMsg: nil,
			expErr: status.Error(codes.Aborted, "Planned error for {Respond with Aborted}; Code {10}"),
		}, {
			name:   "error input; request: AlreadyExists",
			inMsg:  &api.EchoMessage{Value: "Respond with AlreadyExists"},
			expMsg: nil,
			expErr: status.Error(codes.AlreadyExists, "Planned error for {Respond with AlreadyExists}; Code {6}"),
		}, {
			name:   "error input; request: Canceled",
			inMsg:  &api.EchoMessage{Value: "Respond with Canceled"},
			expMsg: nil,
			expErr: status.Error(codes.Canceled, "Planned error for {Respond with Canceled}; Code {1}"),
		}, {
			name:   "error input; request: Cancelled",
			inMsg:  &api.EchoMessage{Value: "Respond with Cancelled"},
			expMsg: nil,
			expErr: status.Error(codes.Canceled, "Planned error for {Respond with Cancelled}; Code {1}"),
		}, {
			name:   "error input; request: DataLoss",
			inMsg:  &api.EchoMessage{Value: "Respond with DataLoss"},
			expMsg: nil,
			expErr: status.Error(codes.DataLoss, "Planned error for {Respond with DataLoss}; Code {15}"),
		}, {
			name:   "error input; request: DeadlineExceeded",
			inMsg:  &api.EchoMessage{Value: "Respond with DeadlineExceeded"},
			expMsg: nil,
			expErr: status.Error(codes.DeadlineExceeded, "Planned error for {Respond with DeadlineExceeded}; Code {4}"),
		}, {
			name:   "error input; request: FailedPrecondition",
			inMsg:  &api.EchoMessage{Value: "Respond with FailedPrecondition"},
			expMsg: nil,
			expErr: status.Error(codes.FailedPrecondition, "Planned error for {Respond with FailedPrecondition}; Code {9}"),
		}, {
			name:   "error input; request: Internal",
			inMsg:  &api.EchoMessage{Value: "Respond with Internal"},
			expMsg: nil,
			expErr: status.Error(codes.Internal, "Planned error for {Respond with Internal}; Code {13}"),
		}, {
			name:   "error input; request: InvalidArgument",
			inMsg:  &api.EchoMessage{Value: "Respond with InvalidArgument"},
			expMsg: nil,
			expErr: status.Error(codes.InvalidArgument, "Planned error for {Respond with InvalidArgument}; Code {3}"),
		}, {
			name:   "error input; request: NotFound",
			inMsg:  &api.EchoMessage{Value: "Respond with NotFound"},
			expMsg: nil,
			expErr: status.Error(codes.NotFound, "Planned error for {Respond with NotFound}; Code {5}"),
		}, {
			name:   "error input; request: OutOfRange",
			inMsg:  &api.EchoMessage{Value: "Respond with OutOfRange"},
			expMsg: nil,
			expErr: status.Error(codes.OutOfRange, "Planned error for {Respond with OutOfRange}; Code {11}"),
		}, {
			name:   "error input; request: PermissionDenied",
			inMsg:  &api.EchoMessage{Value: "Respond with PermissionDenied"},
			expMsg: nil,
			expErr: status.Error(codes.PermissionDenied, "Planned error for {Respond with PermissionDenied}; Code {7}"),
		}, {
			name:   "error input; request: ResourceExhausted",
			inMsg:  &api.EchoMessage{Value: "Respond with ResourceExhausted"},
			expMsg: nil,
			expErr: status.Error(codes.ResourceExhausted, "Planned error for {Respond with ResourceExhausted}; Code {8}"),
		}, {
			name:   "error input; request: Unauthenticated",
			inMsg:  &api.EchoMessage{Value: "Respond with Unauthenticated"},
			expMsg: nil,
			expErr: status.Error(codes.Unauthenticated, "Planned error for {Respond with Unauthenticated}; Code {16}"),
		}, {
			name:   "error input; request: Unavailable",
			inMsg:  &api.EchoMessage{Value: "Respond with Unavailable"},
			expMsg: nil,
			expErr: status.Error(codes.Unavailable, "Planned error for {Respond with Unavailable}; Code {14}"),
		}, {
			name:   "error input; request: Unimplemented",
			inMsg:  &api.EchoMessage{Value: "Respond with Unimplemented"},
			expMsg: nil,
			expErr: status.Error(codes.Unimplemented, "Planned error for {Respond with Unimplemented}; Code {12}"),
		}, {
			name:   "error input; request: Unknown",
			inMsg:  &api.EchoMessage{Value: "Respond with Unknown"},
			expMsg: nil,
			expErr: status.Error(codes.Unknown, "Planned error for {Respond with Unknown}; Code {2}"),
		},
	}

	for i, scenario := range scenarios {
		testName := fmt.Sprintf("[%v] %v", i, scenario.name)
		t.Run(testName, func(t *testing.T) {
			ctx := context.Background()
			svc := EchoServiceServer{}

			actMsg, actErr := svc.Echo(ctx, scenario.inMsg)
			assertEqual(t, scenario.expMsg, actMsg)
			assertEqual(t, scenario.expErr, actErr)
		})
	}
}

func assertEqual(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tact: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
