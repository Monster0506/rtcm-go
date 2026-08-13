package ssr

import "github.com/Monster0506/rtcm-go/core"

type Msg1064 struct {
	SSRHeader
	Corrections []SSRClockCorrection
}

func DecodeMsg1064(payload []byte) (*Msg1064, error) {
	r := core.NewBitReader(payload)
	h := decodeSSRHeader(r, 17)
	m := &Msg1064{SSRHeader: h}
	m.Corrections = make([]SSRClockCorrection, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRClockCorrection(r, 5)
	}
	return m, nil
}
