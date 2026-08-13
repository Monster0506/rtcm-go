package transform

import "github.com/Monster0506/rtcm-go/core"

type Msg1021 struct {
	CoordTransformCommon
	CoordTransformTail
}

func DecodeMsg1021(payload []byte) (*Msg1021, error) {
	r := core.NewBitReader(payload)
	m := &Msg1021{
		CoordTransformCommon: decodeCoordTransformCommon(r),
	}
	m.CoordTransformTail = decodeCoordTransformTail(r)
	return m, nil
}
