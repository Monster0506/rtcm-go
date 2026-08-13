package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1015Satellite struct {
	SatelliteID                int
	AmbiguityStatusFlag        int
	NonSyncCount               int
	IonosphericCorrectionDiffM float64
}

type Msg1015 struct {
	NetworkRTKHeader
	Satellites []Msg1015Satellite
}

func DecodeMsg1015(payload []byte) (*Msg1015, error) {
	r := core.NewBitReader(payload)
	h := decodeNetworkRTKHeader(r)
	m := &Msg1015{NetworkRTKHeader: h}
	m.Satellites = make([]Msg1015Satellite, h.NumGPSSats)
	for i := 0; i < h.NumGPSSats; i++ {
		m.Satellites[i] = Msg1015Satellite{
			SatelliteID:                int(r.ReadUint(6)),
			AmbiguityStatusFlag:        int(r.ReadUint(2)),
			NonSyncCount:               int(r.ReadUint(3)),
			IonosphericCorrectionDiffM: float64(r.ReadInt(17)) * 0.0005,
		}
	}
	return m, nil
}
