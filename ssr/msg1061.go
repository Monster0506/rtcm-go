package ssr

import "github.com/Monster0506/rtcm-go/core"

type Msg1061 struct {
	SSRHeader
	URAs []SSRURA
}

func DecodeMsg1061(payload []byte) (*Msg1061, error) {
	r := core.NewBitReader(payload)
	h := decodeSSRHeader(r, 20)
	m := &Msg1061{SSRHeader: h}
	m.URAs = make([]SSRURA, h.NumSatellites)
	for i := 0; i < h.NumSatellites; i++ {
		m.URAs[i] = decodeSSRURA(r, 6)
	}
	return m, nil
}
