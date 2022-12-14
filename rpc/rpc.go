package rpc

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"github.com/rajatxs/go-hyper/db"
	peer_rpc "github.com/rajatxs/go-hyper/rpc/peer"
)

type RPCHandler struct {
	server *rpc.Server
	codec  *json.Codec
	db     *db.Database
}

func New(db *db.Database) *RPCHandler {
	s := rpc.NewServer()
	c := json.NewCodec()

	s.RegisterCodec(c, "application/json")

	return &RPCHandler{
		server: s,
		codec:  c,
		db:     db,
	}
}

// Returns RPC handler
func (r *RPCHandler) Handler() *rpc.Server {
	return r.server
}

func (r *RPCHandler) RegisterMethods() (err error) {
	if err = r.server.RegisterService(peer_rpc.New(r.db), "peer"); err != nil {
		return err
	}

	return nil
}
