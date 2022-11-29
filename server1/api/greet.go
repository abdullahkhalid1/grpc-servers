package api

import (
	"context"
	"google.golang.org/grpc"
	"strings"
)

type GreetServer struct {
	UnimplementedGreetServiceServer
}

func RegisterGreetServer(s *grpc.Server) {
	RegisterGreetServiceServer(s, &GreetServer{})
}

func (s *GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetResponse, error) {
	msg := "Hello World"
	reqMsg := strings.ToLower(request.GetMessage())
	if reqMsg == "hi" {
		msg = "Hi from server"
	} else if reqMsg == "hello" {
		msg = "Hello from server"
	} else if reqMsg == "salam" {
		msg = "Salam from server"
	}
	return &GreetResponse{Message: msg}, nil
}
