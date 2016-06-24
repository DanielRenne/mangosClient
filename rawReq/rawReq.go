//Package rawReq supports the implementation of a raw Request client.
package rawReq

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/req"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
)

type Client struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func([]byte)

//Connect to a Reply Server.
func (self *Client) Connect(url string) error {

	self.url = url

	var err error

	if self.sock, err = req.NewSocket(); err != nil {
		return err
	}
	self.sock.AddTransport(ipc.NewTransport())
	self.sock.AddTransport(tcp.NewTransport())
	if err = self.sock.Dial(url); err != nil {
		return err
	}

	return nil
}

//Send a request to the reply server.
func (self *Client) Request(payload []byte, handler ResponseHandler) error {

	body := []byte(payload)
	m := mangos.NewMessage(len(body) + 2)
	m.Body = body

	err := self.sock.SendMsg(m)
	go respondToRaw(self.sock, handler)
	return err
}

//handle the reply response.
func respondToRaw(sock mangos.Socket, handler ResponseHandler) {
	m, err := sock.RecvMsg()

	if err != nil {
		return
	}

	handler(m.Body)
}
