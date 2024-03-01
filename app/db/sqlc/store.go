package db

import (
	"context"
	"database/sql"
	"fmt"
)

// embedding (compostition + interface instead of inheritance)
type Store struct {
	*Queries
	db *sql.DB
}

var TkKey = struct{}{}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// execute a function within a database transaction
func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("%v AND %v", err, rbErr)
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
	Transfer    Transfer
	FromAccount Account
	ToAccount   Account
	FromEntry   Entry
	ToEntry     Entry
}

func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, globalErr error) {
	globalErr = s.execTx(ctx, func(q *Queries) error {
		var e error
		txName := ctx.Value(TkKey)
		fmt.Println(txName, "creating transfer...")
		result.Transfer, e = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if e != nil {
			return e
		}
		fmt.Println(txName, "transfer create Entry from...")
		result.FromEntry, e = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if e != nil {
			return e
		}
		fmt.Println(txName, "transfer create Entry to...")
		result.ToEntry, e = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if e != nil {
			return e
		}

		//// TODO: update account balance
		fmt.Println(txName, "GetAccountForUpdate 1...")
		acc1, e := q.GetAccountForUpdate(ctx, arg.FromAccountID)
		if e != nil {
			return e
		}

		fmt.Println(txName, "UpdateAccount 1...")
		result.FromAccount, e = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      acc1.ID,
			Balance: acc1.Balance - arg.Amount,
		})
		if e != nil {
			return e
		}
		fmt.Println(txName, "GetAccountForUpdate 2...")
		acc2, e := q.GetAccountForUpdate(ctx, arg.ToAccountID)
		if e != nil {
			return e
		}

		fmt.Println(txName, "UpdateAccount 2...")
		result.ToAccount, e = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      acc2.ID,
			Balance: acc2.Balance + arg.Amount,
		})
		if e != nil {
			return e
		}

		return nil
	})
	return result, globalErr
}
