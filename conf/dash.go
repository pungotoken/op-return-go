package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// DashParams network params
var DashParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000004),
	PubKeyHashAddrID: 0x4c,
	ScriptHashAddrID: 0x10,
	PrivateKeyID:     0xcc,
}
