package legacyobs

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1012(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x8a, 0x3f, 0x40, 0x00, 0x86, 0x85, 0x03, 0x14, 0x00, 0x24,
		0x34, 0x07, 0xba, 0x41, 0x74, 0x0b, 0xfa, 0x5a, 0x60, 0x2f, 0x10, 0x96,
		0xad, 0xd3, 0x1c, 0xb0, 0x88, 0x0c, 0xc4, 0x4f, 0x68, 0x0a, 0xfe, 0x8f,
		0x1c, 0x09, 0x67, 0xdf, 0x1a, 0x3f, 0xdc, 0xb0, 0x4b, 0x87, 0x18, 0x0c,
		0x0b, 0xdd, 0xff, 0xa5, 0xb7, 0x02, 0xc0, 0x0d, 0x79, 0x7f, 0xf5, 0x64,
		0x1a, 0xe3, 0x5d, 0x56, 0x01, 0xd6, 0xef, 0xe8, 0x74, 0xc0, 0x76, 0x81,
		0xf6, 0x0b, 0xfe, 0x30, 0xe6, 0x20, 0xc5, 0x18, 0x3d, 0x23, 0xb7, 0xfa,
		0x4c, 0x10, 0x2b, 0x8f, 0x51, 0xee, 0xff, 0x66, 0xb9, 0x57, 0x21, 0x76,
		0x20, 0x17, 0x20, 0xfe, 0x83, 0x2c, 0x80, 0x02, 0x00, 0x00, 0x00, 0x00,
		0x14, 0x06, 0xd8, 0xd2, 0x50, 0x1b, 0x13, 0x7f, 0xa5, 0x9c, 0x20, 0x00,
		0x80, 0x00, 0x00, 0x00, 0x04, 0x8a, 0xf5, 0xeb, 0x0a, 0x05, 0xa0, 0x0f,
		0xe9, 0x6c, 0x80, 0xe2, 0x85, 0x39, 0x13, 0xfd, 0x38, 0x0e, 0x8f, 0xad,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1012(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1012: unexpected error: %v", err)
	}
	if msg.MessageType != 1012 {
		t.Fatalf("MessageType = %d, want 1012", msg.MessageType)
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
	testutil.CheckFloat(t, "Satellites[0].L1PseudorangeAmbiguityM", s0.L1PseudorangeAmbiguityM, 22184641.892, 1e-2)
	testutil.CheckFloat(t, "Satellites[0].L1CNRDbHz", s0.L1CNRDbHz, 41.5, 1e-6)
	if s0.L2CodeIndicator != 0 {
		t.Fatalf("Satellites[0].L2CodeIndicator = %d, want 0", s0.L2CodeIndicator)
	}
	testutil.CheckFloat(t, "Satellites[0].L2MinusL1PseudorangeM", s0.L2MinusL1PseudorangeM, 15.06, 1e-6)
	testutil.CheckFloat(t, "Satellites[0].L2PhaserangeM", s0.L2PhaserangeM, 19.2865, 1e-6)
	if s0.L2LockTimeIndicator != 105 {
		t.Fatalf("Satellites[0].L2LockTimeIndicator = %d, want 105", s0.L2LockTimeIndicator)
	}
	testutil.CheckFloat(t, "Satellites[0].L2CNRDbHz", s0.L2CNRDbHz, 35.5, 1e-6)
}
