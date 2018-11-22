package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// DashTestnetParams network params
var DashTestnetParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000003),
	PubKeyHashAddrID: 0x8c,
	ScriptHashAddrID: 0x13,
	PrivateKeyID:     0xef,
}
