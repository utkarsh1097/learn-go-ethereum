package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Keystores in go-ethereum can only contain one wallet key pair per file. However, they can contain multiple files.
func main() {
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)

	// NewKeyStore at a directory with existing key files will result in their auto-load
	// ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)

	password := "yadayadahogehoge"
	// account, err := ks.NewAccount(password)
	// if err != nil {
	// 	log.Fatal("Unable to create new account")

	// }

	currAddrPath := "UTC--2024-04-22T02-04-51.580627000Z--fa2d187a4552ab44272955610f909934b36072e2"

	fileData, err := ioutil.ReadFile(path.Join("wallets", currAddrPath))
	if err != nil {
		log.Fatalf("Could not read the file. Details: \v\n", err)
	}

	fmt.Println(ks.Accounts())
	fmt.Println("OK")

	account, err := ks.Import(fileData, password, password)
	if err != nil {
		log.Fatalf("Unable to import data. Details: \v\n", err)
	}

	fmt.Println(account.Address.Hex()) // 0x20F8D42FB0F667F2E53930fed426f225752453b3

	// if err := os.Remove(path.Join("./wallets/tmp", currAddrPath)); err != nil {
	// 	log.Fatal(err)
	// }

}
