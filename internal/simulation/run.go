package simulation

import (
	"math/rand"

	"github.com/Deflnam/code-research/internal/channel"
	"github.com/Deflnam/code-research/internal/coder"
)

type Config struct {
	P      float64
	N      int
	Length int
}

type Result struct {
	Total    int
	Missed   int
	Detected int
}

func sliceEqual(x, y []byte) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func Run(cfg Config) Result {
	result := Result{}

	for i := 0; i < cfg.N; i++ {
		data := make([]byte, cfg.Length)
		rand.Read(data)

		encoded := coder.Encode(data)
		received := channel.AddNoise(encoded, cfg.P)
		detected := coder.Detect(received)

		hasErrors := !sliceEqual(received, encoded)

		result.Total++

		if hasErrors && !detected {
			result.Detected++
		}
		if hasErrors && detected {
			result.Missed++
		}
	}

	return result
}
