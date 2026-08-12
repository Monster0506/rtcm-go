package rtcm



type Msg1037Satellite struct {
	SatelliteID                int
	AmbiguityStatusFlag        int
	NonSyncCount               int
	IonosphericCorrectionDiffM float64
}



type Msg1037 struct {
	GLONASSNetworkRTKHeader
	Satellites []Msg1037Satellite
}

func DecodeMsg1037(payload []byte) (*Msg1037, error) {
	r := NewBitReader(payload)
	h := decodeGLONASSNetworkRTKHeader(r)
	m := &Msg1037{GLONASSNetworkRTKHeader: h}
	m.Satellites = make([]Msg1037Satellite, h.NumGLONASSDataEntries)
	for i := 0; i < h.NumGLONASSDataEntries; i++ {
		m.Satellites[i] = Msg1037Satellite{
			SatelliteID:                int(r.ReadUint(6)),
			AmbiguityStatusFlag:        int(r.ReadUint(2)),
			NonSyncCount:               int(r.ReadUint(3)),
			IonosphericCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
		}
	}
	return m, nil
}
