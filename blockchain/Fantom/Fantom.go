package Fantom

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type FtmWalletManager struct{}

func (fwm *FtmWalletManager) CreateWallet() (ftmAddress string, ftmPrivateKey string, err error) {
	//var privateKey *ecdsa.PrivateKey
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}
	ftmAddress = crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return ftmAddress, privateKey.D.String(), nil
}

func (fwm *FtmWalletManager) GetBalance(_address string) (balance float64, err error) {
	rpcUrl := "https://rpcapi.fantom.network"
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return 0.0, err
	}

	ctx := context.Background()
	account := common.HexToAddress(_address)
	result, err := client.BalanceAt(ctx, account, nil)
	if err != nil {
		return 0.0, err
	}

	balanceInFTM := new(big.Float).SetInt(result)
	divisor := new(big.Float).SetFloat64(1000000000000000000)
	balanceInFTM = balanceInFTM.Quo(balanceInFTM, divisor)
	balance, _ = balanceInFTM.Float64()
	return balance, nil
}

func (fwm *FtmWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error) {
	return "", nil
}

func (fwm *FtmWalletManager) DestroyWallet(_address string) (err error) {
	return nil
}
