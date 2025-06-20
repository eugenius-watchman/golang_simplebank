package db

import (
	"context"
	"testing"
	"time"

	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/stretchr/testify/require"
)

// TestCreateEntry tests creating an account entry (deposit/withdrawal)
func TestCreateEntry(t *testing.T) {
	// First create a random account to make entries for
	account := createRandomAccount(t)

	// Create entry arguments
	arg := CreateEntryParams{
		AccountID: account.ID, // Which account this entry belongs to
		Amount:    util.RandomMoney(), // Can be positive (deposit) or negative (withdrawal)
	}

	// Create the entry
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err) // Should not return error
	require.NotEmpty(t, entry) // Should return an entry object

	// Verify all fields are correct
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	
	// Verify auto-generated fields
	require.NotZero(t, entry.ID) // ID should be assigned
	require.NotZero(t, entry.CreatedAt) // Timestamp should be set
}

// TestGetEntry tests retrieving an entry by ID
func TestGetEntry(t *testing.T) {
	// First create an account and entry to test with
	account := createRandomAccount(t)
	entry1, err := testQueries.CreateEntry(context.Background(), CreateEntryParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	})
	require.NoError(t, err)

	// Now try to get the entry we just created
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	// Verify all fields match the original entry
	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

// TestListEntries tests listing entries for an account with pagination
func TestListEntries(t *testing.T) {
	// Create an account to make entries for
	account := createRandomAccount(t)

	// Create 5 random entries for this account
	for i := 0; i < 5; i++ {
		testQueries.CreateEntry(context.Background(), CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomMoney(),
		})
	}

	// Test listing entries with pagination (limit 3, skip 1)
	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:  3,
		Offset: 1,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 3) // Should get exactly 3 entries

	// Verify all returned entries belong to our account
	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, account.ID, entry.AccountID)
	}
}