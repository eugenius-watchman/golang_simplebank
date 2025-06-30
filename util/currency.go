package util

// implement logic whether currency is supported or not 
// constants for all supported currencies
const (
	GHS = "GHS"
	EUR = "EUR"
	USD = "USD"
)

// IsSupportedCurrency returns true if currency is supported 
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case GHS, EUR, USD:
		return true
	}
	return false
}