package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
func RandomOwner() string {
	return RandomString(10)
}

func RandomAmount() int64 {
	return int64(RandomInt(0, 3000))
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD", "AUD", "JPY", "CNY", "KRW", "SGD", "VND", "THB", "CHF"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
