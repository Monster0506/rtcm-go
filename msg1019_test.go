package rtcm

import "testing"

func TestDecodeMsg1019(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x3d, 0x3f, 0xb0, 0x90, 0x10, 0x7a, 0xa4, 0xb9, 0x4f, 0x1a,
		0x00, 0x00, 0x36, 0xc2, 0x48, 0x58, 0xb9, 0xf1, 0x57, 0x2e, 0x09, 0x58,
		0x1c, 0x10, 0x53, 0xf3, 0xa6, 0x08, 0x40, 0xce, 0x79, 0x11, 0xf0, 0xa1,
		0x0d, 0xb5, 0xfd, 0x4f, 0x1a, 0x00, 0x82, 0x87, 0x11, 0xb6, 0xbb, 0x00,
		0x09, 0x27, 0x6e, 0xc4, 0x03, 0x1a, 0x4a, 0xce, 0x31, 0x5b, 0x86, 0xff,
		0xaa, 0xe6, 0xda, 0x00, 0x8e, 0x56, 0x2b,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1019(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1019: unexpected error: %v", err)
	}
	if m.MessageType != 1019 {
		t.Fatalf("MessageType = %d, want 1019", m.MessageType)
	}
	if m.SatelliteID != 2 {
		t.Fatalf("SatelliteID = %d, want 2", m.SatelliteID)
	}
	if m.WeekNumber != 257 {
		t.Fatalf("WeekNumber = %d, want 257", m.WeekNumber)
	}
	if m.SVAccuracy != 0 {
		t.Fatalf("SVAccuracy = %d, want 0", m.SVAccuracy)
	}
	if m.CodeOnL2 != 1 {
		t.Fatalf("CodeOnL2 = %d, want 1", m.CodeOnL2)
	}
	if m.IODE != 185 {
		t.Fatalf("IODE = %d, want 185", m.IODE)
	}
	if m.IODC != 185 {
		t.Fatalf("IODC = %d, want 185", m.IODC)
	}
	if m.SVHealth != 0 {
		t.Fatalf("SVHealth = %d, want 0", m.SVHealth)
	}
	if m.L2PDataFlag != false {
		t.Fatalf("L2PDataFlag = %v, want false", m.L2PDataFlag)
	}
	if m.FitInterval != false {
		t.Fatalf("FitInterval = %v, want false", m.FitInterval)
	}
	checkFloat(t, "IDOTSemiCirclesPerS", m.IDOTSemiCirclesPerS, -1.559783413540572e-10, 1e-9)
	checkFloat(t, "TocS", m.TocS, 324000, 1e-9)
	checkFloat(t, "Af2SPerS2", m.Af2SPerS2, 0.0, 1e-9)
	checkFloat(t, "Af1SPerS", m.Af1SPerS, 6.139089236967266e-12, 1e-9)
	checkFloat(t, "Af0S", m.Af0S, -0.00047086644917726517, 1e-9)
	checkFloat(t, "CrsM", m.CrsM, -117.28125, 1e-9)
	checkFloat(t, "DeltaNSemiCirclesPerS", m.DeltaNSemiCirclesPerS, 1.339799382549245e-09, 1e-9)
	checkFloat(t, "M0SemiCircles", m.M0SemiCircles, 0.6883564381860197, 1e-9)
	checkFloat(t, "CucRad", m.CucRad, -5.889683961868286e-06, 1e-9)
	checkFloat(t, "Eccentricity", m.Eccentricity, 0.016119434614665806, 1e-9)
	checkFloat(t, "CusRad", m.CusRad, 8.553266525268555e-06, 1e-9)
	checkFloat(t, "SqrtA", m.SqrtA, 5153.713861465454, 1e-9)
	checkFloat(t, "ToeS", m.ToeS, 324000, 1e-9)
	checkFloat(t, "CicRad", m.CicRad, 2.421438694000244e-07, 1e-9)
	checkFloat(t, "Omega0SemiCircles", m.Omega0SemiCircles, -0.944771918002516, 1e-9)
	checkFloat(t, "CisRad", m.CisRad, 1.6763806343078613e-08, 1e-9)
	checkFloat(t, "I0SemiCircles", m.I0SemiCircles, 0.3080678000114858, 1e-9)
	checkFloat(t, "CrcM", m.CrcM, 210.3125, 1e-9)
	checkFloat(t, "OmegaSemiCircles", m.OmegaSemiCircles, -0.3891187282279134, 1e-9)
	checkFloat(t, "OmegaDotSemiCirclesPerS", m.OmegaDotSemiCirclesPerS, -2.476781446603127e-09, 1e-9)
	checkFloat(t, "TgdS", m.TgdS, -1.7695128917694092e-08, 1e-9)
}
