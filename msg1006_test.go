package rtcm

import "testing"

func TestDecodeMsg1006(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x15, 0x3e, 0xe0, 0x00, 0x03, 0x84, 0x1a, 0x86, 0x92,
		0xbf, 0xb4, 0x4b, 0x4b, 0xf4, 0xfa, 0xb7, 0xdc, 0x37, 0x62, 0x8a,
		0x01, 0x57, 0x1b, 0xa9, 0xd6,
	}
	const tol = 1e-4

	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	msg, err := DecodeMsg1006(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1006: unexpected error: %v", err)
	}

	if msg.MessageType != 1006 {
		t.Fatalf("MessageType = %d, want 1006", msg.MessageType)
	}
	if msg.StationID != 0 {
		t.Fatalf("StationID = %d, want 0", msg.StationID)
	}
	if !msg.GPSIndicator || !msg.GLONASSIndicator || !msg.GalileoIndicator {
		t.Fatalf("indicators = %v/%v/%v, want true/true/true",
			msg.GPSIndicator, msg.GLONASSIndicator, msg.GalileoIndicator)
	}
	checkFloat(t, "ECEFXM", msg.ECEFXM, 1762489.6191, tol)
	checkFloat(t, "ECEFYM", msg.ECEFYM, -5027633.8438, tol)
	checkFloat(t, "ECEFZM", msg.ECEFZM, -3496008.8438000004, tol)
	checkFloat(t, "AntennaHeightM", msg.AntennaHeightM, 0.034300000000000004, tol)
}
