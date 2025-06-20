package util

import (
	"math/rand"
	"strings"
	"time"
)

// lowercase letters for random string generation
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt to generate a random integer between min and max ...inclusive
func RandomInt(min, max int64) int64 {
	// Create new random generator seeded with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Calculate random number in a specified range
	return min + r.Int63n(max-min+1)
}

// RandomString to generate a random string of length n
func RandomString(n int) string {
	// Create new random generator seeded with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var sb strings.Builder //  Efficient string builder
	k := len(alphabet) // length of alphabet

	// Building the string xter by xter
	for i := 0; i < n; i++ {
		// Pick random letter from alphabet
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c) //Add letter to  string
	}
	return sb.String() // Return completed string

}
// RandomOwner generates a random 6-character owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random money amount between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency picks a random currency from supported options
func RandomCurrency() string {
	currencies := []string{"GHS", "USD", "EUR"}
	// Create a new random generator seeded with current time
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Return a random currency from the slice
	return currencies[r.Intn(len(currencies))]
}



// package util

// import (
// 	"math/rand"
// 	"strings"
// 	"time"
// )

// // lowercase letters for random string generatin
// const alphabet = "abcdefghijklmnopqrstuvwxyz"


// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// // RandomInt to generate random number between min and max
// func RandomInt(min, max int64) int64 {
// 	return min + rand.Int63n(max - min + 1)
// }

// // RandomString to generate a random string of length n
// func RandomString(n int) string {
// 	var sb strings.Builder
// 	k := len(alphabet)

// 	for i := 0; i < n; i++ {
// 		c:= alphabet[rand.Intn(k)]
// 		sb.WriteByte(c)
// 	}

// 	return sb.String()
// }

// // Define RandomOwner...generates a random owner name
// func RandomOwner() string {
// 	return RandomString(6)
// }

// // RandomMoney generator...generate random amount of money
// func RandomMoney() int64 {
// 	return RandomInt(0, 1000)
// }

// // Random currency... generate random currency code
// func RandomCurrency() string {
// 	currencies := []string{"GHS","USD","EUR"}
// 	n := len(currencies)
// 	return currencies[rand.Intn(n)]
// }