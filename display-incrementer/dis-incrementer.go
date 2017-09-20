package main

import "fmt"
import "os"

import "github.com/go-redis/redis"
import "github.com/joho/godotenv"

//import "github.com/stianeikeland/go-rpio"

func main() {
	godotenv.Load()

	redisChannelName := os.Getenv("UNSPEAK_CHANNEL_SPOKEN")
	redisAddress := os.Getenv("UNSPEAK_REDIS_ADDRESS")

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pubsub := client.Subscribe(redisChannelName)
	defer pubsub.Close()

	done := make(chan bool)
	msg := make(chan string)

	go pinWorker(msg)
	go redisWorker(pubsub, msg, done)

	//Wait until we are Done...
	<-done

	fmt.Println("Exiting")
}

func redisWorker(pubsub *redis.PubSub, msg chan<- string, done chan<- bool) {
	for {
		message, err := pubsub.ReceiveMessage()
		if err != nil {
			break
		}
		msg <- message.Payload
	}
	close(msg)
	done <- true
}

func pinWorker(msg <-chan string) {

	// err := rpio.Open(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// defer rpio.Close();

	//pin := rpio.Pin(10);
	for {
		message, more := <-msg
		if !more {
			break
		}
		//At this point we should increment the counter
		//pin.Toggle();
		fmt.Println(message)
	}
}
