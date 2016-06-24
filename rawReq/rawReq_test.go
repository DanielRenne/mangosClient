package rawReq

import (
	"github.com/DanielRenne/mangosServer/raw"
	"github.com/go-mangos/mangos"
	"testing"
	"time"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T

func TestRawReq(t *testing.T) {
	tGlobal = t
	var s raw.Server
	err := s.Listen(url, 2, handleRawRequest)

	if err != nil {
		t.Errorf("Error starting listener at rawReq_test.TestRawReq:  %v", err.Error())
		return
	}

	var c Client
	err = c.Connect(url)

	if err != nil {
		t.Errorf("Error connecting client at rawReq_test.TestRawReq:  %v", err.Error())
		return
	}

	c.Request([]byte("TestingSurvey"), handleReply)

	time.Sleep(1 * time.Second)

}

func handleRawRequest(s *raw.Server, m *mangos.Message) {

	if string(m.Body) == "TestRequest" {
		tGlobal.Errorf("Failed to match request message.")
		return
	}
	m.Body = []byte("TestReply")
	s.Reply(m)
}

func handleReply(msg []byte) {
	if string(msg) != "TestReply" {
		tGlobal.Errorf("Failed to match reply message.")
		return
	}
}
