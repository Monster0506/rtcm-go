package station

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
)

func TestDecodeMsg1033(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x39, 0x40, 0x90, 0x00, 0x14, 0x53, 0x45, 0x50, 0x43,
		0x48, 0x4f, 0x4b, 0x45, 0x5f, 0x42, 0x33, 0x45, 0x36, 0x20, 0x20,
		0x20, 0x53, 0x50, 0x4b, 0x45, 0x00, 0x04, 0x35, 0x38, 0x35, 0x36,
		0x0c, 0x53, 0x45, 0x50, 0x54, 0x20, 0x50, 0x4f, 0x4c, 0x41, 0x52,
		0x58, 0x35, 0x05, 0x35, 0x2e, 0x35, 0x2e, 0x30, 0x07, 0x33, 0x30,
		0x37, 0x35, 0x30, 0x32, 0x34, 0xb9, 0x1f, 0xbd,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1033(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1033: unexpected error: %v", err)
	}
	if msg.MessageType != 1033 {
		t.Fatalf("MessageType = %d, want 1033", msg.MessageType)
	}
	if msg.AntennaDescriptor != "SEPCHOKE_B3E6   SPKE" {
		t.Fatalf("AntennaDescriptor = %q", msg.AntennaDescriptor)
	}
	if msg.AntennaSerial != "5856" {
		t.Fatalf("AntennaSerial = %q", msg.AntennaSerial)
	}
	if msg.ReceiverType != "SEPT POLARX5" {
		t.Fatalf("ReceiverType = %q, want %q", msg.ReceiverType, "SEPT POLARX5")
	}
	if msg.FirmwareVersion != "5.5.0" {
		t.Fatalf("FirmwareVersion = %q, want %q", msg.FirmwareVersion, "5.5.0")
	}
	if msg.ReceiverSerial != "3075024" {
		t.Fatalf("ReceiverSerial = %q, want %q", msg.ReceiverSerial, "3075024")
	}
}
