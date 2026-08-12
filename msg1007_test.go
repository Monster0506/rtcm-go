package rtcm

import "testing"



func TestDecodeMsg1007(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x10, 0x3e, 0xf0, 0x2a, 0x0b, 0x53, 0x45, 0x50, 0x43, 0x48,
		0x4f, 0x4b, 0x45, 0x5f, 0x4d, 0x43, 0x07, 0x97, 0xca, 0xff,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1007(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1007: unexpected error: %v", err)
	}
	if m.MessageType != 1007 || m.StationID != 42 ||
		m.AntennaDescriptor != "SEPCHOKE_MC" || m.SetupID != 7 {
		t.Fatalf("mismatch: %+v", m)
	}
}
