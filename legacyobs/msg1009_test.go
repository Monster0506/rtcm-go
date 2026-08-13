package legacyobs

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1009(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x48, 0x3f, 0x10, 0x00, 0x86, 0x85, 0x03, 0x14, 0x00, 0x24,
		0x34, 0x07, 0xba, 0x41, 0x74, 0x0b, 0xfa, 0xc2, 0x20, 0x33, 0x11, 0x3d,
		0xa0, 0x2b, 0xfb, 0x04, 0xb8, 0x71, 0x80, 0xc0, 0xbd, 0xdf, 0xf9, 0x06,
		0xb8, 0xd7, 0x55, 0x80, 0x75, 0xbb, 0xf8, 0xe6, 0x20, 0xc5, 0x18, 0x3d,
		0x23, 0xb7, 0xfa, 0xe5, 0x5c, 0x85, 0xd8, 0x80, 0x5c, 0x83, 0xf9, 0x40,
		0x6d, 0x8d, 0x25, 0x01, 0xb1, 0x37, 0xf9, 0x22, 0xbd, 0x7a, 0xc2, 0x81,
		0x68, 0x03, 0xf8, 0x9f, 0xc6, 0xe6,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1009(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1009: unexpected error: %v", err)
	}
	if msg.MessageType != 1009 {
		t.Fatalf("MessageType = %d, want 1009", msg.MessageType)
	}
	if msg.StationID != 0 {
		t.Fatalf("StationID = %d, want 0", msg.StationID)
	}
	if msg.GLONASSEpochTkMs != 70527000 {
		t.Fatalf("GLONASSEpochTkMs = %d, want 70527000", msg.GLONASSEpochTkMs)
	}
	if len(msg.Satellites) != 8 {
		t.Fatalf("len(Satellites) = %d, want 8", len(msg.Satellites))
	}
	s0 := msg.Satellites[0]
	if s0.SatelliteID != 1 {
		t.Fatalf("Satellites[0].SatelliteID = %d, want 1", s0.SatelliteID)
	}
	if s0.FrequencyChannelNumber != 8 {
		t.Fatalf("Satellites[0].FrequencyChannelNumber = %d, want 8", s0.FrequencyChannelNumber)
	}
	testutil.CheckFloat(t, "Satellites[0].L1PseudorangeM", s0.L1PseudorangeM, 272788.02, 1e-6)
	testutil.CheckFloat(t, "Satellites[0].L1PhaserangeM", s0.L1PhaserangeM, 11.905, 1e-6)
	if s0.L1LockTimeIndicator != 127 {
		t.Fatalf("Satellites[0].L1LockTimeIndicator = %d, want 127", s0.L1LockTimeIndicator)
	}

}
