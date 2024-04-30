package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// A transaction consists of the amount of ether you're transferring, the gas limit, the gas price, a nonce, the receiving address, and optionally data.
	// The transaction must be signed with the private key of the sender before it's broadcasted to the network.
	// Using localhost set up by ganache to test actual transfer
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Load private key
	privateKey, err := crypto.HexToECDSA("f1b3f8e0d52caec13491368449ab8d90f3d222a3e485aa7f02591bbceb5efba5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Every transaction needs a nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// The next step is to set the amount of ETH that we'll be transferring.
	// However we must convert ether to wei since that's what the Ethereum blockchain uses.
	value := big.NewInt(1000000000000000000) // in wei (1 eth)

	// The gas limit for a standard ETH transfer is 21000 units.
	gasLimit := uint64(21000) // in units

	// The gas price must be set in wei.
	// At the time of this tutorial getting written, a gas price that will get your transaction included pretty fast in a block is 30 gwei.
	// gasPrice := big.NewInt(30000000000) // in wei (30 gwei)

	// gasPrice can be set to a suggested value to account for market fluctuations
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x68dB32D26d9529B2a142927c6f1af248fc6Ba7e9")

	// Generate unsigned transaction
	// Data field is nil for sending ETH. It is used while interacting with smart contracts
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	// Sign the transaction with the private key of the sender
	// It requires the EIP155 signer, which we derive the chain ID from the client.
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Broadcast the transaction to the entire network by calling SendTransaction on the client which takes in the signed transaction

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())

}
