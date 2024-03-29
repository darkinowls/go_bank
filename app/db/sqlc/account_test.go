package sqlc

import (
	"app/util"
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, arg.Owner, acc.Owner)
	require.Equal(t, arg.Balance, acc.Balance)
	require.Equal(t, arg.Currency, acc.Currency)

	require.Positive(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)
	return acc

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// independent test
	acc1 := createRandomAccount(t)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      acc1.ID,
		Balance: acc1.Balance + 1,
	}
	acc2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc2)

	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, arg.Balance, acc2.Balance)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err)
	require.Empty(t, acc2)
}

func TestListAccounts(t *testing.T) {
	createRandomAccount(t)
	createRandomAccount(t)

	arg := ListAccountsParams{
		Limit:  2,
		Offset: 0,
	}
	accs, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accs, 2)

	for _, acc := range accs {
		require.NotEmpty(t, acc)
	}
}
