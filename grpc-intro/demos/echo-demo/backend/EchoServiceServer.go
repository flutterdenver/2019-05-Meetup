package main

import (
	"backend/api"
	"backend/stringx"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EchoServiceServer struct{}

func (s *EchoServiceServer) Echo(_ context.Context, msg *api.EchoMessage) (*api.EchoMessage, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "echo.EchoMessage is required")
	}

	plannedErrorFor := func(c codes.Code) error {
		return status.Errorf(c, "Planned error for {%v}; Code {%d}", msg.Value, c)
	}

	switch msg.Value {
	case "Respond with Aborted":
		return nil, plannedErrorFor(codes.Aborted)

	case "Respond with AlreadyExists":
		return nil, plannedErrorFor(codes.AlreadyExists)

	case "Respond with Canceled", "Respond with Cancelled": // Note the support both spellings;
		return nil, plannedErrorFor(codes.Canceled)

	case "Respond with DataLoss":
		return nil, plannedErrorFor(codes.DataLoss)

	case "Respond with DeadlineExceeded":
		return nil, plannedErrorFor(codes.DeadlineExceeded)

	case "Respond with FailedPrecondition":
		return nil, plannedErrorFor(codes.FailedPrecondition)

	case "Respond with Internal":
		return nil, plannedErrorFor(codes.Internal)

	case "Respond with InvalidArgument":
		return nil, plannedErrorFor(codes.InvalidArgument)

	case "Respond with NotFound":
		return nil, plannedErrorFor(codes.NotFound)

	case "Respond with OutOfRange":
		return nil, plannedErrorFor(codes.OutOfRange)

	case "Respond with PermissionDenied":
		return nil, plannedErrorFor(codes.PermissionDenied)

	case "Respond with ResourceExhausted":
		return nil, plannedErrorFor(codes.ResourceExhausted)

	case "Respond with Unauthenticated":
		return nil, plannedErrorFor(codes.Unauthenticated)

	case "Respond with Unavailable":
		return nil, plannedErrorFor(codes.Unavailable)

	case "Respond with Unimplemented":
		return nil, plannedErrorFor(codes.Unimplemented)

	case "Respond with Unknown":
		return nil, plannedErrorFor(codes.Unknown)
	}

	if msg.Reverse {
		msg.Value = stringx.Reverse(msg.Value)
	}

	return msg, nil
}
