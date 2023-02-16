package Tron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcutil/base58"
	ethCommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	AddressLength = 21
	AddressPrefix = "41"
)

type Address [AddressLength]byte

type TronWalletManager struct{}

func (twm *TronWalletManager) CreateWallet() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	address, err := getTronAddress(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}

	privateKeyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))
	tronAddressString, err := encode58Check(address.Bytes())
	if err != nil {
		return "", "", err
	}

	return tronAddressString, privateKeyHex, nil
}

func (twm *TronWalletManager) GetBalance(address string) (float64, error) {
	return 0.0, nil
}

func (twm *TronWalletManager) SendTransaction(fromAddress string, toAddress string, amount float64) (string, error) {
	return "", nil
}

func (twm *TronWalletManager) DestroyWallet(address string) error {
	return nil
}

func getTronAddress(publicKey *ecdsa.PublicKey) (Address, error) {
	address := crypto.PubkeyToAddress(*publicKey)
	tronAddress, err := addressLedgerToTron(address.Bytes())
	if err != nil {
		return Address{}, err
	}

	return tronAddress, nil
}

func addressLedgerToTron(ledgerAddress []byte) (Address, error) {
	addr := ethCommon.BytesToAddress(crypto.Keccak256(ledgerAddress[1:])[12:])
	addressTron := make([]byte, AddressLength)
	addressPrefix, err := FromHex(AddressPrefix)
	if err != nil {
		return Address{}, err
	}
	addressTron = append(addressTron, addressPrefix...)
	addressTron = append(addressTron, addr.Bytes()...)
	return BytesToAddress(addressTron), nil
}

func FromHex(input string) ([]byte, error) {
	if len(input) == 0 {
		return nil, errors.New("empty hex string")
	}

	return hex.DecodeString(input[:])
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func (a *Address) SetBytes(b []byte) {
	if len(b) > len(a) {
		b = b[len(b)-AddressLength:]
	}
	copy(a[AddressLength-len(b):], b)
}

func (a *Address) Bytes() []byte {
	return a[:]
}

func encode58Check(input []byte) (string, error) {
	h0, err := Hash(input)
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
	inputCheck := append(input, h1[:4]...)

	return base58.Encode(inputCheck), nil
}

func Hash(s []byte) ([]byte, error) {
	h := sha256.New()
	_, err := h.Write(s)
	if err != nil {
		return nil, err
	}
	bs := h.Sum(nil)
	return bs, nil
}
