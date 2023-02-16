package Solana

import (
	"context"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

type SolWalletManager struct{}

func (swm *SolWalletManager) CreateWallet() (_address string, _privateKey string, _err error) {
	account := types.NewAccount()
	privateKeyBytes := account.PrivateKey
	privateKey := hexutil.Encode(privateKeyBytes)
	publicKey := account.PublicKey.ToBase58()
	return publicKey, privateKey, nil
}

func (swm *SolWalletManager) GetBalance(_address string) (balance float64, err error) {
	client := rpc.NewRpcClient(rpc.MainnetRPCEndpoint)
	json, err := client.GetBalance(context.Background(), _address)

	if err != nil {
		return 0.0, err
	}

	balance = float64(json.Result.Value) / 1000000000
	return balance, nil
}

func (swm *SolWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error) {
	return "", nil
}

func (swm *SolWalletManager) DestroyWallet(_address string) (err error) {
	return nil
}
