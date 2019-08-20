package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	_ "github.com/joho/godotenv/autoload"

	kafkaexample "github.com/friendsofgo/kafka-example/pkg"
	"github.com/friendsofgo/kafka-example/pkg/kafka"
)

var user string

type Data struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {

	var (
		host    = os.Getenv("HOST")
		brokers = os.Getenv("KAFKA_BROKERS")
		topic   = os.Getenv("KAFKA_TOPIC")
	)

	fmt.Println("Choose your nickname:")
	_, _ = fmt.Scanf("%s\n", &user)
	if err := joinTheRoom(host); err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			msg, _ := reader.ReadString('\n')
			if err := publishMessage(host, msg); err != nil {
				log.Fatal(err)
			}
		}
	}()

	chMsg := make(chan kafkaexample.Message)
	chErr := make(chan error)
	consumer := kafka.NewConsumer(strings.Split(brokers, ","), topic)

	go func() {
		consumer.Read(context.Background(), chMsg, chErr)
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			goto end
		case m := <-chMsg:
			printMessage(m)
		case err := <-chErr:
			log.Println(err)
		}
	}
end:

	fmt.Println("\nyou have abandoned the room")

}

func printMessage(m kafkaexample.Message) {

	switch {
	case m.Username == user:
		return
	case m.Username == kafkaexample.SystemID:
		fmt.Println(m.Message)
	default:
		fmt.Printf("%s say: %s", m.Username, m.Message)
	}
}

func joinTheRoom(host string) error {
	d := Data{Username: user}
	endpoint := fmt.Sprintf("%s/join", host)
	return do(endpoint, d)
}

func publishMessage(host, message string) error {
	d := Data{Username: user, Message: message}
	endpoint := fmt.Sprintf("%s/publish", host)
	return do(endpoint, d)
}

func do(endpoint string, data Data) error {
	requestBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}

		return errors.New(string(body))
	}

	return nil
}
