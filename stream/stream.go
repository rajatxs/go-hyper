package stream

import "github.com/gorilla/websocket"

type Stream struct {
	id string
	ws *websocket.Conn
}

func New(id string) {

}
