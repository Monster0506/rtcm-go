package rtcm



type Msg1059 struct {
	SSRHeader
	SatelliteBiases []SSRSatelliteCodeBiases
}

func DecodeMsg1059(payload []byte) (*Msg1059, error) {
	r := NewBitReader(payload)
	h := decodeSSRHeader(r, 20)
	m := &Msg1059{SSRHeader: h}
	m.SatelliteBiases = make([]SSRSatelliteCodeBiases, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.SatelliteBiases[i] = decodeSSRSatelliteCodeBiases(r, 6)
	}
	return m, nil
}
