# mangosClient sub

To subscribe to any publisher message:

###Server Code

	package main
	
	import (
		"github.com/DanielRenne/mangosServer/pub"
		"log"
		"time"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	//Creates a new Pub Server and broadcasts a plain message
	func main() {
		var s pub.Server
		err := s.Listen(url)
		if err != nil {
			log.Printf("Error:  %v", err.Error())
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go s.Publish([]byte("Publishing Message."))
		}
	
	}


###Client Code 

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



To subscribe to a specific Topic.

###Server Code

	package main
	
	import (
		"github.com/DanielRenne/mangosServer/pub"
		"log"
		"time"
	)
	
	const url = "tcp://127.0.0.1:600"
	
	//Creates a new Pub Server and broadcasts a plain message
	func main() {
		var s pub.Server
		err := s.Listen(url)
		if err != nil {
			log.Printf("Error:  %v", err.Error())
		}
	
		//Code a forever loop to stop main from exiting.
		for {
			time.Sleep(3 * time.Second)
			go s.PublishTopic("TestTopic", "Publishing Message.")
		}
	
	}


###Client Code
	
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
	
		c.Subscribe("TestTopic", handleSubscription)
	
		for {
	
		}
	
	}
	
	func handleSubscription(msg []byte) {
		log.Printf(string(msg))
	}


