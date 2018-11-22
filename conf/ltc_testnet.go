package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// LTCTestnetParams network params
var LTCTestnetParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000006),
	PubKeyHashAddrID: 0x6f,
	ScriptHashAddrID: 0x3a,
	PrivateKeyID:     0xef,
}
