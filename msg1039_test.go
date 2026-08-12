package rtcm

import "testing"



func TestDecodeMsg1039(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x17, 0x40, 0xf0, 0x31, 0x15, 0x18, 0x08, 0x32, 0x03, 0x29,
		0x0a, 0xaf, 0xd1, 0x21, 0x08, 0x0e, 0x10, 0xce, 0x80, 0xc8, 0x32, 0x3f,
		0xa8, 0x80, 0xc8, 0xa6, 0xa7,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1039(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1039: unexpected error: %v", err)
	}
	if m.MessageType != 1039 || m.NetworkID != 3 ||
		m.SubnetworkID != 1 || m.MultipleMessageIndicator != true ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 || m.NumGLONASSDataEntries != 2 {
		t.Fatalf("header mismatch: %+v", m.GLONASSNetworkRTKHeader)
	}
	checkFloat(t, "GLONASSEpochTimeS", m.GLONASSEpochTimeS, 8640.0, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].AmbiguityStatusFlag != 1 || m.Satellites[0].NonSyncCount != 2 || m.Satellites[0].IOD != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	checkFloat(t, "Satellites[0].GeometricCorrectionDiffM", m.Satellites[0].GeometricCorrectionDiffM, -0.75, 1e-9)
	checkFloat(t, "Satellites[0].IonosphericCorrectionDiffM", m.Satellites[0].IonosphericCorrectionDiffM, 0.45, 1e-9)
	if m.Satellites[1].SatelliteID != 12 || m.Satellites[1].AmbiguityStatusFlag != 3 || m.Satellites[1].NonSyncCount != 5 || m.Satellites[1].IOD != 200 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	checkFloat(t, "Satellites[1].GeometricCorrectionDiffM", m.Satellites[1].GeometricCorrectionDiffM, 0.4, 1e-9)
	checkFloat(t, "Satellites[1].IonosphericCorrectionDiffM", m.Satellites[1].IonosphericCorrectionDiffM, -0.35000000000000003, 1e-9)
}
