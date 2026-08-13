package ssr

import "github.com/Monster0506/rtcm-go/core"

type Msg1062 struct {
	SSRHeader
	Corrections []SSRHighRateClock
}

func DecodeMsg1062(payload []byte) (*Msg1062, error) {
	r := core.NewBitReader(payload)
	h := decodeSSRHeader(r, 20)
	m := &Msg1062{SSRHeader: h}
	m.Corrections = make([]SSRHighRateClock, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.Corrections[i] = decodeSSRHighRateClock(r, 6)
	}
	return m, nil
}
