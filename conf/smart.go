package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// SMARTMainNetParams for smart mainnet
var SMARTMainNetParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000008),
	PubKeyHashAddrID: 0x3F,
	ScriptHashAddrID: 0x12,
	PrivateKeyID:     0xBF,
}
