package rtcm

import "testing"



func TestDecodeMsg1030(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x14, 0x40, 0x64, 0xda, 0x30, 0x03, 0x20, 0x82, 0x14, 0x28,
		0x28, 0xf0, 0xc8, 0x64, 0x0e, 0x0a, 0x06, 0x1c, 0x08, 0x02, 0x40, 0x78,
		0x28, 0xaa,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1030(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1030: unexpected error: %v", err)
	}
	if m.MessageType != 1030 || m.StationID != 50 || m.NRefs != 4 {
		t.Fatalf("mismatch: %+v", m)
	}
	checkFloat(t, "GPSEpochTimeS", m.GPSEpochTimeS, 318000, 1e-9)
	if len(m.Satellites) != 2 {
		t.Fatalf("len(Satellites) = %d", len(m.Satellites))
	}
	if m.Satellites[0].SatelliteID != 5 {
		t.Fatalf("Satellites[0] mismatch: %+v", m.Satellites[0])
	}
	checkFloat(t, "Satellites[0].SigmaOcMM", m.Satellites[0].SigmaOcMM, 5.0, 1e-9)
	checkFloat(t, "Satellites[0].SigmaOdPPM", m.Satellites[0].SigmaOdPPM, 0.2, 1e-9)
	checkFloat(t, "Satellites[0].SigmaOhPPM", m.Satellites[0].SigmaOhPPM, 3.0, 1e-9)
	checkFloat(t, "Satellites[0].SigmaIcMM", m.Satellites[0].SigmaIcMM, 50.0, 1e-9)
	checkFloat(t, "Satellites[0].SigmaIdPPM", m.Satellites[0].SigmaIdPPM, 2.0, 1e-9)
	if m.Satellites[1].SatelliteID != 7 {
		t.Fatalf("Satellites[1] mismatch: %+v", m.Satellites[1])
	}
	checkFloat(t, "Satellites[1].SigmaOcMM", m.Satellites[1].SigmaOcMM, 2.5, 1e-9)
	checkFloat(t, "Satellites[1].SigmaOdPPM", m.Satellites[1].SigmaOdPPM, 0.06, 1e-9)
	checkFloat(t, "Satellites[1].SigmaOhPPM", m.Satellites[1].SigmaOhPPM, 0.7000000000000001, 1e-9)
	checkFloat(t, "Satellites[1].SigmaIcMM", m.Satellites[1].SigmaIcMM, 4.0, 1e-9)
	checkFloat(t, "Satellites[1].SigmaIdPPM", m.Satellites[1].SigmaIdPPM, 0.09, 1e-9)
}
