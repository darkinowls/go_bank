package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

var Currencies = []string{USD, EUR, CAD}

func IsValidCurrency(currency string) bool {
	for _, c := range Currencies {
		if c == currency {
			return true
		}
	}
	return false
}
