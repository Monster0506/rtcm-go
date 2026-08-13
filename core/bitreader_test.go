package core

import "testing"

func TestBitReaderReadUint(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		reads    []int
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

func TestBitReaderReadBits38(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x13, 0x3e, 0xd7, 0xd3, 0x02, 0x02, 0x98, 0x0e, 0xde, 0xef,
		0x34, 0xb4, 0xbd, 0x62, 0xac, 0x09, 0x41, 0x98, 0x6f, 0x33, 0x36, 0x0b, 0x98,
	}

	tests := []struct {
		name     string
		skip     int
		expected int64
	}{
		{
			name:     "ECEF-X, positive value",
			skip:     58,
			expected: 11141045999,
		},
		{
			name:     "ECEF-Y, negative value",
			skip:     98,
			expected: -48507297108,
		},
		{
			name:     "ECEF-Z, positive value",
			skip:     138,
			expected: 39755214643,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBitReader(frame)
			remaining := tt.skip
			for remaining > 0 {
				chunk := remaining
				if chunk > 64 {
					chunk = 64
				}
				r.ReadUint(chunk)
				remaining -= chunk
			}
			got := r.ReadBits38()
			if got != tt.expected {
				t.Fatalf("ReadBits38() = %d, want %d", got, tt.expected)
			}
		})
	}
}

func TestBitReaderReadSignMagnitude(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		nbits    int
		expected int64
	}{
		{
			name:     "spec worked example, negative",
			data:     []byte{0b10000101},
			nbits:    8,
			expected: -5,
		},
		{
			name:     "spec worked example, positive",
			data:     []byte{0b00000101},
			nbits:    8,
			expected: 5,
		},
		{
			name:     "zero",
			data:     []byte{0x00},
			nbits:    8,
			expected: 0,
		},
		{
			name:     "negative zero collapses to zero",
			data:     []byte{0x80},
			nbits:    8,
			expected: 0,
		},
		{
			name:     "max positive magnitude, 24-bit (GLONASS DF111 width)",
			data:     []byte{0x7F, 0xFF, 0xFF},
			nbits:    24,
			expected: 1<<23 - 1,
		},
		{
			name:     "max negative magnitude, 24-bit (GLONASS DF111 width)",
			data:     []byte{0xFF, 0xFF, 0xFF},
			nbits:    24,
			expected: -(1<<23 - 1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewBitReader(tt.data)
			got := r.ReadSignMagnitude(tt.nbits)
			if got != tt.expected {
				t.Fatalf("ReadSignMagnitude(%d) = %d, want %d", tt.nbits, got, tt.expected)
			}
		})
	}
}

func TestBitReaderReadSignMagnitudeSequential(t *testing.T) {
	data := []byte{0b10000101, 0b00000101}
	r := NewBitReader(data)
	if got := r.ReadSignMagnitude(8); got != -5 {
		t.Fatalf("first read = %d, want -5", got)
	}
	if got := r.ReadSignMagnitude(8); got != 5 {
		t.Fatalf("second read = %d, want 5", got)
	}
}
