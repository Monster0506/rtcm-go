package rtcm

type Msg1004Satellite struct {
	GPSL1ExtendedObservation
	GPSL2ExtendedObservation
}

type Msg1004 struct {
	GPSObservationHeader
	Satellites []Msg1004Satellite
}

func DecodeMsg1004(payload []byte) (*Msg1004, error) {
	r := NewBitReader(payload)
	h, numSats := decodeGPSObservationHeader(r)
	sats := make([]Msg1004Satellite, numSats)
	for i := range sats {
		sats[i] = Msg1004Satellite{
			GPSL1ExtendedObservation: decodeGPSL1ExtendedObservation(r),
			GPSL2ExtendedObservation: decodeGPSL2ExtendedObservation(r),
		}
	}
	return &Msg1004{GPSObservationHeader: h, Satellites: sats}, nil
}
