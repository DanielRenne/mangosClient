# mangosClient rawReq

Example Code to start a raw reply server, request a message, reply with a new message.  The server will have 2 workers. 

###Server Code
	package main

	import (
		"github.com/DanielRenne/mangosServer/raw"
		"github.com/go-mangos/mangos"
		"log"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	func main() {
		var s raw.Server
	
		err := s.Listen(url, 2, handleRawRequest)
		if err != nil {
			log.Printf("Error:  %v", err.Error)
		}
	
		//Code a forever loop to stop main from exiting.
		for {
	
		}
	
	}
	
	func handleRawRequest(s *raw.Server, m *mangos.Message) {
	
		log.Printf(string(m.Body))
		m.Body = []byte("Custom Response to the Request")
		s.Reply(m)
	}



###Client Code 

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
