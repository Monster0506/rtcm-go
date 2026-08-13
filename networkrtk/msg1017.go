package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1017Satellite struct {
	SatelliteID                int
	AmbiguityStatusFlag        int
	NonSyncCount               int
	GeometricCorrectionDiffM   float64
	IODE                       int
	IonosphericCorrectionDiffM float64
}

type Msg1017 struct {
	NetworkRTKHeader
	Satellites []Msg1017Satellite
}

func DecodeMsg1017(payload []byte) (*Msg1017, error) {
	r := core.NewBitReader(payload)
	h := decodeNetworkRTKHeader(r)
	m := &Msg1017{NetworkRTKHeader: h}
	m.Satellites = make([]Msg1017Satellite, h.NumGPSSats)
	for i := 0; i < h.NumGPSSats; i++ {
		m.Satellites[i] = Msg1017Satellite{
			SatelliteID:                int(r.ReadUint(6)),
			AmbiguityStatusFlag:        int(r.ReadUint(2)),
			NonSyncCount:               int(r.ReadUint(3)),
			GeometricCorrectionDiffM:   float64(r.ReadInt(17)) * 0.0005,
			IODE:                       int(r.ReadUint(8)),
			IonosphericCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
		}
	}
	return m, nil
}
