package Ethereum

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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
	return 0.0, nil
}

func (ewm *EthWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error) {
	return "", nil
}

func (ewm *EthWalletManager) DestroyWallet(_address string) (err error) {
	return nil
}
