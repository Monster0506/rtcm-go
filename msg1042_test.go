package rtcm

import "testing"

func TestDecodeMsg1042(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x40, 0x41, 0x23, 0x07, 0x6a, 0x1d, 0xae, 0x0d, 0x35, 0x61,
		0xfd, 0xbf, 0xdd, 0xca, 0xe4, 0x30, 0x86, 0x17, 0xcc, 0x82, 0x4d, 0x7d,
		0xe2, 0xf5, 0x5e, 0x87, 0xea, 0xa4, 0x00, 0x48, 0x1c, 0xa7, 0x85, 0x19,
		0x54, 0xa2, 0xa1, 0x07, 0x29, 0xab, 0x00, 0x01, 0x61, 0xd1, 0x8a, 0x76,
		0x03, 0xff, 0xd8, 0x28, 0x0b, 0xc4, 0xdd, 0x11, 0x21, 0xb1, 0x0d, 0x0f,
		0xce, 0x7f, 0xec, 0xfc, 0x01, 0x80, 0x10, 0x0d, 0x38, 0x18,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1042(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1042: unexpected error: %v", err)
	}
	if m.MessageType != 1042 {
		t.Fatalf("MessageType = %d, want 1042", m.MessageType)
	}
	if m.SatelliteID != 12 {
		t.Fatalf("SatelliteID = %d, want 12", m.SatelliteID)
	}
	if m.WeekNumber != 949 {
		t.Fatalf("WeekNumber = %d, want 949", m.WeekNumber)
	}
	if m.SVURAI != 0 {
		t.Fatalf("SVURAI = %d, want 0", m.SVURAI)
	}
	if m.AODE != 3 {
		t.Fatalf("AODE = %d, want 3", m.AODE)
	}
	if m.AODC != 2 {
		t.Fatalf("AODC = %d, want 2", m.AODC)
	}
	if m.SVHealth != false {
		t.Fatalf("SVHealth = %v, want false", m.SVHealth)
	}
	checkFloat(t, "IDOTSemiCirclesPerS", m.IDOTSemiCirclesPerS, -1.3505996321327984e-10, 1e-9)
	checkFloat(t, "TocS", m.TocS, 316800, 1e-9)
	checkFloat(t, "A2SPerS2", m.A2SPerS2, -1.3552527156068805e-19, 1e-9)
	checkFloat(t, "A1SPerS", m.A1SPerS, -7.778666599733697e-12, 1e-9)
	checkFloat(t, "A0S", m.A0S, -0.00021217693574726582, 1e-9)
	checkFloat(t, "CrsM", m.CrsM, -102.984375, 1e-9)
	checkFloat(t, "DeltaNSemiCirclesPerS", m.DeltaNSemiCirclesPerS, 1.1275460565229878e-09, 1e-9)
	checkFloat(t, "M0SemiCircles", m.M0SemiCircles, -0.11344346264377236, 1e-9)
	checkFloat(t, "CucRad", m.CucRad, -5.0924718379974365e-06, 1e-9)
	checkFloat(t, "Eccentricity", m.Eccentricity, 0.001100340741686523, 1e-9)
	checkFloat(t, "CusRad", m.CusRad, 4.862435162067413e-06, 1e-9)
	checkFloat(t, "SqrtA", m.SqrtA, 5282.629014968872, 1e-9)
	checkFloat(t, "ToeS", m.ToeS, 316800, 1e-9)
	checkFloat(t, "CicRad", m.CicRad, 4.0978193283081055e-08, 1e-9)
	checkFloat(t, "Omega0SemiCircles", m.Omega0SemiCircles, 0.9092594981193542, 1e-9)
	checkFloat(t, "CisRad", m.CisRad, -1.862645149230957e-08, 1e-9)
	checkFloat(t, "I0SemiCircles", m.I0SemiCircles, 0.31285916129127145, 1e-9)
	checkFloat(t, "CrcM", m.CrcM, 274.09375, 1e-9)
	checkFloat(t, "OmegaSemiCircles", m.OmegaSemiCircles, -0.4671555492095649, 1e-9)
	checkFloat(t, "OmegaDotSemiCirclesPerS", m.OmegaDotSemiCirclesPerS, -2.2137101041153073e-09, 1e-9)
	checkFloat(t, "Tgd1S", m.Tgd1S, 2.4e-09, 1e-9)
	checkFloat(t, "Tgd2S", m.Tgd2S, 4e-10, 1e-9)
}
