package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type Msg1009Satellite struct {
	GLONASSL1Observation
}

type Msg1009 struct {
	GLONASSObservationHeader
	Satellites []Msg1009Satellite
}

func DecodeMsg1009(payload []byte) (*Msg1009, error) {
	r := core.NewBitReader(payload)
	h, numSats := decodeGLONASSObservationHeader(r)
	sats := make([]Msg1009Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1009Satellite{GLONASSL1Observation: decodeGLONASSL1Observation(r)}
	}
	return &Msg1009{GLONASSObservationHeader: h, Satellites: sats}, nil
}
