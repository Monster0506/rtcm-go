package rtcm



type Msg1063 struct {
	SSROrbitHeader
	Corrections []SSROrbitCorrection
}

func DecodeMsg1063(payload []byte) (*Msg1063, error) {
	r := NewBitReader(payload)
	h := decodeSSROrbitHeader(r, 17)
	m := &Msg1063{SSROrbitHeader: h}
	m.Corrections = make([]SSROrbitCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSROrbitCorrection(r, 5)
	}
	return m, nil
}
