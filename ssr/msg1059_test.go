package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1059(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0f, 0x42, 0x34, 0xda, 0x30, 0x59, 0x87, 0xd0, 0x90, 0x22,
		0x88, 0x00, 0xfa, 0x1f, 0xc7, 0xc0, 0xda, 0x64, 0xe8,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1059(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1059: unexpected error: %v", err)
	}
	if m.MessageType != 1059 || m.EpochTimeS != 318000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSRHeader)
	}
	testutil.CheckFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.SatelliteBiases) != 1 || m.SatelliteBiases[0].SatelliteID != 5 {
		t.Fatalf("mismatch: %+v", m.SatelliteBiases)
	}
	if len(m.SatelliteBiases[0].CodeBiases) != 2 {
		t.Fatalf("len(CodeBiases)")
	}
	if m.SatelliteBiases[0].CodeBiases[0].SignalTrackingMode != 0 {
		t.Fatalf("SignalTrackingMode mismatch")
	}
	testutil.CheckFloat(t, "CodeBiasM", m.SatelliteBiases[0].CodeBiases[0].CodeBiasM, 5.0, 1e-9)
	if m.SatelliteBiases[0].CodeBiases[1].SignalTrackingMode != 7 {
		t.Fatalf("SignalTrackingMode mismatch")
	}
	testutil.CheckFloat(t, "CodeBiasM", m.SatelliteBiases[0].CodeBiases[1].CodeBiasM, -9.0, 1e-9)
}
