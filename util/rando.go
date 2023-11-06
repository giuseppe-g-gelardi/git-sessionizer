package util

import (
	"math/rand"
)

func Rando(max ...int) int {
	if len(max) == 0 {
		return 3 // Default value is 3
	}
	return rand.Intn(max[0])
}
