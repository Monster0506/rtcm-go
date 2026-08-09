package rtcm

import "testing"

func TestBitReaderReadUint(t *testing.T) {
	r := NewBitReader([]byte{0xB4, 0xE0})
	got := r.ReadUint(12)
	want := uint64(0xB4E)
	if got != want {
		t.Fatalf("ReadUint(12) = %#x, want %#x", got, want)
	}
}
