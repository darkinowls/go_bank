package sqlc

import (
	"context"
	"database/sql"
	"fmt"
)

// embedding (compostition + interface instead of inheritance)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, globalErr error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

var TkKey = struct{}{}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		Queries: New(db),
		db:      db,
	}
}

// execute a function within a database transaction
func (s *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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

func (s *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (result TransferTxResult, globalErr error) {
	globalErr = s.execTx(ctx, func(q *Queries) error {
		var e error
		result.Transfer, e = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if e != nil {
			return e
		}
		result.FromEntry, e = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if e != nil {
			return e
		}
		result.ToEntry, e = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if e != nil {
			return e
		}

		// update account balance
		// ORDER MATTERS to avoid deadlocks (cross deadlocks)
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, e = addMoney(ctx, q, arg.FromAccountID, arg.ToAccountID,
				-arg.Amount, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, e = addMoney(ctx, q, arg.ToAccountID, arg.FromAccountID,
				arg.Amount, -arg.Amount)
		}
		if e != nil {
			return e
		}

		return nil
	})
	return result, globalErr
}

func addMoney(ctx context.Context, q *Queries,
	accountID1 int64, accountID2 int64,
	amount1 int64, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return

}
