package legacyobs

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1011(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x73, 0x3f, 0x30, 0x00, 0x86, 0x85, 0x03, 0x14, 0x00, 0x24,
		0x34, 0x07, 0xba, 0x41, 0x74, 0x0b, 0xf8, 0x17, 0x88, 0x4b, 0x56, 0xe9,
		0x58, 0x44, 0x06, 0x62, 0x27, 0xb4, 0x05, 0x7f, 0x02, 0x59, 0xf7, 0xc6,
		0x8f, 0xec, 0x12, 0xe1, 0xc6, 0x03, 0x02, 0xf7, 0x7f, 0xe0, 0x58, 0x01,
		0xaf, 0x2f, 0xfc, 0x83, 0x5c, 0x6b, 0xaa, 0xc0, 0x3a, 0xdd, 0xfc, 0x07,
		0x68, 0x1f, 0x60, 0xbf, 0x8e, 0x62, 0x0c, 0x51, 0x83, 0xd2, 0x3b, 0x7f,
		0x81, 0x5c, 0x7a, 0x8f, 0x77, 0xf5, 0xca, 0xb9, 0x0b, 0xb1, 0x00, 0xb9,
		0x07, 0xf2, 0x00, 0x08, 0x00, 0x00, 0x00, 0x50, 0x1b, 0x63, 0x49, 0x40,
		0x6c, 0x4d, 0xfe, 0x40, 0x01, 0x00, 0x00, 0x00, 0x09, 0x15, 0xeb, 0xd6,
		0x14, 0x0b, 0x40, 0x1f, 0xc0, 0xe2, 0x85, 0x39, 0x13, 0xf8, 0x27, 0xe0,
		0x58,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1011(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1011: unexpected error: %v", err)
	}
	if msg.MessageType != 1011 {
		t.Fatalf("MessageType = %d, want 1011", msg.MessageType)
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
	if s0.L2CodeIndicator != 0 {
		t.Fatalf("Satellites[0].L2CodeIndicator = %d, want 0", s0.L2CodeIndicator)
	}
	testutil.CheckFloat(t, "Satellites[0].L2MinusL1PseudorangeM", s0.L2MinusL1PseudorangeM, 15.06, 1e-6)
	testutil.CheckFloat(t, "Satellites[0].L2PhaserangeM", s0.L2PhaserangeM, 19.2865, 1e-6)
	if s0.L2LockTimeIndicator != 105 {
		t.Fatalf("Satellites[0].L2LockTimeIndicator = %d, want 105", s0.L2LockTimeIndicator)
	}
}
