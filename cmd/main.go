package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
	"url-shortener/internal/database"
	"url-shortener/internal/models"
	"url-shortener/internal/server"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func initKafka() *kafka.Producer {
    producer, err := kafka.NewProducer(&kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
    })
    if err != nil {
        log.Fatal("failed to create kafka producer:", err)
    }

    return producer
}


func startgRPC() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	srv := server.NewServer(database.DB, database.RedisClient, initKafka(), "click-events")

	server.RegisterURLServiceServer(grpcServer, srv)

	log.Println("gRPC server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	r := gin.Default()

	// Initialize DB
	database.ConnectDB()
	database.NewRedis()

	go startgRPC()
	time.Sleep(1 * time.Second)

	// gRPC connection
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal("failed to connect to gRPC server:", err)
	}
	defer conn.Close()

	// Create gRPC client
	client := server.NewURLServiceClient(conn)

	// Route
	r.POST("/shorten", func(c *gin.Context) {
		var reqBody struct {
			URL string `json:"url"`
		}

		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(400, gin.H{"error": "invalid request"})
			return
		}

		// Create context (IMPORTANT)
		ctx := context.Background()

		// Call gRPC method
		resp, err := client.Shorten(ctx, &server.ShortenRequest{
			LongUrl: reqBody.URL,
		})
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"short_url": "localhost:3000/" + resp.ShortCode,
		})
	})

	r.GET("/:shortcode", func(c *gin.Context) {
		shortcode := c.Param("shortcode")

		var url models.URL
		result := database.DB.Where("short_code = ?", shortcode).First(&url)

		if result.Error != nil {
			fmt.Println(err)
		} else {
			var clicks models.Clicks
			
			ip := c.ClientIP()
			userAgent := c.Request.UserAgent()

			clicks = models.Clicks {
				ShortCode: url.ShortCode, 
				IP: ip,
				UserAgent: userAgent,
			}

			forward := database.DB.Create(&clicks).Error
			if forward != nil {
				fmt.Println("Couldn't track the hit request")
			} else {
				fmt.Println("Successfully Created a Click Record in Database!")
			}

			c.Redirect(http.StatusFound, url.LongURL)
		}
	})



	fmt.Println("Server is running on localhost:3000")

	r.Run(":3000")
}
