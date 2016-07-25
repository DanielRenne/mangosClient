//Package sub supports the implementation of a subscription client.
package sub

import (
	"github.com/go-mangos/mangos"
	"github.com/go-mangos/mangos/protocol/sub"
	"github.com/go-mangos/mangos/transport/ipc"
	"github.com/go-mangos/mangos/transport/tcp"
	"strings"
	"sync"
)

type topicHandlerMap struct {
	M        sync.RWMutex
	handlers map[string]ResponseHandler
}

type Client struct {
	url  string
	sock mangos.Socket
}

type ResponseHandler func([]byte)

var topicHandlers topicHandlerMap
var isReading bool

func init() {
	topicHandlers.handlers = make(map[string]ResponseHandler)
}

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

	if !isReading {
		go subscribeRoutine(self, handler)
	}

	topicHandlers.M.Lock()
	topicHandlers.handlers[topic] = handler
	topicHandlers.M.Unlock()

	return nil
}

func subscribeRoutine(self *Client, handler ResponseHandler) {
	isReading = true
	var msg []byte
	var err error
	for {
		if msg, err = self.sock.Recv(); err != nil {
			continue
		}

		msgString := string(msg)

		var topic string

		if strings.ContainsAny(msgString, "|") {
			topic = msgString[0:strings.Index(msgString, "|")]
		}

		msgTopic := strings.Replace(string(msg), topic+"|", "", -1)

		topicHandlers.M.RLock()
		handleFunc := topicHandlers.handlers[topic]

		if handleFunc != nil {
			handleFunc([]byte(msgTopic))
		}

		topicHandlers.M.RUnlock()

	}
}
