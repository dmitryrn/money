package main

import (
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	ErrNegativeNewAccountBalance = errors.New("new account balance can not be negative")
)

type (
	IdLessAccount struct {
		Name    string          `json:"name"`
		Balance decimal.Decimal `json:"balance"`
	}

	Account struct {
		ID uint64 `json:"id"`
		IdLessAccount
	}

	IdLessTx struct {
		Sum        decimal.Decimal `json:"sum"`
		AccountID  uint64          `json:"account_id"`
		CategoryID uint64          `json:"category_id"`
	}

	Tx struct {
		ID uint64 `json:"id"`
		IdLessTx
	}

	Category struct {
		ID                uint64          `json:"id"`
		Name              string          `json:"name"`
		AvailableForSpend decimal.Decimal `json:"available_for_spend"`
	}
)

type Service struct {
}

func (s *Service) NewTx(
	sum decimal.Decimal,
	account Account,
	category Category,
) (
	IdLessTx,
	Account,
	Category,
) {
	newTx := IdLessTx{
		Sum:        sum,
		AccountID:  account.ID,
		CategoryID: category.ID,
	}

	account.Balance = account.Balance.Add(sum)
	category.AvailableForSpend = category.AvailableForSpend.Add(sum)

	return newTx, account, category
}

func (s *Service) UpdateTx(
	tx Tx,
	newSum decimal.Decimal,
	account *Account,
	newAccount *Account,
	category *Category,
	newCategory *Category,
) (
	Tx,
	*Account,
	*Account,
	*Category,
	*Category,
) {
	if account == nil || newAccount == nil || category == nil || newCategory == nil {
		panic("some of the required params is nil")
	}

	*account, *category = s.revertTx(tx, *account, *category)

	tx.Sum = newSum

	_, *newAccount, *newCategory = s.NewTx(tx.Sum, *newAccount, *newCategory)
	tx.CategoryID = newCategory.ID
	tx.AccountID = newAccount.ID

	return tx, account, newAccount, category, newCategory
}

func (s *Service) revertTx(
	tx Tx,
	account Account,
	category Category,
) (
	Account,
	Category,
) {
	account.Balance = account.Balance.Sub(tx.Sum)
	category.AvailableForSpend = category.AvailableForSpend.Sub(tx.Sum)

	return account, category
}

func (s *Service) TransferBudgeted(
	amount decimal.Decimal,
	categoryA Category,
	categoryB Category,
) (
	Category,
	Category,
) {
	categoryA.AvailableForSpend = categoryA.AvailableForSpend.Sub(amount)
	categoryB.AvailableForSpend = categoryB.AvailableForSpend.Add(amount)

	return categoryA, categoryB
}

func (s *Service) NewAccount(
	name string,
	balance decimal.Decimal,
	availableForBudgetingCategory Category,
) (
	IdLessAccount,
	Category,
	error,
) {
	if balance.IsNegative() {
		return IdLessAccount{}, Category{}, ErrNegativeNewAccountBalance
	}

	account := IdLessAccount{
		Name:    name,
		Balance: balance,
	}

	availableForBudgetingCategory.AvailableForSpend = availableForBudgetingCategory.AvailableForSpend.Add(balance)

	return account, availableForBudgetingCategory, nil
}
