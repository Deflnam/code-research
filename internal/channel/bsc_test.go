package channel

import "testing"

func TestNoiseZero(t *testing.T) {
	data := []byte{1, 2, 3}
	original := append([]byte{}, data...)

	result := AddNoise(data, 0)

	for i := range result {
		if result[i] != original[i] {
			t.Errorf("at index %d: expected %d, got %d", i, original[i], result[i])
		}
	}

}
