package rtcm

import "testing"



func TestDecodeMsg1022(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x54, 0x3f, 0xe4, 0xa2, 0xa8, 0x29, 0xa3, 0x9d, 0x1a, 0x19,
		0x99, 0x1b, 0x29, 0x15, 0x41, 0x4d, 0x1c, 0xe8, 0xc8, 0xd4, 0xe0, 0xcc,
		0xc8, 0x1c, 0x03, 0x48, 0x1d, 0x09, 0x00, 0x5e, 0xe0, 0x13, 0x88, 0x5d,
		0xc0, 0x0f, 0x12, 0x07, 0xc6, 0xbb, 0x90, 0xa8, 0xc9, 0xc0, 0x01, 0xe8,
		0x48, 0x1f, 0xfe, 0x17, 0xb8, 0x00, 0x00, 0xf4, 0x24, 0x03, 0xd0, 0x90,
		0x00, 0x0e, 0xb7, 0x9a, 0x2b, 0xf1, 0x48, 0x65, 0xd3, 0xc1, 0x08, 0xe8,
		0xd7, 0x18, 0x7a, 0x12, 0x00, 0x7a, 0x12, 0x00, 0x5b, 0x8d, 0x80, 0x4c,
		0x4b, 0x40, 0xe0, 0xbc, 0x8f, 0xc3,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1022(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1022: unexpected error: %v", err)
	}
	if m.MessageType != 1022 || m.SourceName != "EPSG:4326" ||
		m.TargetName != "EPSG:25832" ||
		m.SystemIdentificationNumber != 7 ||
		m.UtilizedTransformationMessageIndicator != 3 ||
		m.PlateNumber != 9 || m.ComputationIndicator != 0 ||
		m.HeightIndicator != 0 {
		t.Fatalf("common header mismatch: %+v", m)
	}
	checkFloat(t, "LatitudeOfOriginArcSec", m.LatitudeOfOriginArcSec, -97152, 1e-9)
	checkFloat(t, "LongitudeOfOriginArcSec", m.LongitudeOfOriginArcSec, 194304, 1e-9)
	checkFloat(t, "NSExtensionArcSec", m.NSExtensionArcSec, 10000, 1e-9)
	checkFloat(t, "EWExtensionArcSec", m.EWExtensionArcSec, 12000, 1e-9)
	checkFloat(t, "TranslationXM", m.TranslationXM, 123.456, 1e-9)
	checkFloat(t, "TranslationYM", m.TranslationYM, -234.567, 1e-9)
	checkFloat(t, "TranslationZM", m.TranslationZM, 345.678, 1e-9)
	checkFloat(t, "RotationXArcSec", m.RotationXArcSec, 20.0, 1e-9)
	checkFloat(t, "RotationYArcSec", m.RotationYArcSec, -20.0, 1e-9)
	checkFloat(t, "RotationZArcSec", m.RotationZArcSec, 10.0, 1e-9)
	checkFloat(t, "ScaleCorrectionPPM", m.ScaleCorrectionPPM, 40.0, 1e-9)
	checkFloat(t, "XPM", m.XPM, 123456.789, 1e-9)
	checkFloat(t, "YPM", m.YPM, -987654.321, 1e-9)
	checkFloat(t, "ZPM", m.ZPM, 555555.555, 1e-9)
	if m.HorizontalQualityIndicator != 3 ||
		m.VerticalQualityIndicator != 4 {
		t.Fatalf("tail mismatch: %+v", m)
	}
	checkFloat(t, "AddSemiMajorAxisSourceM", m.AddSemiMajorAxisSourceM, 1000.0, 1e-9)
	checkFloat(t, "AddSemiMinorAxisSourceM", m.AddSemiMinorAxisSourceM, 2000.0, 1e-9)
	checkFloat(t, "AddSemiMajorAxisTargetM", m.AddSemiMajorAxisTargetM, 1500.0, 1e-9)
	checkFloat(t, "AddSemiMinorAxisTargetM", m.AddSemiMinorAxisTargetM, 2500.0, 1e-9)
}
