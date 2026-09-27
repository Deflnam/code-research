package channel

import (
	"math/rand"
)

func AddNoise(data []byte, p float64) []byte {
	out := make([]byte, len(data))
	copy(out, data)
	for i := range out {
		for bit := 0; bit < 8; bit++ {
			if rand.Float64() < p {
				out[i] ^= 1 << bit
			}
		}
	}
	return out
}
