package conf

import (
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/miron/src/lib/btcsuite/btcd/wire"
)

// BTCTestNet3Params for bitcoin testnet v3
var BTCTestNet3Params = chaincfg.Params{
	Net:              wire.BitcoinNet(0x00000001),
	PubKeyHashAddrID: 0x6f, // starts with m or n
	ScriptHashAddrID: 0xc4, // starts with 2
	PrivateKeyID:     0xef, // starts with 9 (uncompressed) or c (compressed)
}
