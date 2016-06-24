# mangosClient Respondent

Example Code to start a survey server, send a survey, and receive a survey response.
###Survey Server Code

	package main
	
	import (
		"github.com/DanielRenne/mangosServer/survey"
		"log"
		"time"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	func main() {
		var s survey.Server
	
		err := s.Listen(url, 500, 2, handleSurveyResponse)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go s.Send([]byte("Sending Survey"))
		}
	
	}
	
	func handleSurveyResponse(msg []byte) {
		//Process Survey Results.
		log.Printf(string(msg))
	}

	
###Respondent Client Code

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
	