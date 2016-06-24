package main

import (
	"github.com/DanielRenne/mangosClient/rawReq"
	"log"
	"time"
)

const url = "tcp://127.0.0.1:600"

func main() {

	var c rawReq.Client

	err := c.Connect(url)
	if err != nil {
		log.Println("Error connecting to Reply server:  " + err.Error())
	}

	for {
		time.Sleep(3 * time.Second)
		c.Request([]byte("Send me a Reply"), handleReply)
	}
}

func handleReply(msg []byte) {
	log.Println(string(msg))
}
