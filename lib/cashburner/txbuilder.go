package cashburner

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcwallet/wallet/txauthor"
	"gitlab.com/instak_backend/op_return/lib/utxorq"
)

// Build p2pkh transaction (standart value transfer transaction)
func Build(network chaincfg.Params, key string, forwardAddress string, changeAddress string, amount int, fee int, utxos []*utxorq.UTXOValue, message string) (string, error) {
	inputs := make([]*wire.TxIn, 0)
	outputs := make([]*wire.TxOut, 0)

	// we expect private key in wif uncompressed format
	// so need to decode it first
	wif, err := btcutil.DecodeWIF(key)
	if err != nil {
		return "", fmt.Errorf("failed parsing wif private key")
	}
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeCompressed(), &network)
	// addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &network)
	sourceAddressScript, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &network)
	if err != nil {
		return "", fmt.Errorf("failed creating source address script")
	}

	// generate inputs for utxos of forwarded wallet's address
	inputs = make([]*wire.TxIn, 0, len(utxos))
	for _, v := range utxos {
		txHash, err := chainhash.NewHashFromStr(v.Hash)
		if err != nil {
			return "", err
		}
		previousOutPoint := wire.OutPoint{Hash: *txHash, Index: uint32(v.Pos)}
		inputs = append(inputs, wire.NewTxIn(&previousOutPoint, nil, nil))
	}

	// create OP_RETURN script
	forwardAddressScript, err := txscript.NullDataScript([]byte(message))
	if err != nil {
		return "", err
	}
	forwardTxOut := wire.NewTxOut(int64(amount), forwardAddressScript)
	outputs = append(outputs, forwardTxOut)

	// process change
	valueCount := uint64(0)
	valueChange := uint64(0)

	for _, v := range utxos {
		valueCount += uint64(v.Value)
	}

	fmt.Println("transaction value: ", amount)
	fmt.Println("transaction fee: ", fee)

	valueChange = valueCount - uint64(amount+fee)

	if valueChange > 0 {

		fmt.Println("utxo change amount: ", valueChange)
		fmt.Println("utxo change address: ", changeAddress)

		// process address to which we will send change
		changeAddressData, err := btcutil.DecodeAddress(changeAddress, &network)
		if err != nil {
			return "", err
		}
		changeAddressScript, err := txscript.PayToAddrScript(changeAddressData)
		if err != nil {
			return "", err
		}
		changeTxOut := wire.NewTxOut(int64(valueChange), changeAddressScript)
		outputs = append(outputs, changeTxOut)
	}

	// unsigned transaction structure
	unsignedTransaction := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}
	tx := &txauthor.AuthoredTx{
		Tx: unsignedTransaction,
	}

	// sign new transaction
	err = signTransaction(tx, wif, sourceAddressScript)
	if err != nil {
		return "", err
	}

	// signed transaction data
	var signedTx bytes.Buffer
	tx.Tx.Serialize(&signedTx)
	signedTxEncoded := hex.EncodeToString(signedTx.Bytes())

	return signedTxEncoded, nil
}

// BuildDouble OP_RETURN transaction with additional forwarding of funds
func BuildDouble(network chaincfg.Params, key string, forwardAddress string, amount int, fee int, utxos []*utxorq.UTXOValue) (string, error) {
	inputs := make([]*wire.TxIn, 0)
	outputs := make([]*wire.TxOut, 0)

	// we expect private key in wif uncompressed format
	// so need to decode it first
	wif, err := btcutil.DecodeWIF(key)
	if err != nil {
		return "", fmt.Errorf("failed parsing wif private key")
	}
	addresspubkey, _ := btcutil.NewAddressPubKey(wif.PrivKey.PubKey().SerializeUncompressed(), &network)
	sourceAddressScript, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), &network)
	changeAddress := sourceAddressScript.String()
	println("change address: ", changeAddress)
	if err != nil {
		return "", fmt.Errorf("failed creating source address script")
	}

	// generate inputs for utxos of forwarded wallet's address
	inputs = make([]*wire.TxIn, 0, len(utxos))
	for _, v := range utxos {
		txHash, err := chainhash.NewHashFromStr(v.Hash)
		if err != nil {
			return "", err
		}
		previousOutPoint := wire.OutPoint{Hash: *txHash, Index: uint32(v.Pos)}
		inputs = append(inputs, wire.NewTxIn(&previousOutPoint, nil, nil))
	}

	// process address to which we will forward payment
	address, err := btcutil.DecodeAddress(forwardAddress, &network)
	if err != nil {
		return "", err
	}
	forwardAddressScript, err := txscript.PayToAddrScript(address)
	if err != nil {
		return "", err
	}
	forwardTxOut := wire.NewTxOut(int64(amount), forwardAddressScript)
	outputs = append(outputs, forwardTxOut)

	// process change
	// if changeAddress &&
	valueCount := uint64(0)
	valueChange := uint64(0)
	_ = valueChange

	for _, v := range utxos {
		valueCount += uint64(v.Value)
	}

	valueChange = valueCount - uint64(amount+fee)
	// fmt.Println("transaction change amount: ", valueChange)

	if valueChange > 0 {
		// process address to which we will send change
		changeAddressData, err := btcutil.DecodeAddress(changeAddress, &network)
		if err != nil {
			return "", err
		}
		changeAddressScript, err := txscript.PayToAddrScript(changeAddressData)
		if err != nil {
			return "", err
		}
		changeTxOut := wire.NewTxOut(int64(valueChange), changeAddressScript)
		outputs = append(outputs, changeTxOut)
	}

	// unsigned transaction structure
	unsignedTransaction := &wire.MsgTx{
		Version:  wire.TxVersion,
		TxIn:     inputs,
		TxOut:    outputs,
		LockTime: 0,
	}
	tx := &txauthor.AuthoredTx{
		Tx: unsignedTransaction,
	}

	// sign new transaction
	err = signTransaction(tx, wif, sourceAddressScript)
	if err != nil {
		return "", err
	}

	// signed transaction data
	var signedTx bytes.Buffer
	tx.Tx.Serialize(&signedTx)
	signedTxEncoded := hex.EncodeToString(signedTx.Bytes())

	return signedTxEncoded, nil
}

func signTransaction(tx *txauthor.AuthoredTx, wif *btcutil.WIF, sourceAddressScript btcutil.Address) error {
	for k := range tx.Tx.TxIn {
		pkScript, _ := txscript.PayToAddrScript(sourceAddressScript)

		sig, err := txscript.RawTxInSignature(tx.Tx, k, pkScript, txscript.SigHashAll, wif.PrivKey)
		if err != nil {
			return err
		}

		pk := (*btcec.PublicKey)(&wif.PrivKey.PublicKey)
		pkData := pk.SerializeCompressed()
		// pkData := pk.SerializeUncompressed()
		script, err := txscript.NewScriptBuilder().AddData(sig).AddData(pkData).Script()
		if err != nil {
			return err
		}

		tx.Tx.TxIn[k].SignatureScript = script
	}
	return nil
}
