package rtcm

import "testing"

func TestBitReaderReadUint(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		reads []int
		expected []uint64
	}{
		{
			name:     "byte-crossing 12-bit read",
			data:     []byte{0xB4, 0xE0},
			reads:    []int{12},
			expected: []uint64{0xB4E},
		},
		{
			name:     "zero-width read does not advance position",
			data:     []byte{0xFF},
			reads:    []int{0, 8},
			expected: []uint64{0, 0xFF},
		},
		{
			name:     "full 64-bit width, byte-aligned",
			data:     []byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF},
			reads:    []int{64},
			expected: []uint64{0x0123456789ABCDEF},
		},
		{
			name:     "sequential reads, second starts at non-byte-aligned offset",
			data:     []byte{0xB4, 0xE0, 0x3C},
			reads:    []int{12, 12},
			expected: []uint64{0xB4E, 0x03C},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBitReader(tt.data)
			for i, nbits := range tt.reads {
				got := r.ReadUint(nbits)
				want := tt.expected[i]
				if got != want {
					t.Fatalf("read %d: ReadUint(%d) = %#x, want %#x", i, nbits, got, want)
				}
			}
		})
	}
}

func TestBitReaderReadInt(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		reads    []int
		expected []int64
	}{
		{
			name:     "negative value, sign bit set",
			data:     []byte{0xFF},
			reads:    []int{8},
			expected: []int64{-1},
		},
		{
			name:     "positive value, sign bit clear",
			data:     []byte{0x7F},
			reads:    []int{8},
			expected: []int64{127},
		},
		{
			name:     "full 64-bit width, all-ones pattern",
			data:     []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF},
			reads:    []int{64},
			expected: []int64{-1},
		},
		{
			name:     "full 64-bit width, min int64 pattern",
			data:     []byte{0x80, 0, 0, 0, 0, 0, 0, 0},
			reads:    []int{64},
			expected: []int64{-9223372036854775808},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBitReader(tt.data)
			for i, nbits := range tt.reads {
				got := r.ReadInt(nbits)
				want := tt.expected[i]
				if got != want {
					t.Fatalf("read %d: ReadInt(%d) = %d, want %d", i, nbits, got, want)
				}
			}
		})
	}
}
