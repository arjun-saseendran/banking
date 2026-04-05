package service

import (
	"github.com/arjun-saseendran/banking/domain"
	"github.com/arjun-saseendran/banking/dto"
	"github.com/arjun-saseendran/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	account := domain.NewAccount(req.CustomerId, req.AccountType, req.Amount)

	if newAccount, err := s.repo.Save(account); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil

	}
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.IsTransactionTypeWithDrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewUnexpectedError("Insufficient balance in the account!")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: req.TransactionDate,
	}
	transaction, err := s.repo.SaveTransaction(t)
	if err != nil {
		return nil, err
	}
	response := transaction.ToDto()
	return &response, nil

}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
