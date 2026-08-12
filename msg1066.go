package rtcm



type Msg1066 struct {
	SSROrbitHeader
	Corrections []SSRCombinedCorrection
}

func DecodeMsg1066(payload []byte) (*Msg1066, error) {
	r := NewBitReader(payload)
	h := decodeSSROrbitHeader(r, 17)
	m := &Msg1066{SSROrbitHeader: h}
	m.Corrections = make([]SSRCombinedCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRCombinedCorrection(r, 5)
	}
	return m, nil
}
