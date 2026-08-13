package misc

import (
	"testing"

	"github.com/Monster0506/rtcm-go/core"
)

func TestDecodeMsg1029(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x13, 0x40, 0x50, 0x17, 0x00, 0x84, 0x73, 0x6e, 0x0a, 0x0a,
		0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x72, 0x74, 0x63, 0x6d, 0x59, 0x9b,
		0x37,
	}
	payload, _, err := core.ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1029(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1029: unexpected error: %v", err)
	}
	if m.MessageType != 1029 || m.StationID != 23 ||
		m.MJDNumber != 132 || m.SecondsOfDayUTC != 59100 ||
		m.NumCharacters != 10 || m.Text != "hello rtcm" {
		t.Fatalf("mismatch: %+v", m)
	}
}
