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
