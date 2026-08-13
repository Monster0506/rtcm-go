package core

import "testing"

func TestReadLengthPrefixedString(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want string
	}{
		{
			name: "zero-length string",
			data: []byte{0x00},
			want: "",
		},
		{
			name: "real 1008 antenna descriptor",
			data: []byte{
				0x14, 'S', 'E', 'P', 'C', 'H', 'O', 'K', 'E', '_', 'B',
				'3', 'E', '6', ' ', ' ', ' ', 'S', 'P', 'K', 'E',
			},
			want: "SEPCHOKE_B3E6   SPKE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReadLengthPrefixedString(NewBitReader(tt.data))
			if got != tt.want {
				t.Fatalf("ReadLengthPrefixedString() = %q, want %q", got, tt.want)
			}
		})
	}
}
