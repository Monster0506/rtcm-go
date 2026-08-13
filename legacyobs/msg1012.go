package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type Msg1012Satellite struct {
	GLONASSL1ExtendedObservation
	GLONASSL2ExtendedObservation
}

type Msg1012 struct {
	GLONASSObservationHeader
	Satellites []Msg1012Satellite
}

func DecodeMsg1012(payload []byte) (*Msg1012, error) {
	r := core.NewBitReader(payload)
	h, numSats := decodeGLONASSObservationHeader(r)
	sats := make([]Msg1012Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1012Satellite{
			GLONASSL1ExtendedObservation: decodeGLONASSL1ExtendedObservation(r),
			GLONASSL2ExtendedObservation: decodeGLONASSL2ExtendedObservation(r),
		}
	}
	return &Msg1012{GLONASSObservationHeader: h, Satellites: sats}, nil
}
