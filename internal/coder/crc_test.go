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

func TestDetectCorrupted(t *testing.T) {
	data := []byte{1, 2, 3, 4}
	encoded := Encode(data)
	encoded[0] = 67

	if Detect(encoded) {
		t.Errorf("expected detected=false, got true")
	}
}

func TestDetectShort(t *testing.T) {
	data := []byte{1, 2}

	if Detect(data) {
		t.Errorf("expected detected=false, got true")
	}
}
