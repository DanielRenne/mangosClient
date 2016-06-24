//Package respondent supports the implementation of a respondent client.
package respondent

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/respondent"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

type Client struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func(*Client, []byte)

//Connect to a Survey Server.
func (self *Client) Connect(url string, handler ResponseHandler) error {

	self.url = url

	var err error

	if self.sock, err = respondent.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Dial(url); err != nil {
		return err
	}

	go surveyMessageRoutine(self, handler)

	return nil
}

//Repsond to the Survey.
func (self *Client) Respond(payload []byte) error {
	return self.sock.Send(payload)
}

//Listens for survey messages and invokes the handler.
func surveyMessageRoutine(self *Client, handler ResponseHandler) {
	var msg []byte
	var err error
	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}

		handler(self, []byte(msg))
	}
}
