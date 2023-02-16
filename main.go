package main

import (
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"golang-disposable-wallet/blockchain"
	"golang-disposable-wallet/blockchain/Ethereum"
)

func main() {
	var walletManager blockchain.WalletManager
	walletManager = &Ethereum.EthWalletManager{}
	ethAddress, ethPrivateKey, _ := walletManager.CreateWallet()
	balance, _ := walletManager.GetBalance(ethAddress)
	fmt.Printf("Address: %s\n", ethAddress)
	fmt.Printf("Private Key: %s\n", ethPrivateKey)
	fmt.Printf("Balance: %f\n", balance)
	//walletManager = &Solana.SolWalletManager{}
	//solAddress, solPrivateKey, _ := walletManager.CreateWallet()
	//walletManager = &Tron.TronWalletManager{}
	//tronAddress, tronPrivateKey, _ := walletManager.CreateWallet()
	//fmt.Printf("Address: %s\n", solAddress)
	//fmt.Printf("Private Key: %s\n", solPrivateKey)
	//fmt.Printf("Address: %s\n", tronAddress)
	//fmt.Printf("Private Key: %s\n", tronPrivateKey)
}
