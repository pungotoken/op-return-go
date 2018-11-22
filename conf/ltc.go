package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// LTCParams network params
var LTCParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000007),
	PubKeyHashAddrID: 0x30,
	ScriptHashAddrID: 0x32,
	PrivateKeyID:     0xb0,
}
