package rtcm

import "testing"

func TestCRC24Q(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected uint32
	}{
		{
			name:     "full message including trailer is zero",
			data:     []byte{0xd3, 0x00, 0x04, 0x4c, 0xe0, 0x00, 0x80, 0xed, 0xed, 0xd6},
			expected: 0,
		},
		{
			name:     "header-only bytes produce the trailer",
			data:     []byte{0xd3, 0x00, 0x04, 0x4c, 0xe0, 0x00, 0x80},
			expected: 0xEDEDD6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CRC24Q(tt.data)
			if got != tt.expected {
				t.Fatalf("CRC24Q() = %#x, want %#x", got, tt.expected)
			}
		})
	}
}
