package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStore(t *testing.T) {
	store := NewStore(dbCon)

	//acc1, err := store.GetAccount(ctx, 1)
	//require.NoError(t, err)
	//acc2, err := store.GetAccount(ctx, 2)
	//require.NoError(t, err)

	acc1, _ := createRandomAccount(t)
	acc2, _ := createRandomAccount(t)

	// concurently run n instances of the same test
	n := 2
	errs := make(chan error)
	txs := make(chan TransferTxResult)
	amount := int64(10)

	for i := 0; i < n; i++ {
		name := fmt.Sprintf("transaction-%d", i+1)
		go func() {
			ctx := context.WithValue(context.Background(), TkKey, name)
			transferResult, err := store.TransferTx(ctx,
				TransferTxParams{
					FromAccountID: acc1.ID,
					ToAccountID:   acc2.ID,
					Amount:        amount,
				},
			)
			errs <- err
			txs <- transferResult
		}()
	}

	existed := make(map[int]bool)
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
		tx := <-txs
		require.NotEmpty(t, tx)
		require.NotZero(t, tx.Transfer.ID)
		require.Equal(t, acc1.ID, tx.FromEntry.AccountID)
		require.Equal(t, acc2.ID, tx.ToEntry.AccountID)

		// Transfer
		require.Equal(t, amount, tx.Transfer.Amount)
		transfer, err := store.GetTransfer(context.Background(), tx.Transfer.ID)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)
		require.NoError(t, err)

		// Entry
		entry, err := store.GetEntry(context.Background(), tx.FromEntry.ID)
		require.NoError(t, err)
		require.Equal(t, acc1.ID, entry.AccountID)
		require.Equal(t, -amount, entry.Amount)

		entry, err = store.GetEntry(context.Background(), tx.ToEntry.ID)
		require.NoError(t, err)
		require.Equal(t, acc2.ID, entry.AccountID)
		require.Equal(t, amount, entry.Amount)

		// check account balances
		fmt.Println(">> tx:", tx.FromAccount.Balance, tx.ToAccount.Balance)
		fromAccount := tx.FromAccount
		require.Equal(t, acc1.ID, fromAccount.ID)

		toAccount := tx.ToAccount
		require.Equal(t, acc2.ID, toAccount.ID)

		diff1 := acc1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - acc2.Balance
		require.Equal(t, diff1, diff2)
		require.Positive(t, diff1)
		require.True(t, diff1%amount == 0)
		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)

		require.NotContains(t, existed, k)
		existed[k] = true
	}

	account1, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account1)
	require.Equal(t, account1.Balance, acc1.Balance-amount*int64(n))

	account2, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Balance, acc2.Balance+amount*int64(n))

}
