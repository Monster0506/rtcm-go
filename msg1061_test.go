package rtcm

import "testing"



func TestDecodeMsg1061(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x0a, 0x42, 0x54, 0xda, 0x30, 0x59, 0x87, 0xd0, 0x90, 0x22,
		0xd6, 0xa1, 0xfa, 0x9a,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1061(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1061: unexpected error: %v", err)
	}
	if m.MessageType != 1061 || m.EpochTimeS != 318000 ||
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
