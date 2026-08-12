package rtcm



type Msg1060 struct {
	SSROrbitHeader
	Corrections []SSRCombinedCorrection
}

func DecodeMsg1060(payload []byte) (*Msg1060, error) {
	r := NewBitReader(payload)
	h := decodeSSROrbitHeader(r, 20)
	m := &Msg1060{SSROrbitHeader: h}
	m.Corrections = make([]SSRCombinedCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRCombinedCorrection(r, 6)
	}
	return m, nil
}
