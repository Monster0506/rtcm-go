package ephemeris

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1020(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x2d, 0x3f, 0xc2, 0x4b, 0xb3, 0x78, 0xcf, 0xa0, 0xf4, 0x96,
		0x4c, 0xb5, 0xd1, 0xa0, 0x0d, 0x84, 0xba, 0x00, 0x21, 0x1b, 0xf2, 0xa7,
		0xf6, 0x48, 0xbf, 0x59, 0x16, 0x63, 0x80, 0x2d, 0x16, 0xf4, 0xa5, 0x01,
		0x50, 0x92, 0xc2, 0x4c, 0x00, 0x00, 0x00, 0x1a, 0x00, 0x00, 0x08, 0x00,
		0xc2, 0xb9, 0xe5,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1020(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1020: unexpected error: %v", err)
	}
	if m.MessageType != 1020 {
		t.Fatalf("MessageType = %d, want 1020", m.MessageType)
	}
	if m.SatelliteID != 9 {
		t.Fatalf("SatelliteID = %d, want 9", m.SatelliteID)
	}
	if m.FrequencyChannelNumber != 5 {
		t.Fatalf("FrequencyChannelNumber = %d, want 5", m.FrequencyChannelNumber)
	}
	if m.P1 != 1 {
		t.Fatalf("P1 = %d, want 1", m.P1)
	}
	if m.Tk != 2492 {
		t.Fatalf("Tk = %d, want 2492", m.Tk)
	}
	if m.En != 0 {
		t.Fatalf("En = %d, want 0", m.En)
	}
	if m.MFT != 5 {
		t.Fatalf("MFT = %d, want 5", m.MFT)
	}
	if m.MNT != 73 {
		t.Fatalf("MNT = %d, want 73", m.MNT)
	}
	if m.MM != 1 {
		t.Fatalf("MM = %d, want 1", m.MM)
	}
	if m.NA != 73 {
		t.Fatalf("NA = %d, want 73", m.NA)
	}
	if m.MN4 != 8 {
		t.Fatalf("MN4 = %d, want 8", m.MN4)
	}
	if m.MP != 3 {
		t.Fatalf("MP = %d, want 3", m.MP)
	}
	if m.AlmanacHealth != true {
		t.Fatalf("AlmanacHealth = %v, want true", m.AlmanacHealth)
	}
	if m.AlmanacHealthAvailability != true {
		t.Fatalf("AlmanacHealthAvailability = %v, want true", m.AlmanacHealthAvailability)
	}
	if m.MSbOfBn != false {
		t.Fatalf("MSbOfBn = %v, want false", m.MSbOfBn)
	}
	if m.P2 != true {
		t.Fatalf("P2 = %v, want true", m.P2)
	}
	if m.P3 != true {
		t.Fatalf("P3 = %v, want true", m.P3)
	}
	if m.MLn3rdString != false {
		t.Fatalf("MLn3rdString = %v, want false", m.MLn3rdString)
	}
	if m.MP4 != true {
		t.Fatalf("MP4 = %v, want true", m.MP4)
	}
	if m.AvailabilityOfAdditionalData != true {
		t.Fatalf("AvailabilityOfAdditionalData = %v, want true", m.AvailabilityOfAdditionalData)
	}
	if m.MLn5thString != false {
		t.Fatalf("MLn5thString = %v, want false", m.MLn5thString)
	}
	testutil.CheckFloat(t, "TbMin", m.TbMin, 1185, 1e-9)
	testutil.CheckFloat(t, "XnFirstDerivativeKmPerS", m.XnFirstDerivativeKmPerS, -2.059713363647461, 1e-9)
	testutil.CheckFloat(t, "XnKm", m.XnKm, 19637.81884765625, 1e-9)
	testutil.CheckFloat(t, "XnSecondDerivativeKmPerS2", m.XnSecondDerivativeKmPerS2, 0.0, 1e-9)
	testutil.CheckFloat(t, "YnFirstDerivativeKmPerS", m.YnFirstDerivativeKmPerS, 0.8449039459228516, 1e-9)
	testutil.CheckFloat(t, "YnKm", m.YnKm, 33.10888671875, 1e-9)
	testutil.CheckFloat(t, "YnSecondDerivativeKmPerS2", m.YnSecondDerivativeKmPerS2, -1.862645149230957e-09, 1e-9)
	testutil.CheckFloat(t, "ZnFirstDerivativeKmPerS", m.ZnFirstDerivativeKmPerS, -2.4976272583007812, 1e-9)
	testutil.CheckFloat(t, "ZnKm", m.ZnKm, -16217.08740234375, 1e-9)
	testutil.CheckFloat(t, "ZnSecondDerivativeKmPerS2", m.ZnSecondDerivativeKmPerS2, 2.7939677238464355e-09, 1e-9)
	testutil.CheckFloat(t, "GammaN", m.GammaN, 1.8189894035458565e-12, 1e-9)
	testutil.CheckFloat(t, "TauNS", m.TauNS, -0.00017513707280158997, 1e-9)
	testutil.CheckFloat(t, "MDeltaTauS", m.MDeltaTauS, -3.725290298461914e-09, 1e-9)
	testutil.CheckFloat(t, "TauCS", m.TauCS, -1.3969838619232178e-09, 1e-9)
	testutil.CheckFloat(t, "MTauGPSS", m.MTauGPSS, 7.450580596923828e-09, 1e-9)
}
