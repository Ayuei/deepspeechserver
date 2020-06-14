# deepspeechserver

A websocket-based hosted deepspeech server

## Dependencies
- go-astideepspeech (follow build instructions [go-astideepspeech](https://github.com/asticode/go-astideepspeech))
- gorilla-websocket

## How to use

1. Start up the server
```
go run server.go
```

2. Set up a client to send messages (client.go is given as an example)
```
go run client.go
```

This is a work in progress, currently quite due to inefficently in marshalling.
