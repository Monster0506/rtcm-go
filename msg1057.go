package rtcm



type Msg1057 struct {
	SSROrbitHeader
	Corrections []SSROrbitCorrection
}

func DecodeMsg1057(payload []byte) (*Msg1057, error) {
	r := NewBitReader(payload)
	h := decodeSSROrbitHeader(r, 20)
	m := &Msg1057{SSROrbitHeader: h}
	m.Corrections = make([]SSROrbitCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSROrbitCorrection(r, 6)
	}
	return m, nil
}
