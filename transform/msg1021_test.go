package transform

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1021(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x47, 0x3f, 0xd4, 0xa2, 0xa8, 0x29, 0xa3, 0x9d, 0x1a, 0x19,
		0x99, 0x1b, 0x29, 0x15, 0x41, 0x4d, 0x1c, 0xe8, 0xc8, 0xd4, 0xe0, 0xcc,
		0xc8, 0x1c, 0x03, 0x48, 0x1d, 0x09, 0x00, 0x5e, 0xe0, 0x13, 0x88, 0x5d,
		0xc0, 0x0f, 0x12, 0x07, 0xc6, 0xbb, 0x90, 0xa8, 0xc9, 0xc0, 0x01, 0xe8,
		0x48, 0x1f, 0xfe, 0x17, 0xb8, 0x00, 0x00, 0xf4, 0x24, 0x03, 0xd0, 0x90,
		0x00, 0xf4, 0x24, 0x00, 0xf4, 0x24, 0x00, 0xb7, 0x1b, 0x00, 0x98, 0x96,
		0x81, 0xc0, 0xbd, 0x66, 0x4d,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1021(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1021: unexpected error: %v", err)
	}
	if m.MessageType != 1021 || m.SourceName != "EPSG:4326" ||
		m.TargetName != "EPSG:25832" ||
		m.SystemIdentificationNumber != 7 ||
		m.UtilizedTransformationMessageIndicator != 3 ||
		m.PlateNumber != 9 || m.ComputationIndicator != 0 ||
		m.HeightIndicator != 0 {
		t.Fatalf("common header mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "LatitudeOfOriginArcSec", m.LatitudeOfOriginArcSec, -97152, 1e-9)
	testutil.CheckFloat(t, "LongitudeOfOriginArcSec", m.LongitudeOfOriginArcSec, 194304, 1e-9)
	testutil.CheckFloat(t, "NSExtensionArcSec", m.NSExtensionArcSec, 10000, 1e-9)
	testutil.CheckFloat(t, "EWExtensionArcSec", m.EWExtensionArcSec, 12000, 1e-9)
	testutil.CheckFloat(t, "TranslationXM", m.TranslationXM, 123.456, 1e-9)
	testutil.CheckFloat(t, "TranslationYM", m.TranslationYM, -234.567, 1e-9)
	testutil.CheckFloat(t, "TranslationZM", m.TranslationZM, 345.678, 1e-9)
	testutil.CheckFloat(t, "RotationXArcSec", m.RotationXArcSec, 20.0, 1e-9)
	testutil.CheckFloat(t, "RotationYArcSec", m.RotationYArcSec, -20.0, 1e-9)
	testutil.CheckFloat(t, "RotationZArcSec", m.RotationZArcSec, 10.0, 1e-9)
	testutil.CheckFloat(t, "ScaleCorrectionPPM", m.ScaleCorrectionPPM, 40.0, 1e-9)
	if m.HorizontalQualityIndicator != 3 ||
		m.VerticalQualityIndicator != 4 {
		t.Fatalf("tail mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "AddSemiMajorAxisSourceM", m.AddSemiMajorAxisSourceM, 1000.0, 1e-9)
	testutil.CheckFloat(t, "AddSemiMinorAxisSourceM", m.AddSemiMinorAxisSourceM, 2000.0, 1e-9)
	testutil.CheckFloat(t, "AddSemiMajorAxisTargetM", m.AddSemiMajorAxisTargetM, 1500.0, 1e-9)
	testutil.CheckFloat(t, "AddSemiMinorAxisTargetM", m.AddSemiMinorAxisTargetM, 2500.0, 1e-9)
}
