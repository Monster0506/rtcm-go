package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1038Satellite struct {
	SatelliteID              int
	AmbiguityStatusFlag      int
	NonSyncCount             int
	GeometricCorrectionDiffM float64
	IOD                      int
}

type Msg1038 struct {
	GLONASSNetworkRTKHeader
	Satellites []Msg1038Satellite
}

func DecodeMsg1038(payload []byte) (*Msg1038, error) {
	r := core.NewBitReader(payload)
	h := decodeGLONASSNetworkRTKHeader(r)
	m := &Msg1038{GLONASSNetworkRTKHeader: h}
	m.Satellites = make([]Msg1038Satellite, h.NumGLONASSDataEntries)
	for i := 0; i < h.NumGLONASSDataEntries; i++ {
		m.Satellites[i] = Msg1038Satellite{
			SatelliteID:              int(r.ReadUint(6)),
			AmbiguityStatusFlag:      int(r.ReadUint(2)),
			NonSyncCount:             int(r.ReadUint(3)),
			GeometricCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
			IOD:                      int(r.ReadUint(8)),
		}
	}
	return m, nil
}
