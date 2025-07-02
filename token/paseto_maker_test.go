package token

import (
	"testing"
	"time"

	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoToken(t *testing.T) {
	// Create new JWT maker with random secret key
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// Create token that expires immediately (negative duration)
	username := util.RandomOwner()
	token, err := maker.CreateToken(username, -time.Minute) // Negative duration makes it expired
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// Try to verify the expired token
	payload, err := maker.VerifyToken(token)

	// We EXPECT this to fail with an error
	require.Error(t, err)                               // Should get an error
	require.EqualError(t, err, ErrExpiredToken.Error()) // Specific expired token error
	require.Nil(t, payload)                             // Shouldnt get any payload back
}

// TestInvalidPasetoToken verifies that the token system properly rejects
// invalid or tampered-with PASETO tokens
func TestInvalidPasetoToken(t *testing.T) {
    //  create a valid token maker with random secret key
    maker, err := NewPasetoMaker(util.RandomString(32))
    require.NoError(t, err) // Make sure no error creating the maker

    // create valid token first then later modify to make invalid
    username := util.RandomOwner()
    validToken, err := maker.CreateToken(username, time.Minute)
    require.NoError(t, err)
    require.NotEmpty(t, validToken)

    // Testing different invalid token scenarios
    testCases := []struct {
        name  string       // Name of test case
        token string       // The invalid token to test
        check func(*testing.T, *Payload, error) // How to verify results
    }{
        {
            name: "EmptyToken",
            token: "",
            check: func(t *testing.T, payload *Payload, err error) {
                require.Error(t, err) // Should get an error
                require.EqualError(t, err, ErrInvalidToken.Error()) // Specific error
                require.Nil(t, payload) // No payload should be returned
            },
        },
        {
            name: "RandomString",
            token: util.RandomString(32), // Totally random string
            check: func(t *testing.T, payload *Payload, err error) {
                require.Error(t, err)
                require.EqualError(t, err, ErrInvalidToken.Error())
                require.Nil(t, payload)
            },
        },
        {
            name: "TamperedToken",
            token: func() string {
                // Take valid token and change one character
                return validToken[:len(validToken)-1] + "x"
            }(),
            check: func(t *testing.T, payload *Payload, err error) {
                require.Error(t, err)
                require.EqualError(t, err, ErrInvalidToken.Error())
                require.Nil(t, payload)
            },
        },
        {
            name: "WrongKeyToken",
            token: func() string {
                // Creating token with different key
                badMaker, _ := NewPasetoMaker(util.RandomString(32))
                token, _ := badMaker.CreateToken(username, time.Minute)
                return token
            }(),
            check: func(t *testing.T, payload *Payload, err error) {
                require.Error(t, err)
                require.EqualError(t, err, ErrInvalidToken.Error())
                require.Nil(t, payload)
            },
        },
    }

    // Run all test cases
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            payload, err := maker.VerifyToken(tc.token)
            tc.check(t, payload, err)
        })
    }
}