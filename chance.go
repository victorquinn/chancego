package chance

import (
	"math/rand"
	"strings"
)

// Char will return a random character as a string. Note this is by default
// using a simple English alphabet, i18n being punted on at this time. It also
// intentionally returns a string rather than a rune for ease of use in the
// primare use case this is intended to solve, namely rapid testing
func Char() string {
	return CharFromPool("abcdefghijklmnopqrstuvwxyz")
}

// CharFromPool takes a string containing all the characters to add to the pool
// and then selects one of them randomly
func CharFromPool(p string) string {
	pool := strings.Split(p, "")
	if len(pool) < 1 {
		panic("Cannot call CharFromPool with pool less than 1 character")
	}
	i := rand.Int()
	return pool[i % len(pool)]
}
