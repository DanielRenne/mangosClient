//Package sub supports the implementation of a subscription client.
package sub

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"strings"
)

type Client struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func([]byte)

func (self *Client) Connect(url string) error {

	self.url = url

	var err error

	if self.sock, err = sub.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Dial(url); err != nil {
		return err
	}

	return nil
}

func (self *Client) Subscribe(topic string, handler ResponseHandler) error {

	// Empty byte array effectively subscribes to everything
	err := self.sock.SetOption(mangos.OptionSubscribe, []byte(topic))
	if err != nil {
		return err
	}

	go subscribeRoutine(self, topic, handler)

	return nil
}

func subscribeRoutine(self *Client, topic string, handler ResponseHandler) {
	var msg []byte
	var err error
	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}

		msgTopic := strings.Replace(string(msg), topic+"|", "", -1)
		handler([]byte(msgTopic))
	}
}
