package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{conns: make(map[*websocket.Conn]bool)}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("incomming connection from client: ", ws.RemoteAddr())

	// Add connection to the connection map and set status to true
	// TODO: add Mutex here to keep map thread safe
	s.conns[ws] = true

	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)

		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("read error: ", err)
			// dont break connection if it failes
			continue
		}

		msg := buf[:n]

		fmt.Println(string(msg))
		ws.Write([]byte("thank you for the message"))
	}

}

func main() {
	server := NewServer()

	http.Handle("/ws", websocket.Handler(server.handleWS))

	http.ListenAndServe(":3000", nil)
}


