package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type Msg1011Satellite struct {
	GLONASSL1Observation
	GLONASSL2Observation
}

type Msg1011 struct {
	GLONASSObservationHeader
	Satellites []Msg1011Satellite
}

func DecodeMsg1011(payload []byte) (*Msg1011, error) {
	r := core.NewBitReader(payload)
	h, numSats := decodeGLONASSObservationHeader(r)
	sats := make([]Msg1011Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1011Satellite{
			GLONASSL1Observation: decodeGLONASSL1Observation(r),
			GLONASSL2Observation: decodeGLONASSL2Observation(r),
		}
	}
	return &Msg1011{GLONASSObservationHeader: h, Satellites: sats}, nil
}
