package transform

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1026(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x1e, 0x40, 0x20, 0x31, 0x03, 0xb9, 0xac, 0xa0, 0x0f, 0x11,
		0x94, 0xd8, 0x00, 0x0e, 0xe6, 0xb2, 0x80, 0x04, 0x78, 0x68, 0xc0, 0x00,
		0x05, 0xf5, 0xe1, 0x01, 0xff, 0x8d, 0x8f, 0x20, 0x00, 0xdb, 0x54, 0x86,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1026(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1026: unexpected error: %v", err)
	}
	if m.MessageType != 1026 || m.SystemIdentificationNumber != 3 ||
		m.ProjectionType != 4 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "LatitudeOfFalseOriginDeg", m.LatitudeOfFalseOriginDeg, 11.0, 1e-9)
	testutil.CheckFloat(t, "LongitudeOfFalseOriginDeg", m.LongitudeOfFalseOriginDeg, -22.0, 1e-9)
	testutil.CheckFloat(t, "LatitudeOfStandardParallel1Deg", m.LatitudeOfStandardParallel1Deg, 5.5, 1e-9)
	testutil.CheckFloat(t, "LatitudeOfStandardParallel2Deg", m.LatitudeOfStandardParallel2Deg, 6.6, 1e-9)
	testutil.CheckFloat(t, "EastingOfFalseOriginM", m.EastingOfFalseOriginM, 50000.0, 1e-9)
	testutil.CheckFloat(t, "NorthingOfFalseOriginM", m.NorthingOfFalseOriginM, -30000.0, 1e-9)
}
