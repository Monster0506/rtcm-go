package networkrtk

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
	"github.com/Monster0506/rtcm-go/internal/testutil"
)

func TestDecodeMsg1013(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x10, 0x3f, 0x50, 0x63, 0xe5, 0xe1, 0x58, 0x78, 0x08, 0x48,
		0xfb, 0x60, 0x06, 0x48, 0x6a, 0x00, 0x0a, 0x6c, 0x2a, 0x15,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1013(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1013: unexpected error: %v", err)
	}
	if m.MessageType != 1013 || m.StationID != 99 ||
		m.MJDNumber != 58849 || m.SecondsOfDayUTC != 45296 ||
		m.LeapSeconds != 18 {
		t.Fatalf("mismatch: %+v", m)
	}
	if len(m.Announcements) != 2 {
		t.Fatalf("len(Announcements) = %d", len(m.Announcements))
	}
	if m.Announcements[0].MessageID != 1005 || m.Announcements[0].SyncFlag != true {
		t.Fatalf("Announcements[0] mismatch: %+v", m.Announcements[0])
	}
	testutil.CheckFloat(t, "Announcements[0].TransmissionIntervalS", m.Announcements[0].TransmissionIntervalS, 5.0, 1e-9)
	if m.Announcements[1].MessageID != 1077 || m.Announcements[1].SyncFlag != false {
		t.Fatalf("Announcements[1] mismatch: %+v", m.Announcements[1])
	}
	testutil.CheckFloat(t, "Announcements[1].TransmissionIntervalS", m.Announcements[1].TransmissionIntervalS, 1.0, 1e-9)
}
