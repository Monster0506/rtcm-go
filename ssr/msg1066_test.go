package ssr

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1066(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x22, 0x42, 0xa7, 0x53, 0x02, 0xe6, 0x1f, 0x42, 0x40, 0x94,
		0x87, 0xe7, 0x96, 0x03, 0x0d, 0x40, 0xb6, 0xc2, 0x00, 0x07, 0xd0, 0x7e,
		0xc7, 0x80, 0x2e, 0xe1, 0xf2, 0x6f, 0xc8, 0xd9, 0x03, 0xbf, 0xd7, 0x4f,
		0x58, 0x57, 0x10, 0xd3,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1066(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1066: unexpected error: %v", err)
	}
	if m.MessageType != 1066 || m.EpochTimeS != 60000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.SatelliteReferenceDatum != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSROrbitHeader)
	}
	testutil.CheckFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.Corrections) != 1 {
		t.Fatalf("len(Corrections)")
	}
	if m.Corrections[0].SatelliteID != 5 || m.Corrections[0].IOD != 33 {
		t.Fatalf("orbit correction mismatch")
	}
	testutil.CheckFloat(t, "DeltaRadialM", m.Corrections[0].DeltaRadialM, -10.0, 1e-9)
	testutil.CheckFloat(t, "DeltaAlongTrackM", m.Corrections[0].DeltaAlongTrackM, 80.0, 1e-9)
	testutil.CheckFloat(t, "DeltaCrossTrackM", m.Corrections[0].DeltaCrossTrackM, -120.0, 1e-9)
	testutil.CheckFloat(t, "DotDeltaRadialMPerS", m.Corrections[0].DotDeltaRadialMPerS, 0.004, 1e-9)
	testutil.CheckFloat(t, "DotDeltaAlongTrackMPerS", m.Corrections[0].DotDeltaAlongTrackMPerS, -0.02, 1e-9)
	testutil.CheckFloat(t, "DotDeltaCrossTrackMPerS", m.Corrections[0].DotDeltaCrossTrackMPerS, 0.024, 1e-9)
	testutil.CheckFloat(t, "DeltaClockC0M", m.Corrections[0].DeltaClockC0M, -11.1111, 1e-9)
	testutil.CheckFloat(t, "DeltaClockC1MPerS", m.Corrections[0].DeltaClockC1MPerS, 0.222222, 1e-9)
	testutil.CheckFloat(t, "DeltaClockC2MPerS2", m.Corrections[0].DeltaClockC2MPerS2, -0.0066666600000000005, 1e-9)
}
