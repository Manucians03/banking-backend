package util

const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
	AUD = "AUD"
	CAD = "CAD"
	CHF = "CHF"
	CNY = "CNY"
	VND = "VND"
	SGD = "SGD"
	THB = "THB"
	KRW = "KRW"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, JPY, AUD, CAD, CHF, CNY, VND, SGD, THB, KRW:
		return true
	}
	return false
}
