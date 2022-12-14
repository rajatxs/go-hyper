package peer

import (
	"bytes"
	"encoding/base64"
	"errors"
)

const (
	AddressByteLength     = 0x1c
	AuthKeyByteLength     = 0x41
	ExchangeKeyByteLength = 0x61
	KeyBundleByteLength   = AuthKeyByteLength + ExchangeKeyByteLength
)

type KeyBundle struct {
	Address     []byte `json:"address"`
	AuthKey     []byte `json:"authKey"`
	ExchangeKey []byte `json:"exchangeKey"`
}

// Returns new instance of KeyBundle
func NewKeyBundle(addr, authKey, exchKey []byte) *KeyBundle {
	return &KeyBundle{
		Address:     addr,
		AuthKey:     authKey,
		ExchangeKey: exchKey,
	}
}

// Returns encoded peer address
func (kb *KeyBundle) EncodedAddress() string {
	return base64.RawURLEncoding.EncodeToString(kb.Address)
}

// Returns encoded peer auth key
func (kb *KeyBundle) EncodedAuthKey() string {
	return base64.RawURLEncoding.EncodeToString(kb.AuthKey)
}

// Returns encoded peer exchange key
func (kb *KeyBundle) EncodedExchangeKey() string {
	return base64.RawURLEncoding.EncodeToString(kb.ExchangeKey)
}

// Returns validity of address
func (kb *KeyBundle) VerifyAddress() bool {
	return bytes.Equal(kb.Address, DeriveAddress(kb.AuthKey, kb.ExchangeKey))
}

// Performs sanity check on keybundle
func (kb *KeyBundle) SanityCheck() error {
	switch {
	default:
		return nil
	case len(kb.AuthKey) != AuthKeyByteLength:
		return errors.New("invalid auth key length")
	case len(kb.ExchangeKey) != ExchangeKeyByteLength:
		return errors.New("invalid exchange key length")
	case len(kb.Address) != AddressByteLength:
		return errors.New("invalid address length")
	case !kb.VerifyAddress():
		return errors.New("invalid key bundle")
	}
}

func (kb *KeyBundle) VerifySignature() {

}
