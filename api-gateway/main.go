package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pbA "github.com/khouini/microservices-hello-world/service-a/pb"
	pbB "github.com/khouini/microservices-hello-world/service-b/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayServer struct {
	serviceAClient pbA.ServiceAHelloClient
	serviceBClient pbB.ServiceBHelloClient
}

func main() {
	// Set up gRPC connections
	serviceAConn, err := grpc.Dial("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to Service A: %v", err)
	}
	defer serviceAConn.Close()
	serviceAClient := pbA.NewServiceAHelloClient(serviceAConn)

	serviceBConn, err := grpc.Dial("localhost:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect to Service B: %v", err)
	}
	defer serviceBConn.Close()
	serviceBClient := pbB.NewServiceBHelloClient(serviceBConn)

	// Create gateway server
	gateway := &GatewayServer{
		serviceAClient: serviceAClient,
		serviceBClient: serviceBClient,
	}

	// Setup Gin router
	r := gin.Default()

	// Routes
	r.GET("/service-a", gateway.HandleServiceA)
	r.GET("/service-b", gateway.HandleServiceB)

	// Start server
	log.Println("API Gateway listening on :8080")
	r.Run(":8080")
}

func (g *GatewayServer) HandleServiceA(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := g.serviceAClient.SayHello(ctx, &pbA.ServiceAHelloRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}

func (g *GatewayServer) HandleServiceB(c *gin.Context) {
	name := c.DefaultQuery("name", "World")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := g.serviceBClient.SayHello(ctx, &pbB.ServiceBHelloRequest{Name: name})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": resp.Message,
	})
}
