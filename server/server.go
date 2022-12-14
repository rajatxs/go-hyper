package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rajatxs/go-hyper/log"
)

type Server struct {
	instance *http.Server
	router   *mux.Router
	hostname string
	port     int
}

// Returns new instance of Server
func New(hostname string, port int, router *mux.Router) *Server {
	return &Server{
		hostname: hostname,
		port:     port,
		router:   router,
	}
}

// Returns server address
func (s *Server) Address() string {
	return fmt.Sprintf("%s:%d", s.hostname, s.port)
}

// Set of operation need to be execute before running server
func (s *Server) presetup() error {
	wsUpgrader := &websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		EnableCompression: false,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	s.router.HandleFunc("/ws", handleWebsocketRequest(wsUpgrader))
	s.router.HandleFunc("/ping", handlePingRequest())
	return nil
}

// Starts server instance
func (s *Server) Start() (err error) {
	addr := s.Address()

	if err = s.presetup(); err != nil {
		log.ErrorF("couldn't setup server %s", err.Error())
		return err
	}

	s.instance = &http.Server{
		Addr:    addr,
		Handler: s.router}

	log.Infof("starting server at %s", addr)
	return s.instance.ListenAndServe()
}

// Shutdowns server instance
func (s *Server) Stop() (err error) {
	err = s.instance.Close()

	if err == http.ErrServerClosed {
		return nil
	} else {
		return err
	}
}
