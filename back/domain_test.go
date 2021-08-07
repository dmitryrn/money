package main

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestService_NewTx(t *testing.T) {
	s := Service{}

	debitCard := Account{
		ID: 1,
		IdLessAccount: IdLessAccount{
			Name:    "Debit card",
			Balance: decimal.RequireFromString("100"),
		},
	}
	groceriesCategory := Category{
		ID: 1,
		IdLessCategory: IdLessCategory{
			Name:              "groceries",
			AvailableForSpend: decimal.Zero,
		},
	}

	sum := decimal.RequireFromString("-10")
	idLessTx, debitCard, groceriesCategory := s.NewTx(sum, debitCard, groceriesCategory)

	assert.Equal(t, IdLessTx{
		Sum:        sum,
		AccountID:  debitCard.ID,
		CategoryID: groceriesCategory.ID,
	}, idLessTx)

	assert.Equal(t, decimal.RequireFromString("90"), debitCard.Balance)

	assert.Equal(t, decimal.RequireFromString("-10"), groceriesCategory.AvailableForSpend)
}

func TestService_UpdateTx(t *testing.T) {
	s := Service{}

	t.Run("changing account (should update account only)", func(t *testing.T) {
		tx := Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-10"),
				AccountID:  1,
				CategoryID: 1,
			},
		}

		category := &Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("40"),
			},
		}

		debitCard := &Account{
			ID: 1,
			IdLessAccount: IdLessAccount{
				Name:    "Debit card",
				Balance: decimal.RequireFromString("90"),
			},
		}

		creditCard := &Account{
			ID: 2,
			IdLessAccount: IdLessAccount{
				Name:    "Credit card",
				Balance: decimal.RequireFromString("100"),
			},
		}

		tx, debitCard, creditCard, category, category = s.UpdateTx(tx, tx.Sum, debitCard, creditCard, category, category)

		assert.Equal(t, decimal.RequireFromString("-10"), tx.Sum)

		assert.Equal(t, decimal.RequireFromString("100"), debitCard.Balance)
		assert.Equal(t, decimal.RequireFromString("90"), creditCard.Balance)

		assert.Equal(t, Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("40"),
			},
		}, *category)

		assert.Equal(t, Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-10"),
				AccountID:  2,
				CategoryID: 1,
			},
		}, tx)
	})

	t.Run("changing category (should update category only)", func(t *testing.T) {
		tx := Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-10"),
				AccountID:  1,
				CategoryID: 1,
			},
		}

		groceries := &Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("40"),
			},
		}
		restaurants := &Category{
			ID: 2,
			IdLessCategory: IdLessCategory{
				Name:              "restaurants",
				AvailableForSpend: decimal.RequireFromString("50"),
			},
		}

		account := &Account{
			ID: 1,
			IdLessAccount: IdLessAccount{
				Name:    "account",
				Balance: decimal.RequireFromString("90"),
			},
		}

		tx, account, account, groceries, restaurants = s.UpdateTx(tx, tx.Sum, account, account, groceries, restaurants)

		assert.Equal(t, decimal.RequireFromString("-10"), tx.Sum)

		assert.Equal(t, decimal.RequireFromString("90"), account.Balance)

		assert.Equal(t, Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("50"),
			},
		}, *groceries)
		assert.Equal(t, Category{
			ID: 2,
			IdLessCategory: IdLessCategory{
				Name:              "restaurants",
				AvailableForSpend: decimal.RequireFromString("40"),
			},
		}, *restaurants)

		assert.Equal(t, Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-10"),
				AccountID:  1,
				CategoryID: 2,
			},
		}, tx)
	})

	t.Run("changing tx sum", func(t *testing.T) {
		tx := Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-10"),
				AccountID:  1,
				CategoryID: 1,
			},
		}

		category := &Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("40"),
			},
		}

		account := &Account{
			ID: 1,
			IdLessAccount: IdLessAccount{
				Name:    "account",
				Balance: decimal.RequireFromString("90"),
			},
		}

		tx, account, account, category, category = s.UpdateTx(tx, decimal.RequireFromString("-15"), account, account, category, category)

		assert.Equal(t, Account{
			ID: 1,
			IdLessAccount: IdLessAccount{
				Name:    "account",
				Balance: decimal.RequireFromString("85"),
			},
		}, *account)

		assert.Equal(t, Category{
			ID: 1,
			IdLessCategory: IdLessCategory{
				Name:              "groceries",
				AvailableForSpend: decimal.RequireFromString("35"),
			},
		}, *category)

		assert.Equal(t, Tx{
			ID: 1,
			IdLessTx: IdLessTx{
				Sum:        decimal.RequireFromString("-15"),
				AccountID:  1,
				CategoryID: 1,
			},
		}, tx)
	})

	// todo: changing account and category
	// todo: changing account and sum
	// todo: changing category and sum
	// todo: changing account and category and sum
}

func TestService_revertTx(t *testing.T) {
	s := Service{}

	tx := Tx{
		ID: 1,
		IdLessTx: IdLessTx{
			Sum:        decimal.RequireFromString("-10"),
			AccountID:  1,
			CategoryID: 1,
		},
	}
	account := Account{
		ID: 1,
		IdLessAccount: IdLessAccount{
			Name:    "account",
			Balance: decimal.RequireFromString("90"),
		},
	}
	category := Category{
		ID: 1,
		IdLessCategory: IdLessCategory{
			Name:              "groceries",
			AvailableForSpend: decimal.RequireFromString("50"),
		},
	}

	account, category = s.revertTx(tx, account, category)

	assert.Equal(t, decimal.RequireFromString("100"), account.Balance)

	assert.Equal(t, decimal.RequireFromString("60"), category.AvailableForSpend)
}
