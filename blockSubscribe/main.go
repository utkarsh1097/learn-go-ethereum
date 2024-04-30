package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

// How to set up subscription to get events when there is a new block mined
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("INFURA_API_KEY")

	// Need Ethereum provider that supports RPC over websockets
	// Infura endpoint to access Ethereum network
	client, err := ethclient.Dial(fmt.Sprintf("wss://sepolia.infura.io/ws/v3/%v", apiKey))
	if err != nil {
		log.Fatal(err)
	}

	// Subscription for the latest block headers
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// Use a select statement to listen for new messages
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:

			fmt.Println(header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex()) // Must be equal to header.Hash().Hex()
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}

	// To get the full contents of the block, we can pass the block header hash to the client's BlockByHash function.

}
