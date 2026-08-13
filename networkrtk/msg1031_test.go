package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1031(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x13, 0x40, 0x76, 0xd1, 0x80, 0x19, 0x04, 0x10, 0xa1, 0x41,
		0x47, 0x86, 0x43, 0x20, 0x70, 0x50, 0x30, 0xe0, 0x40, 0x12, 0xbf, 0xf2,
		0x5e,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1031(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1031: unexpected error: %v", err)
	}
	if m.MessageType != 1031 || m.StationID != 50 || m.NRefs != 4 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "GLONASSEpochTimeS", m.GLONASSEpochTimeS, 55856, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	testutil.CheckFloat(t, "Satellites[0].SigmaOcMM", m.Satellites[0].SigmaOcMM, 5.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].SigmaOdPPM", m.Satellites[0].SigmaOdPPM, 0.2, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].SigmaOhPPM", m.Satellites[0].SigmaOhPPM, 3.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].SigmaIcMM", m.Satellites[0].SigmaIcMM, 50.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].SigmaIdPPM", m.Satellites[0].SigmaIdPPM, 2.0, 1e-9)
	if m.Satellites[1].SatelliteID != 7 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	testutil.CheckFloat(t, "Satellites[1].SigmaOcMM", m.Satellites[1].SigmaOcMM, 2.5, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].SigmaOdPPM", m.Satellites[1].SigmaOdPPM, 0.06, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].SigmaOhPPM", m.Satellites[1].SigmaOhPPM, 0.7000000000000001, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].SigmaIcMM", m.Satellites[1].SigmaIcMM, 4.0, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].SigmaIdPPM", m.Satellites[1].SigmaIdPPM, 0.09, 1e-9)
}
