package ephemeris

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1046(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x3f, 0x41, 0x61, 0x54, 0x04, 0x16, 0x6b, 0xfb, 0xb1, 0x4b,
		0x40, 0x00, 0x01, 0xf4, 0x13, 0x5e, 0x68, 0xc3, 0xe9, 0xe8, 0xa0, 0xbc,
		0x23, 0x37, 0x08, 0x3f, 0xef, 0x60, 0x00, 0x7d, 0xaa, 0xcc, 0x22, 0xe6,
		0xa8, 0x12, 0xf4, 0x19, 0x4b, 0x4f, 0xff, 0xde, 0x0a, 0x11, 0x82, 0xcf,
		0xff, 0xd2, 0x72, 0x3a, 0xe3, 0xe1, 0xf0, 0x5c, 0x6c, 0xc0, 0x7c, 0x0f,
		0xfb, 0xf1, 0xe0, 0x4c, 0x15, 0x00, 0x25, 0x41, 0xd4,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1046(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1046: unexpected error: %v", err)
	}
	if m.MessageType != 1046 {
		t.Fatalf("MessageType = %d, want 1046", m.MessageType)
	}
	if m.SatelliteID != 5 {
		t.Fatalf("SatelliteID = %d, want 5", m.SatelliteID)
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
	if m.E5bSignalHealthStatus != 0 {
		t.Fatalf("E5bSignalHealthStatus = %d, want 0", m.E5bSignalHealthStatus)
	}
	if m.E1BSignalHealthStatus != 0 {
		t.Fatalf("E1BSignalHealthStatus = %d, want 0", m.E1BSignalHealthStatus)
	}
	if m.E5bDataValidityStatus != false {
		t.Fatalf("E5bDataValidityStatus = %v, want false", m.E5bDataValidityStatus)
	}
	if m.E1BDataValidityStatus != false {
		t.Fatalf("E1BDataValidityStatus = %v, want false", m.E1BDataValidityStatus)
	}
	testutil.CheckFloat(t, "IDOTSemiCirclesPerS", m.IDOTSemiCirclesPerS, -3.1377567211166024e-11, 1e-9)
	testutil.CheckFloat(t, "TocS", m.TocS, 318000, 1e-9)
	testutil.CheckFloat(t, "Af2SPerS2", m.Af2SPerS2, 0.0, 1e-9)
	testutil.CheckFloat(t, "Af1SPerS", m.Af1SPerS, 3.552713678800501e-12, 1e-9)
	testutil.CheckFloat(t, "Af0S", m.Af0S, 0.004728707484900951, 1e-9)
	testutil.CheckFloat(t, "CrsM", m.CrsM, -44.1875, 1e-9)
	testutil.CheckFloat(t, "DeltaNSemiCirclesPerS", m.DeltaNSemiCirclesPerS, 1.169496499642264e-09, 1e-9)
	testutil.CheckFloat(t, "M0SemiCircles", m.M0SemiCircles, 0.06877923710271716, 1e-9)
	testutil.CheckFloat(t, "CucRad", m.CucRad, -1.9818544387817383e-06, 1e-9)
	testutil.CheckFloat(t, "Eccentricity", m.Eccentricity, 0.00023969111498445272, 1e-9)
	testutil.CheckFloat(t, "CusRad", m.CusRad, 4.159286618232727e-06, 1e-9)
	testutil.CheckFloat(t, "SqrtA", m.SqrtA, 5440.592296600342, 1e-9)
	testutil.CheckFloat(t, "ToeS", m.ToeS, 318000, 1e-9)
	testutil.CheckFloat(t, "CicRad", m.CicRad, -5.587935447692871e-09, 1e-9)
	testutil.CheckFloat(t, "Omega0SemiCircles", m.Omega0SemiCircles, -0.24508378840982914, 1e-9)
	testutil.CheckFloat(t, "CisRad", m.CisRad, -5.587935447692871e-09, 1e-9)
	testutil.CheckFloat(t, "I0SemiCircles", m.I0SemiCircles, 0.30577638652175665, 1e-9)
	testutil.CheckFloat(t, "CrcM", m.CrcM, 248.15625, 1e-9)
	testutil.CheckFloat(t, "OmegaSemiCircles", m.OmegaSemiCircles, -0.446898490190506, 1e-9)
	testutil.CheckFloat(t, "OmegaDotSemiCirclesPerS", m.OmegaDotSemiCirclesPerS, -1.8883383745560423e-09, 1e-9)
	testutil.CheckFloat(t, "BGDE5aE1S", m.BGDE5aE1S, 4.423782229423523e-09, 1e-9)
	testutil.CheckFloat(t, "BGDE5bE1S", m.BGDE5bE1S, 4.889443516731262e-09, 1e-9)
}
