package token

import (
	"testing"
	"time"

	"github.com/eugenius-watchman/golang_simplebank/util"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
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

func TestExpiredJWTToken(t *testing.T) {
	// Create new JWT maker with random secret key
	maker, err := NewJWTMaker(util.RandomString(32))
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

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)

}

// This test verifies our system rejects tokens using the "none" algorithm
// which is a security risk if accepted
// func TestInvalidJWTTokenAlgNone(t *testing.T) {
//     // 1. Create a valid payload (token content) with random username
//     //    that would normally last for 1 minute
//     payload, err := NewPayload(util.RandomOwner(), time.Minute)
//     require.NoError(t, err) // Ensure no error creating payload

//     // 2. Create a UNSAFE token using the "none" signing algorithm
//     //    - jwt.NewWithClaims creates a new token structure
//     //    - jwt.SigningMethodNone means no cryptographic signature
//     //    This is dangerous because anyone could modify these tokens!
//     jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)

//     // 3. Convert the token to a string format
//     //    - jwt.UnsafeAllowNoneSignatureType is required to create these tokens
//     //    - Normally you should NEVER use this in production!
//     token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
//     require.NoError(t, err) // Ensure token string was created

//     // 4. Create our normal JWT maker that uses proper signing
//     maker, err := NewJWTMaker(util.RandomString(32))
//     require.NoError(t, err) // Ensure maker was created

//     // 5. Try to verify the unsafe "none" algorithm token
//     //    This should FAIL because we only accept properly signed tokens
//     payload, err = maker.VerifyToken(token)

//     // 6. Verify the results:
//     require.Error(t, err) // We expect an error
//     require.EqualError(t, err, ErrInvalidToken.Error()) // Specifically "invalid token" error
//     require.Nil(t, payload) // Should get no payload back
// }

// func TestInvalidJWTToken(t *testing.T) {
//     // Create a valid token
//     maker, err := NewJWTMaker(util.RandomString(32))
//     require.NoError(t, err)

//     token, err := maker.CreateToken(util.RandomOwner(), time.Minute)
//     require.NoError(t, err)
//     require.NotEmpty(t, token)

//     // Test cases for different invalid scenarios
//     testCases := []struct {
//         name    string
//         token   string
//         check   func(t *testing.T, payload *Payload, err error)
//     }{
//         {
//             name: "EmptyToken",
//             token: "",
//             check: func(t *testing.T, payload *Payload, err error) {
//                 require.Error(t, err)
//                 require.EqualError(t, err, ErrInvalidToken.Error())
//                 require.Nil(t, payload)
//             },
//         },
//         {
//             name: "RandomString",
//             token: util.RandomString(32),
//             check: func(t *testing.T, payload *Payload, err error) {
//                 require.Error(t, err)
//                 require.EqualError(t, err, ErrInvalidToken.Error())
//                 require.Nil(t, payload)
//             },
//         },
//         {
//             name: "TamperedToken",
//             token: func() string {
//                 // Take a valid token and change one character
//                 return token[:len(token)-1] + "x"
//             }(),
//             check: func(t *testing.T, payload *Payload, err error) {
//                 require.Error(t, err)
//                 require.EqualError(t, err, ErrInvalidToken.Error())
//                 require.Nil(t, payload)
//             },
//         },
//     }

//     for _, tc := range testCases {
//         t.Run(tc.name, func(t *testing.T) {
//             payload, err := maker.VerifyToken(tc.token)
//             tc.check(t, payload, err)
//         })
//     }
// }
