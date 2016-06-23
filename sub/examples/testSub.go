package main

import (
	"github.com/DanielRenne/mangosClient/sub"
	"log"
)

const url = "tcp://127.0.0.1:600"

//Creates a new Pub Server and broadcasts a plain message
func main() {

	var c sub.Client
	err := c.Connect(url)

	if err != nil {
		log.Printf("Error connecting client:  %v", err.Error())
		return
	}

	c.Subscribe("", handleSubscription)

	for {

	}

}

func handleSubscription(msg []byte) {
	log.Printf(string(msg))
}
