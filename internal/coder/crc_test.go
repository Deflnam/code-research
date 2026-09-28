package coder

import "testing"

func TestEncodeDetect(t *testing.T) {
	data := []byte{1, 2, 3, 4}
	encoded := Encode(data)
	detected := Detect(encoded)

	if !detected {
		t.Errorf("expected detected=true, got false")
	}
}
