package rtcm

import "testing"



func TestDecodeMsg1016(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x13, 0x3f, 0x80, 0x31, 0x09, 0xb4, 0x61, 0x06, 0x40, 0x65,
		0x21, 0x55, 0xfa, 0x24, 0x21, 0x33, 0xa0, 0x32, 0x0c, 0x80, 0x88, 0xda,
		0xf4,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1016(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1016: unexpected error: %v", err)
	}
	if m.MessageType != 1016 || m.NetworkID != 3 ||
		m.SubnetworkID != 1 || m.MultipleMessageIndicator != true ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 || m.NumGPSSats != 2 {
		t.Fatalf("header mismatch: %+v", m.NetworkRTKHeader)
	}
	checkFloat(t, "GPSEpochTimeS", m.GPSEpochTimeS, 31800.0, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].AmbiguityStatusFlag != 1 || m.Satellites[0].NonSyncCount != 2 || m.Satellites[0].IODE != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	checkFloat(t, "Satellites[0].GeometricCorrectionDiffM", m.Satellites[0].GeometricCorrectionDiffM, -0.75, 1e-9)
	if m.Satellites[1].SatelliteID != 12 || m.Satellites[1].AmbiguityStatusFlag != 3 || m.Satellites[1].NonSyncCount != 5 || m.Satellites[1].IODE != 200 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	checkFloat(t, "Satellites[1].GeometricCorrectionDiffM", m.Satellites[1].GeometricCorrectionDiffM, 0.4, 1e-9)
}
