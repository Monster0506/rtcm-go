package rtcm

type Msg1006 struct {
	Msg1005
	AntennaHeightM float64
}

func DecodeMsg1006(payload []byte) (*Msg1006, error) {
	r := NewBitReader(payload)
	m1005 := decodeMsg1005Fields(r)
	height := float64(r.ReadUint(16)) * 0.0001
	return &Msg1006{Msg1005: m1005, AntennaHeightM: height}, nil
}
