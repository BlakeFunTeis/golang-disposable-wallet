package Ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

type EthWalletManager struct{}

func (ewm *EthWalletManager) CreateWallet() (_address string, _privateKey string, _err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	privateKeyHex := hexutil.Encode(crypto.FromECDSA(privateKey))

	return address, privateKeyHex, nil
}

func (ewm *EthWalletManager) GetBalance(_address string) (balance float64, err error) {
	key := os.Getenv("INFURA_API_KEY")
	MainEndpoint := os.Getenv("ETHEREUM_MAIN_NET")
	client, err := ethclient.Dial(MainEndpoint + "/" + key)
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(_address)
	balanceAt, err := client.BalanceAt(context.Background(), address, nil)
	balance = toEther(balanceAt.Int64())
	if err != nil {
		panic(err)
	}
	return balance, err
}

func (ewm *EthWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error) {
	return "", nil
}

func (ewm *EthWalletManager) DestroyWallet(_address string) (err error) {
	return nil
}

func toEther(_wei int64) (_ether float64) {
	_ether = float64(_wei) / 1000000000000000000
	return _ether
}
