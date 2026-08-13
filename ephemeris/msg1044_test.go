package ephemeris

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1044(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x3d, 0x41, 0x43, 0x30, 0x39, 0xf9, 0x01, 0x23, 0xd8, 0x10,
		0x3c, 0xab, 0xf0, 0x60, 0x4e, 0x20, 0x1d, 0x6f, 0x34, 0x57, 0xfa, 0xfc,
		0x00, 0x3c, 0x48, 0x18, 0x0a, 0x3a, 0x86, 0xac, 0x71, 0x3b, 0x50, 0xc7,
		0xff, 0x27, 0x14, 0x86, 0x5d, 0x3c, 0x01, 0x08, 0x84, 0x7b, 0x35, 0x64,
		0x8a, 0xe3, 0xe5, 0x7b, 0x87, 0x0f, 0xff, 0x3f, 0x1f, 0xd5, 0xab, 0xe8,
		0x50, 0x3b, 0x38, 0x48, 0xa7, 0x6f, 0x64,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1044(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1044: unexpected error: %v", err)
	}
	if m.MessageType != 1044 {
		t.Fatalf("MessageType = %d, want 1044", m.MessageType)
	}
	if m.SatelliteID != 3 {
		t.Fatalf("SatelliteID = %d, want 3", m.SatelliteID)
	}
	if m.IODE != 42 {
		t.Fatalf("IODE = %d, want 42", m.IODE)
	}
	if m.CodesOnL2 != 2 {
		t.Fatalf("CodesOnL2 = %d, want 2", m.CodesOnL2)
	}
	if m.WeekNumber != 1000 {
		t.Fatalf("WeekNumber = %d, want 1000", m.WeekNumber)
	}
	if m.URA != 5 {
		t.Fatalf("URA = %d, want 5", m.URA)
	}
	if m.SVHealth != 0 {
		t.Fatalf("SVHealth = %d, want 0", m.SVHealth)
	}
	if m.IODC != 900 {
		t.Fatalf("IODC = %d, want 900", m.IODC)
	}
	if m.FitInterval != true {
		t.Fatalf("FitInterval = %v, want true", m.FitInterval)
	}
	testutil.CheckFloat(t, "TocS", m.TocS, 197520, 1e-9)
	testutil.CheckFloat(t, "Af2SPerS2", m.Af2SPerS2, -1.942890293094024e-16, 1e-9)
	testutil.CheckFloat(t, "Af1SPerS", m.Af1SPerS, 3.3082869776990265e-11, 1e-9)
	testutil.CheckFloat(t, "Af0S", m.Af0S, -0.00030469195917248726, 1e-9)
	testutil.CheckFloat(t, "CrsM", m.CrsM, -31.25, 1e-9)
	testutil.CheckFloat(t, "DeltaNSemiCirclesPerS", m.DeltaNSemiCirclesPerS, 5.684341886080801e-10, 1e-9)
	testutil.CheckFloat(t, "M0SemiCircles", m.M0SemiCircles, 0.057489047292619944, 1e-9)
	testutil.CheckFloat(t, "CucRad", m.CucRad, -5.979090929031372e-07, 1e-9)
	testutil.CheckFloat(t, "Eccentricity", m.Eccentricity, 0.00011497805826365948, 1e-9)
	testutil.CheckFloat(t, "CusRad", m.CusRad, 1.218169927597046e-06, 1e-9)
	testutil.CheckFloat(t, "SqrtA", m.SqrtA, 5173.388820648193, 1e-9)
	testutil.CheckFloat(t, "ToeS", m.ToeS, 869136, 1e-9)
	testutil.CheckFloat(t, "CicRad", m.CicRad, -1.0244548320770264e-07, 1e-9)
	testutil.CheckFloat(t, "Omega0SemiCircles", m.Omega0SemiCircles, -0.45991238253191113, 1e-9)
	testutil.CheckFloat(t, "CisRad", m.CisRad, 1.2293457984924316e-07, 1e-9)
	testutil.CheckFloat(t, "I0SemiCircles", m.I0SemiCircles, 0.2587525066919625, 1e-9)
	testutil.CheckFloat(t, "CrcM", m.CrcM, 277.75, 1e-9)
	testutil.CheckFloat(t, "OmegaSemiCircles", m.OmegaSemiCircles, -0.05179193476215005, 1e-9)
	testutil.CheckFloat(t, "OmegaDotSemiCirclesPerS", m.OmegaDotSemiCirclesPerS, -1.4034640116733499e-09, 1e-9)
	testutil.CheckFloat(t, "IDOTSemiCirclesPerS", m.IDOTSemiCirclesPerS, -7.707967597525567e-11, 1e-9)
	testutil.CheckFloat(t, "TgdS", m.TgdS, -9.313225746154785e-09, 1e-9)
}
