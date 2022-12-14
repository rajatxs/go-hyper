package peer

import (
	"crypto/sha256"
	"encoding/base64"
)

// Returns computed address from given auth and exchange key
func DeriveAddress(authKey []byte, exchKey []byte) []byte {
	h := sha256.New224()
	h.Write(append(authKey, exchKey...))
	return h.Sum(nil)
}

// Decodes url safe base64 public address
func DecodeAddress(rawAddr string) ([]byte, bool) {
	if addr, err := base64.RawURLEncoding.DecodeString(rawAddr); err != nil {
		return nil, false
	} else {
		return addr, len(addr) == AddressByteLength
	}
}

// Decodes url safe base64 authentication key
func DecodeAuthKey(rawAuthKey string) ([]byte, bool) {
	if authKey, err := base64.RawURLEncoding.DecodeString(rawAuthKey); err != nil {
		return nil, false
	} else {
		return authKey, len(authKey) == AuthKeyByteLength
	}
}

// Decodes url safe base64 exchange key
func DecodeExchangeKey(rawExchKey string) ([]byte, bool) {
	if exchKey, err := base64.RawURLEncoding.DecodeString(rawExchKey); err != nil {
		return nil, false
	} else {
		return exchKey, len(exchKey) == ExchangeKeyByteLength
	}
}
