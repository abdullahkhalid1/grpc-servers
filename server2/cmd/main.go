package main

import (
	api2 "github.com/abdullahkhalid1/grpc-servers/server1/api"
	"github.com/abdullahkhalid1/grpc-servers/server2/api"
	"github.com/abdullahkhalid1/grpc-servers/server2/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const PORT = "localhost:8000"

func main() {
	err := models.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	conn, err := grpc.Dial("localhost:5005", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	carClient := api2.NewCarServiceClient(conn)
	api.Routes(r, carClient)
	// By default it serves on :8080
	r.Run(PORT)
}
