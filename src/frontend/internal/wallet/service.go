package wallet

type Service struct {
	walletRepository WalletRepository
}

type WalletService interface {
	GetUserWallets() (*[]Wallet, error)
}

func NewWalletService(walletRepository WalletRepository) *Service {
	return &Service{walletRepository}
}

func (s *Service) GetUserWallets() (*[]Wallet, error) {
	return s.walletRepository.FindAll()
}
