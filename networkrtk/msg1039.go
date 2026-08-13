package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1039Satellite struct {
	SatelliteID                int
	AmbiguityStatusFlag        int
	NonSyncCount               int
	GeometricCorrectionDiffM   float64
	IOD                        int
	IonosphericCorrectionDiffM float64
}

type Msg1039 struct {
	GLONASSNetworkRTKHeader
	Satellites []Msg1039Satellite
}

func DecodeMsg1039(payload []byte) (*Msg1039, error) {
	r := core.NewBitReader(payload)
	h := decodeGLONASSNetworkRTKHeader(r)
	m := &Msg1039{GLONASSNetworkRTKHeader: h}
	m.Satellites = make([]Msg1039Satellite, h.NumGLONASSDataEntries)
	for i := 0; i < h.NumGLONASSDataEntries; i++ {
		m.Satellites[i] = Msg1039Satellite{
			SatelliteID:                int(r.ReadUint(6)),
			AmbiguityStatusFlag:        int(r.ReadUint(2)),
			NonSyncCount:               int(r.ReadUint(3)),
			GeometricCorrectionDiffM:   float64(r.ReadInt(17)) * 0.0005,
			IOD:                        int(r.ReadUint(8)),
			IonosphericCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
		}
	}
	return m, nil
}
