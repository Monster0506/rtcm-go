package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1062(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0c, 0x42, 0x64, 0xda, 0x30, 0x59, 0x87, 0xd0, 0x90, 0x22,
		0xf9, 0x37, 0xe4, 0x03, 0xe2, 0x15,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1062(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1062: unexpected error: %v", err)
	}
	if m.MessageType != 1062 || m.EpochTimeS != 318000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSRHeader)
	}
	testutil.CheckFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.Corrections) != 1 || m.Corrections[0].SatelliteID != 5 {
		t.Fatalf("mismatch: %+v", m.Corrections)
	}
	testutil.CheckFloat(t, "HighRateClockCorrectionM", m.Corrections[0].HighRateClockCorrectionM, -22.2222, 1e-9)
}
