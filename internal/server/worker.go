package server

import (
	context "context"
	"encoding/json"
	"log"
	"time"

	"url-shortener/internal/encoding"
	"url-shortener/internal/models"

	"github.com/bwmarrin/snowflake"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka" // Using Confluent Kafka
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	db            *gorm.DB
	redis         *redis.Client
	kafkaProducer *kafka.Producer
	topic         string

	UnimplementedURLServiceServer
}

func NewServer(
	db *gorm.DB,
	redisClient *redis.Client,
	producer *kafka.Producer,
	topic string,
) *Server {

	if db == nil {
		panic("db is nil")
	}
	if redisClient == nil {
		panic("redis is nil")
	}
	if producer == nil {
		log.Println("Kafka disabled")
	}

	return &Server{
		db:            db,
		redis:         redisClient,
		kafkaProducer: producer,
		topic:         topic,
	}
}

func generateSnowflake() int64 {
	// Create a new Node with a Node number of 1
	node, _ := snowflake.NewNode(1)

	// Generate a snowflake ID.
	id := node.Generate().Int64()

	return id
}

func (s *Server) Shorten(ctx context.Context, req *ShortenRequest) (*ShortenResponse, error) {
	id := generateSnowflake()

	code := encoding.EncodeBase62(id)

	url := models.URL{
		ShortCode: code,
		LongURL:   req.LongUrl,
	}

	if err := s.db.Create(&url).Error; err != nil {
		return nil, err
	}

	// Set Redis Cache for future lookup
	s.redis.Set(ctx, code, req.LongUrl, time.Hour*24)

	return &ShortenResponse{
		ShortCode: code,
	}, nil
}

func (s *Server) Resolve(ctx context.Context, req *ResolveRequest) (*ResolveResponse, error) {
	// Look up in redis Cache first for the ShortCode
	val, err := s.redis.Get(ctx, req.ShortCode).Result()
	if err == nil {
		go s.publishClick(req.ShortCode) // look up in Apache Kafka

		return &ResolveResponse{
			LongUrl: val,
		}, nil
	}
	// if Redis cache doesn't have the key/val, lookup in DB and populate in redis for future lookup's
	var url models.URL
	if err := s.db.Where("short_code = ?", req.ShortCode).First(&url).Error; err != nil {
		return nil, err
	}

	// final step, ie., populating in redis
	s.redis.Set(ctx, req.ShortCode, url.LongURL, time.Hour*24)

	go s.publishClick(req.ShortCode)

	return &ResolveResponse{
		LongUrl: url.LongURL,
	}, nil
}

// now comes the fun part Kafka or NATS .. for now, let's go with Kafka event driven
func (s *Server) publishClick(code string) {
	event := map[string]string{
		"short_code": code,
	}

	data, _ := json.Marshal(event)

	s.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &s.topic,
			Partition: kafka.PartitionAny,
		},
		Value: data,
	}, nil)
}

func consumeClicks(db *gorm.DB, reader *kafka.Consumer) {
	for {
		msg, err := reader.ReadMessage(-1) // -1 = block indefinitely
		if err != nil {
			continue
		}

		var event models.Clicks
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			continue
		}

		if err := db.Create(&event).Error; err != nil {
			continue
		}
	}
}
