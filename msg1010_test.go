package rtcm

import "testing"

func TestDecodeMsg1010(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x57, 0x3f, 0x20, 0x00, 0x86, 0x85, 0x03, 0x14, 0x00, 0x24,
		0x34, 0x07, 0xba, 0x41, 0x74, 0x0b, 0xfa, 0x5a, 0x65, 0x84, 0x40, 0x66,
		0x22, 0x7b, 0x40, 0x57, 0xf4, 0x78, 0xec, 0x12, 0xe1, 0xc6, 0x03, 0x02,
		0xf7, 0x7f, 0xe9, 0x6d, 0xc8, 0x35, 0xc6, 0xba, 0xac, 0x03, 0xad, 0xdf,
		0xd0, 0xe9, 0x8e, 0x62, 0x0c, 0x51, 0x83, 0xd2, 0x3b, 0x7f, 0xa4, 0xc1,
		0x5c, 0xab, 0x90, 0xbb, 0x10, 0x0b, 0x90, 0x7f, 0x41, 0x96, 0x50, 0x1b,
		0x63, 0x49, 0x40, 0x6c, 0x4d, 0xfe, 0x96, 0x70, 0x91, 0x5e, 0xbd, 0x61,
		0x40, 0xb4, 0x01, 0xfd, 0x2d, 0x90, 0xbb, 0x83, 0x54,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1010(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1010: unexpected error: %v", err)
	}
	if msg.MessageType != 1010 {
		t.Fatalf("MessageType = %d, want 1010", msg.MessageType)
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
	checkFloat(t, "Satellites[0].L1PseudorangeM", s0.L1PseudorangeM, 272788.02, 1e-6)
	checkFloat(t, "Satellites[0].L1PhaserangeM", s0.L1PhaserangeM, 11.905, 1e-6)
	if s0.L1LockTimeIndicator != 127 {
		t.Fatalf("Satellites[0].L1LockTimeIndicator = %d, want 127", s0.L1LockTimeIndicator)
	}
	checkFloat(t, "Satellites[0].L1PseudorangeAmbiguityM", s0.L1PseudorangeAmbiguityM, 22184641.892, 1e-2)
	checkFloat(t, "Satellites[0].L1CNRDbHz", s0.L1CNRDbHz, 41.5, 1e-6)
}
