package main

import (
	"github.com/DanielRenne/mangosClient/respondent"
	"log"
)

const url = "tcp://127.0.0.1:600"

func main() {

	var c respondent.Client

	c.Connect(url, handleSurveyMessage)

	for {

	}
}

func handleSurveyMessage(c *respondent.Client, msg []byte) {
	log.Println(string(msg))
	c.Respond([]byte("Respondent"))
}
