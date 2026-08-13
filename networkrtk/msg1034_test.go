package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1034(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x17, 0x40, 0xa0, 0x32, 0x4d, 0xa3, 0x01, 0x0a, 0x43, 0xc1,
		0x84, 0xb1, 0x25, 0x43, 0xe8, 0x03, 0x86, 0x03, 0x27, 0x9c, 0x02, 0x59,
		0xf3, 0x80, 0x43, 0x5a, 0x6a,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1034(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1034: unexpected error: %v", err)
	}
	if m.MessageType != 1034 || m.StationID != 50 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "GPSFKPEpochTimeS", m.GPSFKPEpochTimeS, 318000, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].IODE != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	testutil.CheckFloat(t, "Satellites[0].GeometricGradientNorthPPM", m.Satellites[0].GeometricGradientNorthPPM, -5.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].GeometricGradientEastPPM", m.Satellites[0].GeometricGradientEastPPM, 6.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].IonosphericGradientNorthPPM", m.Satellites[0].IonosphericGradientNorthPPM, -70.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].IonosphericGradientEastPPM", m.Satellites[0].IonosphericGradientEastPPM, 80.0, 1e-9)
	if m.Satellites[1].SatelliteID != 7 || m.Satellites[1].IODE != 12 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	testutil.CheckFloat(t, "Satellites[1].GeometricGradientNorthPPM", m.Satellites[1].GeometricGradientNorthPPM, 1.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].GeometricGradientEastPPM", m.Satellites[1].GeometricGradientEastPPM, -2.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].IonosphericGradientNorthPPM", m.Satellites[1].IonosphericGradientNorthPPM, 3.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].IonosphericGradientEastPPM", m.Satellites[1].IonosphericGradientEastPPM, -4.0, 1e-9)
}
