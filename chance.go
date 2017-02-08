package chance

import (
	"math/rand"
	"strings"
)

// @todo make it so this can accept a pool and override the default
func Char() string {
	pool := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	i := rand.Int()
	return pool[i % 26]
}
