package main

import (
	"backend/api"

	"google.golang.org/grpc"
)

func main() {

	var s *grpc.Server // TODO needs to be created

	api.RegisterEchoServiceServer(s, &EchoServiceServer{
		//
	})
}
