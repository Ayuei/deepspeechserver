package main

import (
	"fmt"
	"github.com/ayuei/deepspeechserver/server"
	"github.com/ayuei/deepspeechserver/speech"
	"github.com/ayuei/deepspeechserver/utils"
	"net/http"
)

var Model *speech.Speech

func serveStream(w http.ResponseWriter, r *http.Request){
	fmt.Println("Websocket endpoint hit")

	conn, err := server.Upgrader.Upgrade(w, r, nil)
	utils.Chkerror(err)

	client := &server.Client{
		Conn:        conn,
		ReadBuffer:  make(chan server.SpeechMessage),
		WriteBuffer: make(chan string),
	}

	go client.Read()
	go client.Write()
	Model.Start(client)
}

func setupRoutes(){
	http.HandleFunc("/stream", serveStream)
}

func setupModel(){
	Model = speech.New("deepspeech-0.7.1-models.pbmm", "deepspeech-0.7.1-models.scorer")
}

func main() {
	fmt.Println("Launching backend")
	setupModel()
	setupRoutes()
	utils.Chkerror(http.ListenAndServe(":8080", nil))
}
