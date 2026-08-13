package transform

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1025(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x19, 0x40, 0x10, 0x30, 0x43, 0xb9, 0xac, 0xa0, 0x0f, 0x11,
		0x94, 0xd8, 0x00, 0x00, 0x55, 0x73, 0x00, 0x01, 0x7d, 0x78, 0x40, 0x7f,
		0xe3, 0x63, 0xc8, 0x00, 0x20, 0x73, 0x89,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1025(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1025: unexpected error: %v", err)
	}
	if m.MessageType != 1025 || m.SystemIdentificationNumber != 3 ||
		m.ProjectionType != 1 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "LatitudeOfNaturalOriginDeg", m.LatitudeOfNaturalOriginDeg, 11.0, 1e-9)
	testutil.CheckFloat(t, "LongitudeOfNaturalOriginDeg", m.LongitudeOfNaturalOriginDeg, -22.0, 1e-9)
	testutil.CheckFloat(t, "AddScaleFactorPPM", m.AddScaleFactorPPM, 7.000000000000001, 1e-9)
	testutil.CheckFloat(t, "FalseEastingM", m.FalseEastingM, 50000.0, 1e-9)
	testutil.CheckFloat(t, "FalseNorthingM", m.FalseNorthingM, -30000.0, 1e-9)
}
