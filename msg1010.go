package rtcm

type Msg1010Satellite struct {
	GLONASSL1ExtendedObservation
}

type Msg1010 struct {
	GLONASSObservationHeader
	Satellites []Msg1010Satellite
}

func DecodeMsg1010(payload []byte) (*Msg1010, error) {
	r := NewBitReader(payload)
	h, numSats := decodeGLONASSObservationHeader(r)
	sats := make([]Msg1010Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1010Satellite{GLONASSL1ExtendedObservation: decodeGLONASSL1ExtendedObservation(r)}
	}
	return &Msg1010{GLONASSObservationHeader: h, Satellites: sats}, nil
}
