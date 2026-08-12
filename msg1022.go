package rtcm





type Msg1022 struct {
	CoordTransformCommon
	XPM float64
	YPM float64
	ZPM float64
	CoordTransformTail
}

func DecodeMsg1022(payload []byte) (*Msg1022, error) {
	r := NewBitReader(payload)
	m := &Msg1022{
		CoordTransformCommon: decodeCoordTransformCommon(r),
	}
	m.XPM = float64(r.ReadInt(35)) * 0.001
	m.YPM = float64(r.ReadInt(35)) * 0.001
	m.ZPM = float64(r.ReadInt(35)) * 0.001
	m.CoordTransformTail = decodeCoordTransformTail(r)
	return m, nil
}
