package rtcm

import "testing"



func TestDecodeMsg1063(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x19, 0x42, 0x77, 0x53, 0x02, 0xe6, 0x1f, 0x42, 0x40, 0x94,
		0x87, 0xe7, 0x96, 0x03, 0x0d, 0x40, 0xb6, 0xc2, 0x00, 0x07, 0xd0, 0x7e,
		0xc7, 0x80, 0x2e, 0xe0, 0x15, 0x3f, 0x0e,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1063(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1063: unexpected error: %v", err)
	}
	if m.MessageType != 1063 || m.EpochTimeS != 60000 ||
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
