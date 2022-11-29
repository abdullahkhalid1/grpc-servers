package main

import (
	"flag"
	"fmt"
	"github.com/abdullahkhalid1/grpc-servers/server1/api"
	"github.com/abdullahkhalid1/grpc-servers/server1/models"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port = flag.Int("port", 5005, "The server port")
)

func main() {
	err := models.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterGreetServer(s)
	api.RegisterCarServer(s)
	//models.RegisterDatabaseServiceServer(s, &server{})

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	//
	fmt.Println("Server Started")
	done := make(chan os.Signal, 1)
	fmt.Println("Signal Registered")
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal Done")

	<-done
	s.GracefulStop()
	fmt.Println("Server Stopped")
}
