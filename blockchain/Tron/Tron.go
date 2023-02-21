package Tron

type TronWalletManager struct{}

func (twm *TronWalletManager) CreateWallet() (string, string, error) {
	return "", "", nil
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
