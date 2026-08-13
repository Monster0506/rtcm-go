package ephemeris

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1045(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x3e, 0x41, 0x50, 0xd4, 0x04, 0x16, 0x6b, 0xfb, 0xb9, 0x4b,
		0x40, 0x3f, 0xfe, 0x89, 0xff, 0x97, 0x1c, 0x1b, 0xeb, 0xf0, 0xa0, 0x52,
		0xea, 0xd2, 0x45, 0x93, 0xf0, 0x3c, 0x00, 0x76, 0x35, 0x3c, 0x23, 0xfa,
		0xa8, 0x12, 0xf5, 0x11, 0x4b, 0x4f, 0xfe, 0xfe, 0x0a, 0x10, 0xff, 0x7f,
		0xfe, 0xf2, 0x72, 0x38, 0xa5, 0xd1, 0xef, 0xdf, 0x52, 0x3a, 0x35, 0x4f,
		0xfb, 0xf9, 0x90, 0x34, 0x00, 0xe7, 0x2f, 0xc4,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1045(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1045: unexpected error: %v", err)
	}
	if m.MessageType != 1045 {
		t.Fatalf("MessageType = %d, want 1045", m.MessageType)
	}
	if m.SatelliteID != 3 {
		t.Fatalf("SatelliteID = %d, want 3", m.SatelliteID)
	}
	if m.WeekNumber != 1281 {
		t.Fatalf("WeekNumber = %d, want 1281", m.WeekNumber)
	}
	if m.IODNav != 22 {
		t.Fatalf("IODNav = %d, want 22", m.IODNav)
	}
	if m.SISA != 107 {
		t.Fatalf("SISA = %d, want 107", m.SISA)
	}
	if m.OSHS != 0 {
		t.Fatalf("OSHS = %d, want 0", m.OSHS)
	}
	if m.OSDVS != false {
		t.Fatalf("OSDVS = %v, want false", m.OSDVS)
	}
	testutil.CheckFloat(t, "IDOTSemiCirclesPerS", m.IDOTSemiCirclesPerS, -3.115019353572279e-11, 1e-9)
	testutil.CheckFloat(t, "TocS", m.TocS, 318000, 1e-9)
	testutil.CheckFloat(t, "Af2SPerS2", m.Af2SPerS2, 0.0, 1e-9)
	testutil.CheckFloat(t, "Af1SPerS", m.Af1SPerS, -2.6716406864579767e-12, 1e-9)
	testutil.CheckFloat(t, "Af0S", m.Af0S, -0.00010003114584833384, 1e-9)
	testutil.CheckFloat(t, "CrsM", m.CrsM, -40.125, 1e-9)
	testutil.CheckFloat(t, "DeltaNSemiCirclesPerS", m.DeltaNSemiCirclesPerS, 1.1664269550237805e-09, 1e-9)
	testutil.CheckFloat(t, "M0SemiCircles", m.M0SemiCircles, -0.5413645040243864, 1e-9)
	testutil.CheckFloat(t, "CucRad", m.CucRad, -1.8794089555740356e-06, 1e-9)
	testutil.CheckFloat(t, "Eccentricity", m.Eccentricity, 0.00022546376567333937, 1e-9)
	testutil.CheckFloat(t, "CusRad", m.CusRad, 4.287809133529663e-06, 1e-9)
	testutil.CheckFloat(t, "SqrtA", m.SqrtA, 5440.592414855957, 1e-9)
	testutil.CheckFloat(t, "ToeS", m.ToeS, 318000, 1e-9)
	testutil.CheckFloat(t, "CicRad", m.CicRad, -3.166496753692627e-08, 1e-9)
	testutil.CheckFloat(t, "Omega0SemiCircles", m.Omega0SemiCircles, -0.24508476676419377, 1e-9)
	testutil.CheckFloat(t, "CisRad", m.CisRad, -3.166496753692627e-08, 1e-9)
	testutil.CheckFloat(t, "I0SemiCircles", m.I0SemiCircles, 0.3057721094228327, 1e-9)
	testutil.CheckFloat(t, "CrcM", m.CrcM, 247.90625, 1e-9)
	testutil.CheckFloat(t, "OmegaSemiCircles", m.OmegaSemiCircles, -0.08484991453588009, 1e-9)
	testutil.CheckFloat(t, "OmegaDotSemiCirclesPerS", m.OmegaDotSemiCirclesPerS, -1.8743548935162835e-09, 1e-9)
	testutil.CheckFloat(t, "BGDE5aE1S", m.BGDE5aE1S, 3.026798367500305e-09, 1e-9)
}
