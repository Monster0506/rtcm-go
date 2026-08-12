package rtcm





type Msg1031Satellite struct {
	SatelliteID int
	SigmaOcMM   float64
	SigmaOdPPM  float64
	SigmaOhPPM  float64
	SigmaIcMM   float64
	SigmaIdPPM  float64
}



type Msg1031 struct {
	MessageType                   int
	GLONASSEpochTimeS             float64
	StationID                     int
	NRefs                         int
	NumGLONASSSatSignalsProcessed int
	Satellites                    []Msg1031Satellite
}

func DecodeMsg1031(payload []byte) (*Msg1031, error) {
	r := NewBitReader(payload)
	m := &Msg1031{}
	m.MessageType = int(r.ReadUint(12))
	m.GLONASSEpochTimeS = float64(r.ReadUint(17)) * 1
	m.StationID = int(r.ReadUint(12))
	m.NRefs = int(r.ReadUint(7))
	m.NumGLONASSSatSignalsProcessed = int(r.ReadUint(5))
	m.Satellites = make([]Msg1031Satellite, m.NumGLONASSSatSignalsProcessed)
	for i := 0; i < m.NumGLONASSSatSignalsProcessed; i++ {
		m.Satellites[i] = Msg1031Satellite{
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
