package rtcm





type Msg1030Satellite struct {
	SatelliteID int
	SigmaOcMM   float64
	SigmaOdPPM  float64
	SigmaOhPPM  float64
	SigmaIcMM   float64
	SigmaIdPPM  float64
}



type Msg1030 struct {
	MessageType               int
	GPSEpochTimeS             float64
	StationID                 int
	NRefs                     int
	NumGPSSatSignalsProcessed int
	Satellites                []Msg1030Satellite
}

func DecodeMsg1030(payload []byte) (*Msg1030, error) {
	r := NewBitReader(payload)
	m := &Msg1030{}
	m.MessageType = int(r.ReadUint(12))
	m.GPSEpochTimeS = float64(r.ReadUint(20)) * 1
	m.StationID = int(r.ReadUint(12))
	m.NRefs = int(r.ReadUint(7))
	m.NumGPSSatSignalsProcessed = int(r.ReadUint(5))
	m.Satellites = make([]Msg1030Satellite, m.NumGPSSatSignalsProcessed)
	for i := 0; i < m.NumGPSSatSignalsProcessed; i++ {
		m.Satellites[i] = Msg1030Satellite{
			SatelliteID: int(r.ReadUint(6)),
			SigmaOcMM:   float64(r.ReadUint(8)) * 0.5,
			SigmaOdPPM:  float64(r.ReadUint(9)) * 0.01,
			SigmaOhPPM:  float64(r.ReadUint(6)) * 0.1,
			SigmaIcMM:   float64(r.ReadUint(10)) * 0.5,
			SigmaIdPPM:  float64(r.ReadUint(10)) * 0.01,
		}
	}
	return m, nil
}
