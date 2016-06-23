package sub

import (
	"github.com/DanielRenne/mangosServer/pub"
	"testing"
	"time"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T

func TestSingleSubscription(t *testing.T) {
	tGlobal = t
	var s pub.Server
	err := s.Listen(url)

	if err != nil {
		t.Errorf("Error starting listener at sub_test.TestSingleSubscription:  %v", err.Error())
		return
	}

	var c Client
	err = c.Connect(url)

	if err != nil {
		t.Errorf("Error connecting client at sub_test.TestSingleSubscription:  %v", err.Error())
		return
	}

	c.Subscribe("", handleSubscription)

	time.Sleep(1 * time.Second)

	s.Publish([]byte("TestingSubscription"))

	time.Sleep(1 * time.Second)

}

func handleSubscription(msg []byte) {
	if string(msg) != "TestingSubscription" {
		tGlobal.Errorf("Error matching subscription response at sub_test.handleSubscription")
	}
}
