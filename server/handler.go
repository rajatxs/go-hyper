package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/rajatxs/go-hyper/log"
)

// Handles websocket handshake request
func handleWebsocketRequest(wsUpgrader *websocket.Upgrader) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			conn       *websocket.Conn
			err        error
			addr, sign string
			authOk     bool
		)

		addr, sign, authOk = r.BasicAuth()

		if !authOk {
			fmt.Fprintf(w, "invalid basic auth format")
			return
		}

		log.Debug("sign", sign)

		if conn, err = wsUpgrader.Upgrade(w, r, nil); err != nil {
			log.ErrorF("couldn't upgrade connection %s\n", err.Error())
			return
		} else {
			log.Infof("peer connected %s\n", addr)
		}

		defer conn.Close()

		for {
			var (
				msgtype int
				msg     []byte
			)

			if msgtype, msg, err = conn.ReadMessage(); err != nil {
				log.ErrorF("couldn't read message %s\n", err.Error())
				break
			}

			log.Infof("new message %d %s\n", msgtype, msg)
		}
	})
}

// Handles simple ping request
func handlePingRequest() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Pong!"))
	})
}
