package rtcm

import "testing"



func TestDecodeMsg1057(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x1a, 0x42, 0x14, 0xda, 0x30, 0x5c, 0xc3, 0xe8, 0x48, 0x11,
		0x48, 0x7e, 0x79, 0x60, 0x30, 0xd4, 0x0b, 0x6c, 0x20, 0x00, 0x7d, 0x07,
		0xec, 0x78, 0x02, 0xee, 0x00, 0x00, 0x88, 0x7f,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1057(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1057: unexpected error: %v", err)
	}
	if m.MessageType != 1057 || m.EpochTimeS != 318000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.SatelliteReferenceDatum != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSROrbitHeader)
	}
	checkFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.Corrections) != 1 {
		t.Fatalf("len(Corrections)")
	}
	if m.Corrections[0].SatelliteID != 5 || m.Corrections[0].IOD != 33 {
		t.Fatalf("orbit correction mismatch")
	}
	checkFloat(t, "DeltaRadialM", m.Corrections[0].DeltaRadialM, -10.0, 1e-9)
	checkFloat(t, "DeltaAlongTrackM", m.Corrections[0].DeltaAlongTrackM, 80.0, 1e-9)
	checkFloat(t, "DeltaCrossTrackM", m.Corrections[0].DeltaCrossTrackM, -120.0, 1e-9)
	checkFloat(t, "DotDeltaRadialMPerS", m.Corrections[0].DotDeltaRadialMPerS, 0.004, 1e-9)
	checkFloat(t, "DotDeltaAlongTrackMPerS", m.Corrections[0].DotDeltaAlongTrackMPerS, -0.02, 1e-9)
	checkFloat(t, "DotDeltaCrossTrackMPerS", m.Corrections[0].DotDeltaCrossTrackMPerS, 0.024, 1e-9)
}
