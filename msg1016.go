package rtcm



type Msg1016Satellite struct {
	SatelliteID              int
	AmbiguityStatusFlag      int
	NonSyncCount             int
	GeometricCorrectionDiffM float64
	IODE                     int
}



type Msg1016 struct {
	NetworkRTKHeader
	Satellites []Msg1016Satellite
}

func DecodeMsg1016(payload []byte) (*Msg1016, error) {
	r := NewBitReader(payload)
	h := decodeNetworkRTKHeader(r)
	m := &Msg1016{NetworkRTKHeader: h}
	m.Satellites = make([]Msg1016Satellite, h.NumGPSSats)
	for i := 0; i < h.NumGPSSats; i++ {
		m.Satellites[i] = Msg1016Satellite{
			SatelliteID:              int(r.ReadUint(6)),
			AmbiguityStatusFlag:      int(r.ReadUint(2)),
			NonSyncCount:             int(r.ReadUint(3)),
			GeometricCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
			IODE:                     int(r.ReadUint(8)),
		}
	}
	return m, nil
}
