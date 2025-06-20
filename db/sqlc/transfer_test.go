package db

import (
	"context"
	"testing"
	"time"

	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/stretchr/testify/require"
)

// TestCreateTransfer tests creating a transfer between accounts
func TestCreateTransfer(t *testing.T) {
	// Create 2 random accounts to transfer between
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Create the transfer arguments
	arg := CreateTransferParams{ 
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:       util.RandomMoney(),
	}

	// Create the transfer
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err) // Should not return error
	require.NotEmpty(t, transfer) // Should return transfer object

	// Verify all fields are correct
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	// Verify auto-generated fields
	require.NotZero(t, transfer.ID) // ID must be assigned
	require.NotZero(t, transfer.CreatedAt) // Timestamp should be set
}

// TestGetTransfer tests retrieving a transfer by ID
func TestGetTransfer(t *testing.T) {
	// Create a transfer to test with
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1, err := testQueries.CreateTransfer(context.Background(), // Fixed: typo in "context"
		CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        util.RandomMoney(),
		})
	require.NoError(t, err)

	// Get the transfer that was just created
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID) // Fixed: Changed TestCreateTransfer1 to transfer1

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	// Verify all fields match the original transfer
	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

// TestListTransfers tests list of transfers
func TestListTransfers(t *testing.T) { // Fixed: Typo in function name
	// Create 2 accounts to use for transfers
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// Create 5 random transfers between accounts
	for i := 0; i < 5; i++ { // Fixed: Space after "for"
		testQueries.CreateTransfer(context.Background(), 
			CreateTransferParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        util.RandomMoney(),
			})
	}
	// Test listing transfers with pagination ...limit 3, offset/skip 1
	arg := ListTransfersParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:  3,
		Offset: 1,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 3) // Should get exactly 3 transfers

	// Verify all returned transfers are valid
	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
	}
}