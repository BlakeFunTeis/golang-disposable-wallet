package main

import (
	"fmt"
	_ "github.com/joho/godotenv"
	"golang-disposable-wallet/blockchain"
	"golang-disposable-wallet/blockchain/Ethereum"
)

func main() {
	var walletManager blockchain.WalletManager
	walletManager = &Ethereum.EthWalletManager{}
	ethAddress, ethPrivateKey, _ := walletManager.CreateWallet()
	fmt.Printf("Address: %s\n", ethAddress)
	fmt.Printf("Private Key: %s\n", ethPrivateKey)
}
