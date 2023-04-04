package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

//Executes a function within a database context
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %+v, rollback error: %+v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) setupBalanceInventory(ctx context.Context, accountID int64, amount int64, isSender bool) (updatedAccount Account, entry Entry, err error) {
	// Update the  fromAccount balance
	account, _ := store.GetAccount(ctx, accountID)

	var argsUpdate UpdateAccountParams
	// Create update args param

	if isSender {
		argsUpdate = UpdateAccountParams{
			ID:      accountID,
			Balance: account.Balance - amount,
		}
	} else {
		argsUpdate = UpdateAccountParams{
			ID:      accountID,
			Balance: account.Balance + amount,
		}
	}

	// Deduct the balance from the fromAccount
	if updatedAccount, err = store.UpdateAccount(ctx, argsUpdate); err != nil {
		return Account{}, Entry{}, fmt.Errorf("error updating the account: %d", accountID)
	}

	var argsEntry CreateEntryParams

	if isSender {
		argsEntry = CreateEntryParams{
			AccountID: accountID,
			Amount:    -amount,
		}
	} else {
		argsEntry = CreateEntryParams{
			AccountID: accountID,
			Amount:    amount,
		}
	}

	// Create an entry
	if entry, err = store.CreateEntry(ctx, argsEntry); err != nil {
		return Account{}, Entry{}, fmt.Errorf("cannot create entry for the transaction for account id %d", accountID)
	}

	return updatedAccount, entry, nil

}

// Transfer tx performs a money transfer from one account to another
// It creates a transfer record, add account entries, and updates account's balance within a single db transaction

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		fromUpdatedAccount, fromEntry, err := store.setupBalanceInventory(ctx, arg.FromAccountID, arg.Amount, true)

		if err != nil {
			return err
		}
		toUpdatedAccount, toEntry, err := store.setupBalanceInventory(ctx, arg.ToAccountID, arg.Amount, false)
		if err != nil {
			return err
		}

		argsTransfer := CreateTransferParams(arg)
		// Create the transfer
		transfer, err := store.CreateTransfer(ctx, argsTransfer)
		if err != nil {
			return err
		}

		result = TransferTxResult{
			Transfer:    transfer,
			FromAccount: fromUpdatedAccount,
			ToAccount:   toUpdatedAccount,
			FromEntry:   fromEntry,
			ToEntry:     toEntry,
		}

		return nil

	})

	return result, err
}
