package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

const GOERLI_ENDPOINT = "https://goerli.infura.io/v3/ad0b336d7f2f4082b5a624e50d27df5c"

func main() {
	ethClient, _ := ethclient.Dial(GOERLI_ENDPOINT)
	chainId, _ := ethClient.ChainID(context.Background())

	fmt.Printf("Connected to Goerli - Chain ID: %v", chainId)
}
