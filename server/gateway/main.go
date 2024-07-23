package main

import (
	"log"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/gin-gonic/gin"
	"github.com/ssonit/common"

	pb "github.com/ssonit/common/protos/order"
)

var (
	httpAddr         = common.EnvConfig("HTTP_ADDR", ":3000")
	orderServiceAddr = common.EnvConfig("ORDER_SERVICE_ADDR", "localhost:50051")
)

func main() {
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	log.Printf("New client with order service on %s", orderServiceAddr)

	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}

	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)

	r := gin.Default()

	h := NewHandler(c)
	h.RegisterRoutes(r)

	log.Printf("Starting server on %s", httpAddr)

	r.Run(httpAddr)
}
