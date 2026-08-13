package ssr

import "github.com/Monster0506/rtcm-go/core"

type Msg1068 struct {
	SSRHeader
	Corrections []SSRHighRateClock
}

func DecodeMsg1068(payload []byte) (*Msg1068, error) {
	r := core.NewBitReader(payload)
	h := decodeSSRHeader(r, 17)
	m := &Msg1068{SSRHeader: h}
	m.Corrections = make([]SSRHighRateClock, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRHighRateClock(r, 5)
	}
	return m, nil
}
