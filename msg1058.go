package rtcm



type Msg1058 struct {
	SSRHeader
	Corrections []SSRClockCorrection
}

func DecodeMsg1058(payload []byte) (*Msg1058, error) {
	r := NewBitReader(payload)
	h := decodeSSRHeader(r, 20)
	m := &Msg1058{SSRHeader: h}
	m.Corrections = make([]SSRClockCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRClockCorrection(r, 6)
	}
	return m, nil
}
