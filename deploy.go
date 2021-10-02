package main

import (
	"context"
	"fmt"
	"github.com/darthchudi/dante/bindings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"os"
)

// Deploy creates a new ethereum account and deploys the smart contract
// with the address
func (s *Service) deploy(ctx context.Context, pathToKey, passphrase string) error {
	keyReader, err := os.Open(pathToKey)
	if err != nil {
		return err
	}
	defer keyReader.Close()

	auth, err := bind.NewTransactorWithChainID(keyReader, passphrase, big.NewInt(goerliTestnetChainID))
	if err != nil {
		return err
	}

	address, transaction, store, err := bindings.DeployStore(auth, s.client)
	if err != nil {
		return err
	}
	s.store = store

	fmt.Println("✨ Contract address (pending deploy): ", address.Hex())
	fmt.Println("🏦 Transaction hash (waiting to be mined): ", transaction.Hash())
	fmt.Println("🚀 Successfully deployed contract :)")

	return nil
}
