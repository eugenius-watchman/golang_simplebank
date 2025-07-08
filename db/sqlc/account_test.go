package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/stretchr/testify/require"
)

func randomAccount(owner string) Account {
	if owner == "" {
		owner = util.RandomOwner()
	}
	
	return Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
		CreatedAt: time.Now(),
	}
}

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)

	// using random account helper function
	randomAcc := randomAccount(user.Username)


	// arg := CreateAccountParams{
	// 	Owner:    user.Username,
	// 	Balance:  util.RandomMoney(),
	// 	Currency: util.RandomCurrency(),
	// }

	arg := CreateAccountParams{
		Owner:    randomAcc.Owner,
		Balance:  randomAcc.Balance,
		Currency: randomAcc.Currency,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	// Create Account
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// Test for UpdateAccount funtion
func TestUpdateAccoun(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

// Test for DeleteAccount() function
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

// Test for ListAccount() funcion
func TestListAccount(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Owner: lastAccount.Owner,
		Limit: 5,
		//Offset: 5,
		Offset: 0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	//require.Len(t, accounts, 5)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}

}
