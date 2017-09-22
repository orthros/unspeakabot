package main

import "fmt"
import "os"

import "github.com/go-redis/redis"
import "github.com/joho/godotenv"

import "github.com/stianeikeland/go-rpio"

func main() {
    godotenv.Load()

    redisChannelName := os.Getenv("UNSPEAK_CHANNEL_SPOKEN")
    redisAddress := os.Getenv("UNSPEAK_REDIS_ADDRESS")

    done := make(chan bool)
    msg := make(chan string)

    go pinWorker(msg)
    go redisWorker(redisAddress, redisChannelName, msg, done)

    //Wait until we are Done...
    <-done

    fmt.Println("Exiting")
}

func redisWorker(redisAddress string, redisChannelName string, msg chan<- string, done chan<- bool) {
    client := redis.NewClient(&redis.Options{
        Addr:     redisAddress,
        Password: "", // no password set
        DB:       0,  // use default DB
    })
    defer client.Close()
    fmt.Println("Worker created")

    pubsub := client.Subscribe(redisChannelName)
    defer pubsub.Close()
    fmt.Println("Subscribed")

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
    fmt.Println("In Pin worker")
    err := rpio.Open()

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer rpio.Close();
    fmt.Println("Pin created")
    //pin := rpio.Pin(10);
    for message := range msg {
    //At this point we should increment the counter
    //pin.Toggle();
        fmt.Println(message)
    }
}
