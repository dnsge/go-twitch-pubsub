package main

import (
	"fmt"
	"github.com/dnsge/go-twitch-pubsub"
	"net/http"
)

func main() {
	// Create pooled client
	client := twitchpubsub.NewPubSubPool("my-auth-token", http.Header{})

	// Listen to topic
	_, err := client.Listen("chat_moderator_actions.12826", func(data twitchpubsub.MessageData) {
		fmt.Printf("Moderator action: %s\n", data.Message)
	})

	if err != nil {
		panic(err)
	}

	client.OnStart = func() {
		fmt.Println("Connected!")
	}

	// Start and wait
	err = client.Start()
	if err != nil {
		panic(err)
	}

	wait := make(chan bool)
	<-wait
}
