package main

import (
	"fmt"
	"github.com/ayuei/deepspeechserver/server"
	"github.com/ayuei/deepspeechserver/utils"
	"github.com/cryptix/wav"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"os"
)

const addr = "localhost:8080"

func main(){
	u := url.URL{Scheme: "ws", Host: addr, Path: "/stream"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	utils.Chkerror(err)

	audio := "audio/4507-16021-0012.wav"
	// Stat audio
	i, err := os.Stat(audio)
	if err != nil {
		log.Fatal(fmt.Errorf("stating %s failed: %w", audio, err))
	}

	// Open audio
	f, err := os.Open(audio)
	if err != nil {
		log.Fatal(fmt.Errorf("opening %s failed: %w", audio, err))
	}

	// Create reader
	r, err := wav.NewReader(f, i.Size())
	if err != nil {
		log.Fatal(fmt.Errorf("creating new reader failed: %w", err))
	}

	// Read
	var d []int16
	for {
		// Read sample
		s, err := r.ReadSample()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(fmt.Errorf("reading sample failed: %w", err))
		}

		// Append
		d = append(d, int16(s))
	}

	err = conn.WriteJSON(server.SpeechMessage{PCM: d})
	err = conn.WriteJSON(server.SpeechMessage{PCM: []int16{}})
	utils.Chkerror(err)

	_, p, _ := conn.ReadMessage()

	fmt.Println(string(p))

	conn.Close()
}
