package rtcm





type Msg1025 struct {
	MessageType                 int
	SystemIdentificationNumber  int
	ProjectionType              int
	LatitudeOfNaturalOriginDeg  float64
	LongitudeOfNaturalOriginDeg float64
	AddScaleFactorPPM           float64
	FalseEastingM               float64
	FalseNorthingM              float64
}

func DecodeMsg1025(payload []byte) (*Msg1025, error) {
	r := NewBitReader(payload)
	m := &Msg1025{}
	m.MessageType = int(r.ReadUint(12))
	m.SystemIdentificationNumber = int(r.ReadUint(8))
	m.ProjectionType = int(r.ReadUint(6))
	m.LatitudeOfNaturalOriginDeg = float64(r.ReadInt(34)) * 0.000000011
	m.LongitudeOfNaturalOriginDeg = float64(r.ReadInt(35)) * 0.000000011
	m.AddScaleFactorPPM = float64(r.ReadUint(30)) * 0.00001
	m.FalseEastingM = float64(r.ReadUint(36)) * 0.001
	m.FalseNorthingM = float64(r.ReadInt(35)) * 0.001
	return m, nil
}
