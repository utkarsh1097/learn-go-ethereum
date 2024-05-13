package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// A "block" consists of a Header and a Body.
	// https://www.educative.io/answers/what-are-blocks-in-blockchain-technology
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	// Only block header
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744

	// Full block
	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1527211625
	fmt.Println(block.Difficulty().Uint64()) // 3217000136609065
	fmt.Println(block.Hash().Hex())          // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	fmt.Println(len(block.Transactions()))   // 144

	// Call TransactionCount to return just the countTransactions of transactions in a block.
	countTransactions, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(countTransactions) // 144 == len(block.Transactions())

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over all transactions
	fmt.Println("Iterating over all block transactions")
	for _, tx := range block.Transactions() {
		fmt.Println(tx.Hash().Hex())        // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
		fmt.Println(tx.Value().String())    // 10000000000000000
		fmt.Println(tx.Gas())               // 105000
		fmt.Println(tx.GasPrice().Uint64()) // 102000000000
		fmt.Println(tx.Nonce())             // 110644
		fmt.Println(tx.Data())              // []
		fmt.Println(tx.To().Hex())          // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		// https://ethereum.stackexchange.com/questions/149220/type-types-transaction-has-no-field-or-method-asmessage
		// The snippet below from tutorial is stale. Use [types.Sender]
		/**
		    if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		        fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		    }
		**/

		if fromAddress, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println(fromAddress.Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		}

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(receipt.Status) // 1
		fmt.Println(receipt.Logs)

	}

	// Another way to iterate over transaction without fetching the block is to call the client's TransactionInBlock method.
	// This method accepts only the block hash and the index of the transaction within the block.
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")

	for idx := uint(0); idx < countTransactions; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0xd3a7283055eb1aa9f9511c635637afadbde96265ae91b72b0419a3cbc0ce1a73
	}

	// You can also query for a single transaction directly given the transaction hash by using TransactionByHash.
	txHash := common.HexToHash("0xd3a7283055eb1aa9f9511c635637afadbde96265ae91b72b0419a3cbc0ce1a73")
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0xd3a7283055eb1aa9f9511c635637afadbde96265ae91b72b0419a3cbc0ce1a73
	fmt.Println(isPending)       // false

}
