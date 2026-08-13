package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1035(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x17, 0x40, 0xb0, 0x32, 0x6d, 0x18, 0x08, 0x52, 0x1e, 0x0c,
		0x25, 0x89, 0x2a, 0x1f, 0x40, 0x1c, 0x30, 0x19, 0x3c, 0xe0, 0x12, 0xcf,
		0x9c, 0x00, 0x8c, 0x37, 0x10,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1035(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1035: unexpected error: %v", err)
	}
	if m.MessageType != 1035 || m.StationID != 50 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "GLONASSFKPEpochTimeS", m.GLONASSFKPEpochTimeS, 55856, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].IOD != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	testutil.CheckFloat(t, "Satellites[0].GeometricGradientNorthPPM", m.Satellites[0].GeometricGradientNorthPPM, -5.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].GeometricGradientEastPPM", m.Satellites[0].GeometricGradientEastPPM, 6.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].IonosphericGradientNorthPPM", m.Satellites[0].IonosphericGradientNorthPPM, -70.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].IonosphericGradientEastPPM", m.Satellites[0].IonosphericGradientEastPPM, 80.0, 1e-9)
	if m.Satellites[1].SatelliteID != 7 || m.Satellites[1].IOD != 12 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	testutil.CheckFloat(t, "Satellites[1].GeometricGradientNorthPPM", m.Satellites[1].GeometricGradientNorthPPM, 1.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].GeometricGradientEastPPM", m.Satellites[1].GeometricGradientEastPPM, -2.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].IonosphericGradientNorthPPM", m.Satellites[1].IonosphericGradientNorthPPM, 3.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].IonosphericGradientEastPPM", m.Satellites[1].IonosphericGradientEastPPM, -4.0, 1e-9)
}
