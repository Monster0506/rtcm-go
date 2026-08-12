package rtcm



type Msg1034Satellite struct {
	SatelliteID                 int
	IODE                        int
	GeometricGradientNorthPPM   float64
	GeometricGradientEastPPM    float64
	IonosphericGradientNorthPPM float64
	IonosphericGradientEastPPM  float64
}



type Msg1034 struct {
	MessageType               int
	StationID                 int
	GPSFKPEpochTimeS          float64
	NumGPSSatSignalsProcessed int
	Satellites                []Msg1034Satellite
}

func DecodeMsg1034(payload []byte) (*Msg1034, error) {
	r := NewBitReader(payload)
	m := &Msg1034{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.GPSFKPEpochTimeS = float64(r.ReadUint(20)) * 1
	m.NumGPSSatSignalsProcessed = int(r.ReadUint(5))
	m.Satellites = make([]Msg1034Satellite, m.NumGPSSatSignalsProcessed)
	for i := 0; i < m.NumGPSSatSignalsProcessed; i++ {
		m.Satellites[i] = Msg1034Satellite{
			SatelliteID:                 int(r.ReadUint(6)),
			IODE:                        int(r.ReadUint(8)),
			GeometricGradientNorthPPM:   float64(r.ReadInt(12)) * 0.01,
			GeometricGradientEastPPM:    float64(r.ReadInt(12)) * 0.01,
			IonosphericGradientNorthPPM: float64(r.ReadInt(14)) * 0.01,
			IonosphericGradientEastPPM:  float64(r.ReadInt(14)) * 0.01,
		}
	}
	return m, nil
}
