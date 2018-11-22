package utxorq

// UTXOResponse .
type UTXOResponse struct {
	Msg    string      `json:"msg"`
	Result []UTXOValue `json:"result"`
}

// UTXOValue .
type UTXOValue struct {
	Hash   string `json:"tx_hash"`
	Pos    int    `json:"tx_pos"`
	Height int    `json:"height"`
	Value  int64  `json:"value"`
}
