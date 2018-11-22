package conf

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/wire"
)

// KomodoParams for KMD and Assetchains
var KomodoParams = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000005),
	PubKeyHashAddrID: 0x3C,
	ScriptHashAddrID: 0x55,
	PrivateKeyID:     0xBC,
}
