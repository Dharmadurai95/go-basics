package helpers

import (
	"math/rand"
	"time"
)

// GenerateRandomNum generates a random number up to max
func GenerateRandomNum(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}
