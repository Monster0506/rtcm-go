package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1068(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0c, 0x42, 0xc7, 0x53, 0x02, 0xcc, 0x3e, 0x84, 0x81, 0x2f,
		0x93, 0x7e, 0x40, 0x19, 0xe5, 0x8d,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1068(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1068: unexpected error: %v", err)
	}
	if m.MessageType != 1068 || m.EpochTimeS != 60000 ||
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
