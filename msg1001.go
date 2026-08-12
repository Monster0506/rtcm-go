package rtcm

type Msg1001Satellite struct {
	GPSL1Observation
}

type Msg1001 struct {
	GPSObservationHeader
	Satellites []Msg1001Satellite
}

func DecodeMsg1001(payload []byte) (*Msg1001, error) {
	r := NewBitReader(payload)
	h, numSats := decodeGPSObservationHeader(r)
	sats := make([]Msg1001Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1001Satellite{GPSL1Observation: decodeGPSL1Observation(r)}
	}
	return &Msg1001{GPSObservationHeader: h, Satellites: sats}, nil
}
