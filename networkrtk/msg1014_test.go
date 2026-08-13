package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1014(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0f, 0x3f, 0x60, 0x52, 0x18, 0x32, 0x03, 0x2f, 0xe7, 0xe3,
		0x83, 0x50, 0xc7, 0xfe, 0xc7, 0x88, 0xc3, 0x08, 0x4d,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1014(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1014: unexpected error: %v", err)
	}
	if m.MessageType != 1014 || m.NetworkID != 5 ||
		m.SubnetworkID != 2 || m.NumAuxStations != 3 ||
		m.MasterReferenceStationID != 100 ||
		m.AuxiliaryReferenceStationID != 101 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "DeltaLatitudeDeg", m.DeltaLatitudeDeg, -0.30862500000000004, 1e-9)
	testutil.CheckFloat(t, "DeltaLongitudeDeg", m.DeltaLongitudeDeg, 1.358025, 1e-9)
	testutil.CheckFloat(t, "DeltaHeightM", m.DeltaHeightM, -9.999, 1e-9)
}
