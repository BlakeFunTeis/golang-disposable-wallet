package Tron

import (
	"context"
	protocol_api "github.com/tron-us/go-btfs-common/protos/protocol/api"
	"google.golang.org/grpc"
)

type TronWalletManager struct{}

func (twm *TronWalletManager) CreateWallet() (string, string, error) {
	conn, err := grpc.Dial("grpc.trongrid.io:50051", grpc.WithInsecure())
	if err != nil {
		return "", "", err
	}

	defer conn.Close()
	client := protocol_api.NewWalletClient(conn)
	ctx := context.Background()
	if err != nil {
		return "", "", err
	}
	resp, err := client.GenerateAddress(ctx, &protocol_api.EmptyMessage{})

	return resp.GetAddress(), resp.GetPrivateKey(), nil
}

func (twm *TronWalletManager) GetBalance(_address string) (float64, error) {
	conn, err := grpc.Dial("grpc.trongrid.io:50051", grpc.WithInsecure())
	if err != nil {
		return 0.0, err
	}

	defer conn.Close()
	client := protocol_api.NewWalletClient(conn)
	ctx := context.Background()
	if err != nil {
		return 0.0, err
	}
}

func (twm *TronWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (string, error) {
	return "", nil
}

func (twm *TronWalletManager) DestroyWallet(_address string) error {
	return nil
}
