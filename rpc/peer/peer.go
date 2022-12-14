package peer_rpc

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/rajatxs/go-hyper/db"
	"github.com/rajatxs/go-hyper/peer"
)

type peerRPCHandler struct {
	db *db.Database
}

// Returns new instance of peerRPCHandler
func New(db *db.Database) *peerRPCHandler {
	return &peerRPCHandler{db: db}
}

type DeriveAddressArgs struct {
	AuthKey     string `json:"authKey"`
	ExchangeKey string `json:"exchangeKey"`
}

func (prh *peerRPCHandler) DeriveAddress(r *http.Request, args *DeriveAddressArgs, reply *string) error {
	var (
		authKey []byte
		exchKey []byte
		addr    []byte
		pkb     *peer.KeyBundle
		valid   bool
	)

	if authKey, valid = peer.DecodeAuthKey(args.AuthKey); !valid {
		return errors.New("invalid auth key")
	}

	if exchKey, valid = peer.DecodeExchangeKey(args.ExchangeKey); !valid {
		return errors.New("invalid exchange key")
	}

	addr = peer.DeriveAddress(authKey, exchKey)
	pkb = peer.NewKeyBundle(addr, authKey, exchKey)

	*reply = pkb.EncodedAddress()
	return nil
}

type GetKeyBundleArgs struct {
	Address string `json:"address"`
}

type GetKeyBundleReply struct {
	AuthKey     string `json:"authKey"`
	ExchangeKey string `json:"exchangeKey"`
}

func (prh *peerRPCHandler) GetKeyBundle(r *http.Request, args *GetKeyBundleArgs, reply *GetKeyBundleReply) (err error) {
	var rkb *peer.RawKeyBundle

	rkb, err = prh.db.GetPeerRawKeyBundle(args.Address)

	if err == sql.ErrNoRows {
		return errors.New("key bundle not found")
	} else if err != nil {
		return errors.New("couldn't get key bundle")
	}

	*reply = GetKeyBundleReply{
		AuthKey:     rkb.AuthKey,
		ExchangeKey: rkb.ExchangeKey,
	}
	return nil
}

type SaveKeyBundleArgs struct {
	AuthKey     string `json:"authKey"`
	ExchangeKey string `json:"exchangeKey"`
}

type SaveKeyBundleReply struct {
	Address string `json:"address"`
}

func (prh *peerRPCHandler) SaveKeyBundle(r *http.Request, args *SaveKeyBundleArgs, reply *SaveKeyBundleReply) (err error) {
	return nil
}
