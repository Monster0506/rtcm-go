package rtcm

import "testing"

func TestDecodeMsg1008(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x1e, 0x3f, 0x00, 0x00, 0x14, 'S', 'E', 'P', 'C', 'H',
		'O', 'K', 'E', '_', 'B', '3', 'E', '6', ' ', ' ', ' ', 'S', 'P',
		'K', 'E', 0x00, 0x04, '5', '8', '5', '6', 0xff, 0x68, 0x94,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1008(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1008: unexpected error: %v", err)
	}
	if msg.MessageType != 1008 {
		t.Fatalf("MessageType = %d, want 1008", msg.MessageType)
	}
	if msg.AntennaDescriptor != "SEPCHOKE_B3E6   SPKE" {
		t.Fatalf("AntennaDescriptor = %q, want %q", msg.AntennaDescriptor, "SEPCHOKE_B3E6   SPKE")
	}
	if msg.SetupID != 0 {
		t.Fatalf("SetupID = %d, want 0", msg.SetupID)
	}
	if msg.AntennaSerial != "5856" {
		t.Fatalf("AntennaSerial = %q, want %q", msg.AntennaSerial, "5856")
	}
}
