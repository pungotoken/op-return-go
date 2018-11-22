package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// GAMEParams network params
var GAMEParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000009),
	PubKeyHashAddrID: 0x26,
	ScriptHashAddrID: 0x5,
	PrivateKeyID:     0xA6,
}
