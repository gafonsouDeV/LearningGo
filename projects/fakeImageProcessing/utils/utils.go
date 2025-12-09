package utils

import (
	"math/rand"
	"time"
)

func RandDelay() time.Duration {
	min := 200  // milisegundos
	max := 1200 // milisegundos
	ms := rand.Intn(max-min) + min
	return time.Duration(ms) * time.Millisecond
}
