package rtcm

type Msg1002Satellite struct {
	GPSL1ExtendedObservation
}

type Msg1002 struct {
	GPSObservationHeader
	Satellites []Msg1002Satellite
}

func DecodeMsg1002(payload []byte) (*Msg1002, error) {
	r := NewBitReader(payload)
	h, numSats := decodeGPSObservationHeader(r)
	sats := make([]Msg1002Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1002Satellite{GPSL1ExtendedObservation: decodeGPSL1ExtendedObservation(r)}
	}
	return &Msg1002{GPSObservationHeader: h, Satellites: sats}, nil
}
