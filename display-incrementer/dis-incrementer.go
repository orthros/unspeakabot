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

	// err := rpio.Open(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// defer rpio.Close();

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	pubsub := client.Subscribe(redisChannelName)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage()
		if err != nil {
			panic(err)
		}
		//At this point we should increment the counter
		//pin = rpio.Pin(10);
		//pin.Toggle();
		fmt.Println(msg)
	}

	fmt.Println("Exiting")
}
