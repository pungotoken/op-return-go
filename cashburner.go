package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"

	"github.com/btcsuite/btcd/chaincfg"
	"gitlab.com/instak_backend/op_return/conf"
	"gitlab.com/instak_backend/op_return/lib/cashburner"
	"gitlab.com/instak_backend/op_return/lib/utxorq"
)

func main() {
	defer func() {
		println("\n\n---------------- done ----------------\n\n")
	}()
	println("\n")

	// configure blockchain networks
	networks := map[string]chaincfg.Params{
		"PGT": conf.KomodoParams,
	}
	for _, v := range networks {
		chaincfg.Register(&v)
	}

	// configure arguments
	argAmount := flag.Float64("val", 0, "transaction amount")
	argKey := flag.String("key", "", "wallet privkey")
	argDest := flag.String("dest", "", "destination address")
	argSource := flag.String("src", "", "source address")
	argMessage := flag.String("msg", "", "transaction data message")
	flag.Parse()

	if *argAmount <= float64(0) {
		println("invalid transaction amount value")
		return
	}
	if *argKey == "" {
		println("invalid privkey value")
		return
	}
	if *argSource == "" {
		println("invalid source address value")
		return
	}
	// if *argDest == "" {
	// 	println("invalid destination address value")
	// 	return
	// }
	if *argMessage == "" {
		println("invalid message data value")
		return
	}

	// data
	println("amount: ", *argAmount)
	println("message: ", *argMessage)
	println("source address: ", *argSource)
	println("source privkey: ", *argKey)
	println("\n")

	// fetch utxos
	source := "https://electrum.pungo.network/api/listunspent?ip=electrum1.komodo.build&port=10002&address=" + *argSource
	res, err := http.Get(source)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	var utxos = new(utxorq.UTXOResponse)
	err = json.Unmarshal(body, &utxos)
	if err != nil {
		panic(err.Error())
	}

	// check available utxos
	if len(utxos.Result) == 0 {
		println("no available utxos")
		return
	}

	println("utxos: ")
	for k, v := range utxos.Result {
		println(k, v.Hash, v.Value)
	}
	println("\n")

	// amount to transfer in satoshis, transaction fee is zero for pungo token
	amount := uint64(*argAmount * math.Pow10(int(8)))
	txFee := int64(0)

	// select utxo
	isUTXOSelect := false
	utxoSelect := make([]*utxorq.UTXOValue, 0)
	for _, v := range utxos.Result {
		if uint64(v.Value+txFee) >= amount {
			utxoSelect = append(utxoSelect, &v)
			isUTXOSelect = true
			break
		}
	}
	if !isUTXOSelect {
		println("not enougth value in available single utxos to complete transaction\n")
		return
	}

	// build OP_RETURN transaction
	transaction, err := cashburner.Build(networks["PGT"], *argKey, *argDest, *argSource, int(amount), int(txFee), utxoSelect, *argMessage)
	if err != nil {
		println("error building transaction: ", err.Error())
		return
	}
	println("transaction: ", transaction)
	println("\n")

	// push transaction
	formData := url.Values{
		"ip":    {"electrum1.komodo.build"},
		"port":  {"10002"},
		"rawtx": {transaction},
	}
	resp, err := http.PostForm("https://electrum.pungo.network/api/pushtx", formData)
	if err != nil {
		log.Fatalln("error sending transaction: ", err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["msg"].(string) == "success" {
		explorer := "https://pgt.komodo.build/tx/"
		fmt.Println("\ntransaction hash:", result["result"].(string))
		fmt.Println("transaction explorer:", explorer+result["result"].(string))
	} else {
		fmt.Println(result)
	}

}
