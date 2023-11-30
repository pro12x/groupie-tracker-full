package pkg

import (
	"math/rand"
	"time"
)

func RandomID(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
