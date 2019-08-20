package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"github.com/friendsofgo/kafka-example/pkg/kafka"
)

const (
	topic = "fogo-chat"
)

func main() {


	brokers := os.Getenv("KAFKA_BROKERS")

	publisher := kafka.NewPublisher(strings.Split(brokers, ","), topic)

	r := gin.Default()
	r.POST("/publish", func(c *gin.Context){
		if err := publisher.Publish(context.Background(), "hello apache kafka!"); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusOK, gin.H{"message": "message published"})
	})

	_ = r.Run()
}
