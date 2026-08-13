package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1032(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x14, 0x40, 0x80, 0x32, 0x03, 0x34, 0xc2, 0x98, 0x0e, 0xde,
		0xef, 0xd2, 0xd2, 0xf5, 0x8a, 0xb0, 0x94, 0x19, 0x86, 0xf3, 0x30, 0xfb,
		0x59, 0x29,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1032(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1032: unexpected error: %v", err)
	}
	if m.MessageType != 1032 || m.NonPhysicalReferenceStationID != 50 ||
		m.PhysicalReferenceStationID != 51 || m.ITRFEpochYear != 19 {
		t.Fatalf("mismatch: %+v", m)
	}
	testutil.CheckFloat(t, "ECEFXM", m.ECEFXM, 1114104.5999, 1e-6)
	testutil.CheckFloat(t, "ECEFYM", m.ECEFYM, -4850729.7108000005, 1e-6)
	testutil.CheckFloat(t, "ECEFZM", m.ECEFZM, 3975521.4643, 1e-6)
}
