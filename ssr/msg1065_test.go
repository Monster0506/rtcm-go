package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1065(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0e, 0x42, 0x97, 0x53, 0x02, 0xcc, 0x3e, 0x84, 0x81, 0x28,
		0x80, 0x0f, 0xa1, 0xfc, 0x7c, 0xb5, 0x9a, 0x73,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1065(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1065: unexpected error: %v", err)
	}
	if m.MessageType != 1065 || m.EpochTimeS != 60000 ||
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
