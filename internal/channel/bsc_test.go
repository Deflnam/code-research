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

func TestNoiseOne(t *testing.T) {
	data := []byte{0, 0, 0}
	result := AddNoise(data, 1)

	for i := range result {
		if result[i] != 255 {
			t.Errorf("at index %d: expected 255, got %d", i, result[i])
		}
	}
}

func TestNoiseCopy(t *testing.T) {
	data := []byte{0, 0, 0}
	original := []byte{0, 0, 0}
	_ = AddNoise(data, 0.5)

	for i := range data {
		if data[i] != original[i] {
			t.Errorf("at index %d: expected %d, got %d", i, original[i], data[i])
		}
	}
}
