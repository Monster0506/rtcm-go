package rtcm

import "testing"



func TestDecodeMsg1037(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x11, 0x40, 0xd0, 0x31, 0x15, 0x18, 0x08, 0x32, 0x03, 0x29,
		0x0a, 0xaf, 0xd1, 0x21, 0x9d, 0x01, 0x90, 0x00, 0x85, 0x81, 0x15,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1037(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1037: unexpected error: %v", err)
	}
	if m.MessageType != 1037 || m.NetworkID != 3 ||
		m.SubnetworkID != 1 || m.MultipleMessageIndicator != true ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 || m.NumGLONASSDataEntries != 2 {
		t.Fatalf("header mismatch: %+v", m.GLONASSNetworkRTKHeader)
	}
	checkFloat(t, "GLONASSEpochTimeS", m.GLONASSEpochTimeS, 8640.0, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].AmbiguityStatusFlag != 1 || m.Satellites[0].NonSyncCount != 2 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	checkFloat(t, "Satellites[0].IonosphericCorrectionDiffM", m.Satellites[0].IonosphericCorrectionDiffM, -0.75, 1e-9)
	if m.Satellites[1].SatelliteID != 12 || m.Satellites[1].AmbiguityStatusFlag != 3 || m.Satellites[1].NonSyncCount != 5 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	checkFloat(t, "Satellites[1].IonosphericCorrectionDiffM", m.Satellites[1].IonosphericCorrectionDiffM, 0.4, 1e-9)
}
