package server

import (
	"encoding/json"
	"fmt"
	"github.com/ayuei/deepspeechserver/utils"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type SpeechMessage struct{
	PCM []int16
}

type Client struct {
	*websocket.Conn
	ReadBuffer chan SpeechMessage
	WriteBuffer chan string
}

func (c *Client) Read() {
	for {
		_, p, err := c.ReadMessage()
		fmt.Println("Received Message")

		if utils.Chkerror(err) {
			return
		}

		var msg SpeechMessage
		err = json.Unmarshal(p, &msg)

		fmt.Println("Unmarshalled File")

		if utils.Chkerror(err) {
			return
		}

		c.ReadBuffer <- msg
		fmt.Println("Put message in buffer")
	}
}

func (c *Client) Write() {
	for {
		select{
		case buffer := <- c.WriteBuffer:
			err := c.WriteMessage(1, []byte(buffer))
			if utils.Chkerror(err) {
				return
			}
		}
	}
}

// Define an upgrader
// Upgrades a http connection to a websocket connection
var Upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// Check the origin of the connection
	CheckOrigin: func(r *http.Request) bool {return true},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}
