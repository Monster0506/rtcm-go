package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type Msg1003Satellite struct {
	GPSL1Observation
	GPSL2Observation
}

type Msg1003 struct {
	GPSObservationHeader
	Satellites []Msg1003Satellite
}

func DecodeMsg1003(payload []byte) (*Msg1003, error) {
	r := core.NewBitReader(payload)
	h, numSats := decodeGPSObservationHeader(r)
	sats := make([]Msg1003Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1003Satellite{
			GPSL1Observation: decodeGPSL1Observation(r),
			GPSL2Observation: decodeGPSL2Observation(r),
		}
	}
	return &Msg1003{GPSObservationHeader: h, Satellites: sats}, nil
}
