package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// Serialize the private key to byte array
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// Encode byte array to hex string
	// Strip "0x" (this is always the prefix)
	privateKeyHex := hexutil.Encode(privateKeyBytes)[2:]

	fmt.Printf("Your wallet's private key is %v\n", privateKeyHex)

	// Public key corresponding to this private key

	// To assert public key type
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Strip "0x04" (this is always the prefix)
	publicKeyHex := hexutil.Encode(publicKeyBytes)[4:]

	fmt.Printf("Your wallet's public key is %v\n", publicKeyHex)

	// Public key is used to generate the wallet address
	// For ethereum blockchain, the public address is simply the Keccak-256 hash of the public key, and then we take the last 40 characters (20 bytes) and prefix it with 0x
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("Your wallet's address is %v\n", address) // 0x96216849c49358B10257cb55b28eA603c874b05E

	// Generating the address by ourself
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	manAddress := hexutil.Encode(hash.Sum(nil)[12:])
	fmt.Printf("If generated manually, your wallet's address is %v\n", manAddress) // 0x96216849c49358B10257cb55b28eA603c874b05E
}
