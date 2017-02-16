package chance

import (
	"math/rand"
	"strings"
)

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
