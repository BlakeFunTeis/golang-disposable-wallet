package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"golang-disposable-wallet/routes"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ginInstance := gin.New()
	ginInstance.Use(gin.Logger())
	ginInstance.Use(gin.Recovery())
	gin.SetMode(getMode())
	routesInstance := routes.WebRouteInstance(ginInstance)
	ReadTimeOut, _ := time.ParseDuration(os.Getenv("HTTP_READ_TIMEOUT"))
	WriteTimeOut, _ := time.ParseDuration(os.Getenv("HTTP_WRITE_TIMEOUT"))
	server := &http.Server{
		Addr:           ":" + os.Getenv("HTTP_PORT"),
		Handler:        routesInstance,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("listen: %s\n", err)
	}
	//var walletManager blockchain.WalletManager
	//walletManager = &Fantom.FtmWalletManager{}
	//address, _, _ := walletManager.CreateWallet()
	//walletManager.GetBalance(address)
	//walletManager = &Tron.TronWalletManager{}
	//tronAddress, tronPrivateKey, _ := walletManager.CreateWallet()
	//walletManager.GetBalance("TR7NHqjeKQxGTCi8q8ZY4pL8otSzgjLj6t")
	//walletManager = &Ethereum.EthWalletManager{}
	//ethAddress, ethPrivateKey, _ := walletManager.CreateWallet()
	//balance, _ := walletManager.GetBalance(ethAddress)
	//fmt.Printf("Address: %s\n", ethAddress)
	//fmt.Printf("Private Key: %s\n", ethPrivateKey)
	//fmt.Printf("Balance: %f\n", balance)
	//walletManager = &Solana.SolWalletManager{}
	//solAddress, solPrivateKey, _ := walletManager.CreateWallet()
	//balance, _ := walletManager.GetBalance("EvxyoLY9PwVZmZQPTbkxJDzCiq415Ajyvjg1MyJvogA")
	//fmt.Printf("Balance: %f\n", balance)
	//walletManager = &Tron.TronWalletManager{}
	//tronAddress, tronPrivateKey, _ := walletManager.CreateWallet()
	//fmt.Printf("Address: %s\n", solAddress)
	//fmt.Printf("Private Key: %s\n", solPrivateKey)
	//fmt.Printf("Tron Address: %s\n", tronAddress)
	//fmt.Printf("Tron Private Key: %s\n", tronPrivateKey)
}

func getMode() string {
	var mode = os.Getenv("app_env")
	switch strings.ToLower(mode) {
	case "debug":
		return gin.DebugMode
	case "local":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
