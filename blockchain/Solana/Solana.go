package Solana

type SolWalletManager struct{}

func (swm *SolWalletManager) CreateWallet() (_address string, _privateKey string, _err error) {
	return "", "", nil
}

func (swm *SolWalletManager) GetBalance(_address string) (balance float64, err error) {
	return 0.0, nil
}

func (swm *SolWalletManager) SendTransaction(_fromAddress string, _toAddress string, _amount float64) (txHash string, err error) {
	return "", nil
}

func (swm *SolWalletManager) DestroyWallet(_address string) (err error) {
	return nil
}
