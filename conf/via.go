package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// VIAParams network params
var VIAParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000010),
	PubKeyHashAddrID: 0x47,
	ScriptHashAddrID: 0x21,
	PrivateKeyID:     0xc7,
}
