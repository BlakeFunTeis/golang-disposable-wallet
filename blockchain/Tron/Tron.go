package Tron

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	protocol_api "github.com/tron-us/go-btfs-common/protos/protocol/api"
	"google.golang.org/grpc"
)

const (
	AddressLength = 21
	AddressPrefix = "41"
)

type Address [AddressLength]byte

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
	return 0.0, nil
}

func (twm *TronWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (string, error) {
	return "", nil
}

func (twm *TronWalletManager) DestroyWallet(_address string) error {
	return nil
}

func getTronAddress(_publicKey *ecdsa.PublicKey) (Address, error) {
	address := crypto.PubkeyToAddress(*_publicKey)
	tronAddress, err := addressLedgerToTron(address.Bytes())
	if err != nil {
		return Address{}, err
	}

	return tronAddress, nil
}

func addressLedgerToTron(_ledgerAddress []byte) (Address, error) {
	addr := ethCommon.BytesToAddress(crypto.Keccak256(_ledgerAddress[1:])[12:])
	addressTron := make([]byte, AddressLength)
	addressPrefix, err := FromHex(AddressPrefix)
	if err != nil {
		return Address{}, err
	}
	addressTron = append(addressTron, addressPrefix...)
	addressTron = append(addressTron, addr.Bytes()...)
	return BytesToAddress(addressTron), nil
}

func FromHex(_input string) ([]byte, error) {
	if len(_input) == 0 {
		return nil, errors.New("empty hex string")
	}

	return hex.DecodeString(_input[:])
}

func BytesToAddress(_byte []byte) Address {
	var a Address
	a.SetBytes(_byte)
	return a
}

func (a *Address) SetBytes(_byte []byte) {
	if len(_byte) > len(a) {
		_byte = _byte[len(_byte)-AddressLength:]
	}
	copy(a[AddressLength-len(_byte):], _byte)
}

func (a *Address) Bytes() []byte {
	return a[:]
}

func encode58Check(_input []byte) (string, error) {
	h0, err := Hash(_input)
	if err != nil {
		return "", err
	}
	h1, err := Hash(h0)
	if err != nil {
		return "", err
	}
	if len(h1) < 4 {
		return "", errors.New("base58 encode length error")
	}
	inputCheck := append(_input, h1[:4]...)

	return base58.Encode(inputCheck), nil
}

func Hash(_byte []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(_byte)
	if err != nil {
		return nil, err
	}
	bs := h.Sum(nil)
	return bs, nil
}
