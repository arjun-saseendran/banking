package dto

import "github.com/arjun-saseendran/banking/errs"

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId string `json:"account_id`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type`
	TransactionDate string `json:"transaction_date`
	CustomerId string `json:"customer_id`
	
}
 
func (r TransactionRequest) IsTransactionTypeWithDrawal()bool{
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) IsTransactionTypeDeposit()bool{
	return  r.TransactionType == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError{
	if !r.IsTransactionTypeWithDrawal() && !r.IsTransactionTypeDeposit(){
		return  errs.NewValidationError("Transaction type only be deposit or withdrawal!")
	}
	if r.Amount < 0{
		return errs.NewValidationError("Amount cannot be lessthan zero!")
	}
	return nil
}

type TransactionResponse struct{
	TransactionId string `json:"transaction_id`
	AccountId string `json:"account_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
	
}