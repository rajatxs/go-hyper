package peer

import "errors"

const (
	RawAddressByteLength = 0x26
	RawAuthKeyLength     = 0x57
	RawExchangeKeyLength = 0x82
	RawKeyBundleLength   = RawAuthKeyLength + RawExchangeKeyLength
)

type RawKeyBundle struct {
	Address     string `json:"address"`
	AuthKey     string `json:"authKey"`
	ExchangeKey string `json:"exchangeKey"`
}

// Returns parsed KeyBundle
func (rkb *RawKeyBundle) Parse() (*KeyBundle, error) {
	var valid bool

	kb := &KeyBundle{}

	if kb.Address, valid = DecodeAddress(rkb.Address); !valid {
		return nil, errors.New("couldn't parse address")
	}

	if kb.AuthKey, valid = DecodeAuthKey(rkb.AuthKey); !valid {
		return nil, errors.New("couldn't parse auth key")
	}

	if kb.ExchangeKey, valid = DecodeExchangeKey(rkb.ExchangeKey); !valid {
		return nil, errors.New("couldn't parse exchange key")
	}

	return kb, nil
}
