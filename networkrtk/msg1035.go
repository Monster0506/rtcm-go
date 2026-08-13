package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1035Satellite struct {
	SatelliteID                 int
	IOD                         int
	GeometricGradientNorthPPM   float64
	GeometricGradientEastPPM    float64
	IonosphericGradientNorthPPM float64
	IonosphericGradientEastPPM  float64
}

type Msg1035 struct {
	MessageType                   int
	StationID                     int
	GLONASSFKPEpochTimeS          float64
	NumGLONASSSatSignalsProcessed int
	Satellites                    []Msg1035Satellite
}

func DecodeMsg1035(payload []byte) (*Msg1035, error) {
	r := core.NewBitReader(payload)
	m := &Msg1035{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.GLONASSFKPEpochTimeS = float64(r.ReadUint(17)) * 1
	m.NumGLONASSSatSignalsProcessed = int(r.ReadUint(5))
	m.Satellites = make([]Msg1035Satellite, m.NumGLONASSSatSignalsProcessed)
	for i := 0; i < m.NumGLONASSSatSignalsProcessed; i++ {
		m.Satellites[i] = Msg1035Satellite{
			SatelliteID:                 int(r.ReadUint(6)),
			IOD:                         int(r.ReadUint(8)),
			GeometricGradientNorthPPM:   float64(r.ReadInt(12)) * 0.01,
			GeometricGradientEastPPM:    float64(r.ReadInt(12)) * 0.01,
			IonosphericGradientNorthPPM: float64(r.ReadInt(14)) * 0.01,
			IonosphericGradientEastPPM:  float64(r.ReadInt(14)) * 0.01,
		}
	}
	return m, nil
}
