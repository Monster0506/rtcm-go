package rtcm

import "testing"



func TestDecodeMsg1027(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x21, 0x40, 0x30, 0x31, 0x61, 0xdc, 0xd6, 0x50, 0x07, 0x88,
		0xca, 0x6c, 0x00, 0x6f, 0xc2, 0x3a, 0xc0, 0x1f, 0xc2, 0xf7, 0x00, 0x01,
		0x55, 0xcc, 0x00, 0x05, 0xf5, 0xe1, 0x01, 0xff, 0x8d, 0x8f, 0x20, 0x00,
		0x0e, 0x90, 0x6a,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1027(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1027: unexpected error: %v", err)
	}
	if m.MessageType != 1027 || m.SystemIdentificationNumber != 3 ||
		m.ProjectionType != 5 {
		if m.RectificationFlag != true {
			t.Fatalf("RectificationFlag mismatch")
		}
		t.Fatalf("mismatch: %+v", m)
	}
	checkFloat(t, "LatitudeOfProjectionCenterDeg", m.LatitudeOfProjectionCenterDeg, 11.0, 1e-9)
	checkFloat(t, "LongitudeOfProjectionCenterDeg", m.LongitudeOfProjectionCenterDeg, -22.0, 1e-9)
	checkFloat(t, "AzimuthOfInitialLineDeg", m.AzimuthOfInitialLineDeg, 165.0, 1e-9)
	checkFloat(t, "DiffAngleRectifiedToSkewGridDeg", m.DiffAngleRectifiedToSkewGridDeg, -0.0055, 1e-9)
	checkFloat(t, "AddScaleFactorPPM", m.AddScaleFactorPPM, 7.000000000000001, 1e-9)
	checkFloat(t, "EastingAtProjectionCenterM", m.EastingAtProjectionCenterM, 50000.0, 1e-9)
	checkFloat(t, "NorthingAtProjectionCenterM", m.NorthingAtProjectionCenterM, -30000.0, 1e-9)
}
