![Pungo Token](https://pungotoken.sale/images/token_logo.png)

# OP_RETURN transactions tool

This script can be used to create transaction that include data using OP_RETURN NullData script.

Currently script is hardcoded for PGT (Pungo token) but can be expended for other bitcoin based blockchain networks.

Transaction OP_RETURN is created from single UTXO, so currently you cannot merge several UTXOs to burn larger amount. 

## Options

```
$cashburner -key="123" -src="456" -val=0.1 -msg="message"

src: wallet address
key: wallet private key
val: amount to be burned with OP_RETURN output
msg: message that will be encoded with OP_RETURN output
```

## Get code

For Debian linux there is a prebuilt binary:

```
$git clone git@gitlab.com:instak_backend/op_return.git
$cd op_return

# now you can run the script
$cashburner -val=0.1 -key="UtKDtRQgYAEyybAbNhtCEZy2iSUDhejBhwuawPFcnB1YQXzwH8u61" -src="RXyrzo7iS6iN69m6ZFhQYWbZAXZQxzifAY" -msg="in a galaxy far far away.."
```

Build manually:

```
# make sure you have golang version 1.8+ installed
$go version

# install golang
$apt-get update
$apt-get install golang

$git clone git@gitlab.com:instak_backend/op_return.git
$cd op_return
$dep ensure
$go build ./cashburner.go 

$cashburner -val=0.1 -key="UtKDtRQgYAEyybAbNhtCEZy2iSUDhejBhwuawPFcnB1YQXzwH8u61" -src="RXyrzo7iS6iN69m6ZFhQYWbZAXZQxzifAY" -msg="in a galaxy far far away.."
```

## View transactions

Currently PGT Explorer (insight explorer) doesn't show messages included with OP_RETURN.

To decode OP_RETURN transaction and verify message use a komodo full node client:

```
$./komodo-cli getrawtransaction verbose=1 TX_HASH
```

## Notes

To be on a safe side, for now do not use a wallet (private key) if you'll have large amounts as change from transaction.
