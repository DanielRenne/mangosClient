package respondent

import (
	"github.com/DanielRenne/mangosServer/survey"
	"testing"
	"time"
)

const url = "tcp://127.0.0.1:600"

var tGlobal *testing.T

func TestRespondent(t *testing.T) {
	tGlobal = t
	var s survey.Server
	err := s.Listen(url, 3000, 2, handleSurveyResponse)

	if err != nil {
		t.Errorf("Error starting listener at sub_test.TestSingleSubscription:  %v", err.Error())
		return
	}

	var c Client
	err = c.Connect(url, handleSurveyMessage)

	if err != nil {
		t.Errorf("Error connecting client at sub_test.TestSingleSubscription:  %v", err.Error())
		return
	}

	s.Send([]byte("TestingSurvey"))

	time.Sleep(1 * time.Second)

}

func handleSurveyMessage(c *Client, msg []byte) {
	if string(msg) != "TestingSurvey" {
		tGlobal.Errorf("Error connecting client at respondent_test.handleSurveyMessage")
		return
	}

	c.Respond([]byte("SurveyResponse"))
}

func handleSurveyResponse(msg []byte) {

	if string(msg) != "SurveyResponse" {
		tGlobal.Errorf("Failed to match the survey response message at respondent_test.handleSurveyResponse")
		return
	}

}
