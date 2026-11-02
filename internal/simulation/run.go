package simulation

import (
	"math/rand"

	"github.com/Deflnam/code-research/internal/channel"
	"github.com/Deflnam/code-research/internal/coder"
)

type Config struct {
	P         float64
	Length    int
	MaxErrors int
	MaxBlocks int
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
	errorsCount := 0

	for result.Total < cfg.MaxBlocks {
		data := make([]byte, cfg.Length)
		rand.Read(data)

		encoded := coder.Encode(data)
		received := channel.AddNoise(encoded, cfg.P)
		detected := coder.Detect(received)

		hasErrors := !sliceEqual(received, encoded)

		result.Total++

		if hasErrors {
			errorsCount++
			if detected {
				result.Detected++
			} else {
				result.Missed++
			}
		}

		if errorsCount >= cfg.MaxErrors {
			break
		}
	}

	return result
}
