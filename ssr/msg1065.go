package ssr

import "github.com/Monster0506/rtcm-go/core"

type Msg1065 struct {
	SSRHeader
	SatelliteBiases []SSRSatelliteCodeBiases
}

func DecodeMsg1065(payload []byte) (*Msg1065, error) {
	r := core.NewBitReader(payload)
	h := decodeSSRHeader(r, 17)
	m := &Msg1065{SSRHeader: h}
	m.SatelliteBiases = make([]SSRSatelliteCodeBiases, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.SatelliteBiases[i] = decodeSSRSatelliteCodeBiases(r, 5)
	}
	return m, nil
}
