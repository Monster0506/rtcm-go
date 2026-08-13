package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1064(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x12, 0x42, 0x87, 0x53, 0x02, 0xcc, 0x3e, 0x84, 0x81, 0x2f,
		0xc9, 0xbf, 0x23, 0x64, 0x0e, 0xff, 0x5d, 0x3d, 0x60, 0xaf, 0x15, 0xe3,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1064(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1064: unexpected error: %v", err)
	}
	if m.MessageType != 1064 || m.EpochTimeS != 60000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSRHeader)
	}
	testutil.CheckFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.Corrections) != 1 {
		t.Fatalf("len(Corrections)")
	}
	if m.Corrections[0].SatelliteID != 5 {
		t.Fatalf("clock correction mismatch")
	}
	testutil.CheckFloat(t, "DeltaClockC0M", m.Corrections[0].DeltaClockC0M, -11.1111, 1e-9)
	testutil.CheckFloat(t, "DeltaClockC1MPerS", m.Corrections[0].DeltaClockC1MPerS, 0.222222, 1e-9)
	testutil.CheckFloat(t, "DeltaClockC2MPerS2", m.Corrections[0].DeltaClockC2MPerS2, -0.0066666600000000005, 1e-9)
}
