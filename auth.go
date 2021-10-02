package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"time"
)

type AuthManager struct {
	keystore  *keystore.KeyStore
	account   accounts.Account
	createdAt time.Time
}

// NewAuthManager initializes a new keystore and creates a new ethereum
// account
func NewAuthManager(ctx context.Context) (*AuthManager, error) {
	ks := keystore.NewKeyStore("", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := ks.NewAccount(*passphrase)
	if err != nil {
		return nil, err
	}

	return &AuthManager{
		account:   account,
		keystore:  ks,
		createdAt: time.Now(),
	}, nil
}
