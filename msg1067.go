package rtcm



type Msg1067 struct {
	SSRHeader
	URAs []SSRURA
}

func DecodeMsg1067(payload []byte) (*Msg1067, error) {
	r := NewBitReader(payload)
	h := decodeSSRHeader(r, 17)
	m := &Msg1067{SSRHeader: h}
	m.URAs = make([]SSRURA, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.URAs[i] = decodeSSRURA(r, 5)
	}
	return m, nil
}
