package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
)

func keyToPrivateKey(key string) *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		panic(err)
	}
	return privateKey
}

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	pkeybytes := crypto.FromECDSA(privateKey)
	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey"))
	}
	addr := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println("pkey:", hex.EncodeToString(pkeybytes))
	fmt.Println("pubaddr:", addr)
}
