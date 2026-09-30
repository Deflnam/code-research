package coder

import (
	"hash/crc32"
)

func Encode(data []byte) []byte {
	crc := crc32.ChecksumIEEE(data)
	out := make([]byte, len(data)+4)
	copy(out, data)

	out[len(data)] = byte(crc >> 24)
	out[len(data)+1] = byte(crc >> 16)
	out[len(data)+2] = byte(crc >> 8)
	out[len(data)+3] = byte(crc)

	return out
}

func Detect(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	n := len(data) - 4
	expected := uint32(data[n])<<24 |
		uint32(data[n+1])<<16 |
		uint32(data[n+2])<<8 |
		uint32(data[n+3])

	crc := crc32.ChecksumIEEE(data[:n])
	return crc == expected
}
