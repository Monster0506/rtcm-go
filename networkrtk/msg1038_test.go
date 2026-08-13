package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1038(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x13, 0x40, 0xe0, 0x31, 0x15, 0x18, 0x08, 0x32, 0x03, 0x29,
		0x0a, 0xaf, 0xd1, 0x21, 0x09, 0x9d, 0x01, 0x90, 0x64, 0x00, 0x2c, 0x16,
		0x6a,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1038(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1038: unexpected error: %v", err)
	}
	if m.MessageType != 1038 || m.NetworkID != 3 ||
		m.SubnetworkID != 1 || m.MultipleMessageIndicator != true ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 || m.NumGLONASSDataEntries != 2 {
		t.Fatalf("header mismatch: %+v", m.GLONASSNetworkRTKHeader)
	}
	testutil.CheckFloat(t, "GLONASSEpochTimeS", m.GLONASSEpochTimeS, 8640.0, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 || m.Satellites[0].AmbiguityStatusFlag != 1 || m.Satellites[0].NonSyncCount != 2 || m.Satellites[0].IOD != 33 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	testutil.CheckFloat(t, "Satellites[0].GeometricCorrectionDiffM", m.Satellites[0].GeometricCorrectionDiffM, -0.75, 1e-9)
	if m.Satellites[1].SatelliteID != 12 || m.Satellites[1].AmbiguityStatusFlag != 3 || m.Satellites[1].NonSyncCount != 5 || m.Satellites[1].IOD != 200 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	testutil.CheckFloat(t, "Satellites[1].GeometricCorrectionDiffM", m.Satellites[1].GeometricCorrectionDiffM, 0.4, 1e-9)
}
