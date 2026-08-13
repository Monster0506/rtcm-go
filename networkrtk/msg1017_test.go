package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1017(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x17, 0x3f, 0x90, 0x31, 0x09, 0xb4, 0x61, 0x06, 0x40, 0x65,
		0x21, 0x55, 0xfa, 0x24, 0x21, 0x01, 0xc2, 0x19, 0xd0, 0x19, 0x06, 0x47,
		0xf5, 0x10, 0xa6, 0x29, 0x07,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1017(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1017: unexpected error: %v", err)
	}
	if m.MessageType != 1017 || m.NetworkID != 3 ||
		m.SubnetworkID != 1 || m.MultipleMessageIndicator != true ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 || m.NumGPSSats != 2 {
		t.Fatalf("header mismatch: %+v", m.NetworkRTKHeader)
	}
	testutil.CheckFloat(t, "GPSEpochTimeS", m.GPSEpochTimeS, 31800.0, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].AmbiguityStatusFlag != 1 || m.Satellites[0].NonSyncCount != 2 || m.Satellites[0].IODE != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	testutil.CheckFloat(t, "Satellites[0].GeometricCorrectionDiffM", m.Satellites[0].GeometricCorrectionDiffM, -0.75, 1e-9)
	testutil.CheckFloat(t, "Satellites[0].IonosphericCorrectionDiffM", m.Satellites[0].IonosphericCorrectionDiffM, 0.45, 1e-9)
	if m.Satellites[1].SatelliteID != 12 || m.Satellites[1].AmbiguityStatusFlag != 3 || m.Satellites[1].NonSyncCount != 5 || m.Satellites[1].IODE != 200 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	testutil.CheckFloat(t, "Satellites[1].GeometricCorrectionDiffM", m.Satellites[1].GeometricCorrectionDiffM, 0.4, 1e-9)
	testutil.CheckFloat(t, "Satellites[1].IonosphericCorrectionDiffM", m.Satellites[1].IonosphericCorrectionDiffM, -0.35000000000000003, 1e-9)
}
