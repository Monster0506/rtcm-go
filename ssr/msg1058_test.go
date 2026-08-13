package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1058(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x12, 0x42, 0x24, 0xda, 0x30, 0x59, 0x87, 0xd0, 0x90, 0x22,
		0xfc, 0x9b, 0xf2, 0x36, 0x40, 0xef, 0xf5, 0xd3, 0xd6, 0x11, 0xb4, 0xf8,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1058(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1058: unexpected error: %v", err)
	}
	if m.MessageType != 1058 || m.EpochTimeS != 318000 ||
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
