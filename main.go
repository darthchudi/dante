package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const (
	clientAddress = "https://mainnet.infura.io"
	contractAddress = ""
)


func main () {
	ctx := context.Background()

	ethereumClient, err := ethclient.Dial(clientAddress)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}

	store, err := NewStore(common.HexToAddress(contractAddress), ethereumClient)
	if err != nil {
		log.Fatalf("Failed to initialize contract bindings: %v", err)
	}


}