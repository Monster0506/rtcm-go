package rtcm

import "testing"



func TestDecodeMsg1067(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0a, 0x42, 0xb7, 0x53, 0x02, 0xcc, 0x3e, 0x84, 0x81, 0x2d,
		0x60, 0x5d, 0x75, 0xd5,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1067(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1067: unexpected error: %v", err)
	}
	if m.MessageType != 1067 || m.EpochTimeS != 60000 ||
		m.UpdateIntervalCode != 5 ||
		m.MultipleMessageIndicator != true ||
		m.IODSSR != 3 || m.SSRProviderID != 4001 ||
		m.SSRSolutionID != 2 || m.NumSatellites != 1 {
		t.Fatalf("header mismatch: %+v", m.SSRHeader)
	}
	checkFloat(t, "UpdateIntervalS", m.UpdateIntervalS, 30, 1e-9)
	if len(m.URAs) != 1 || m.URAs[0].SatelliteID != 5 ||
		m.URAs[0].URACode != 43 || m.URAs[0].URAClass != 5 ||
		m.URAs[0].URAValue != 3 {
		t.Fatalf("URA mismatch: %+v", m.URAs)
	}
}
