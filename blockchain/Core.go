package blockchain

type WalletManager interface {
	CreateWallet() (_address string, _privateKey string, _err error)
	GetBalance(_address string) (balance float64, err error)
	SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error)
	DestroyWallet(_address string) (err error)
}
