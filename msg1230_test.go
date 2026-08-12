package rtcm

import "testing"



func TestDecodeMsg1230(t *testing.T) {
	frame := []byte{
		0xd3, 0x00, 0x08, 0x4c, 0xe0, 0x07, 0x8a, 0x04, 0xd2, 0xef, 0x1f, 0x60,
		0x6e, 0x5a,
	}
	payload, _, err := ParseFrame(frame)
	if err != nil {
		t.Fatalf("ParseFrame: unexpected error: %v", err)
	}
	m, err := DecodeMsg1230(payload)
	if err != nil {
		t.Fatalf("DecodeMsg1230: unexpected error: %v", err)
	}
	if m.MessageType != 1230 || m.StationID != 7 ||
		m.CodePhaseBiasIndicator != true || m.FDMASignalsMask != 10 {
		t.Fatalf("mismatch: %+v", m)
	}
	if m.L1PCodePhaseBiasM != nil || m.L2PCodePhaseBiasM != nil {
		t.Fatalf("expected L1P/L2P absent: %+v", m)
	}
	if m.L1CACodePhaseBiasM == nil {
		t.Fatal("expected L1CACodePhaseBiasM present")
	}
	checkFloat(t, "L1CACodePhaseBiasM", *m.L1CACodePhaseBiasM, 24.68, 1e-9)
	if m.L2CACodePhaseBiasM == nil {
		t.Fatal("expected L2CACodePhaseBiasM present")
	}
	checkFloat(t, "L2CACodePhaseBiasM", *m.L2CACodePhaseBiasM, -86.42, 1e-9)
}
