package rtcm

type Msg1046 struct {
	GalileoEphemerisCommon
	BGDE5bE1S             float64
	E5bSignalHealthStatus int
	E5bDataValidityStatus bool
	E1BSignalHealthStatus int
	E1BDataValidityStatus bool
}

func DecodeMsg1046(payload []byte) (*Msg1046, error) {
	r := NewBitReader(payload)
	c := decodeGalileoEphemerisCommon(r)
	m := &Msg1046{GalileoEphemerisCommon: c}
	m.BGDE5bE1S = float64(r.ReadInt(10)) * twoPow(-32)
	m.E5bSignalHealthStatus = int(r.ReadUint(2))
	m.E5bDataValidityStatus = r.ReadUint(1) != 0
	m.E1BSignalHealthStatus = int(r.ReadUint(2))
	m.E1BDataValidityStatus = r.ReadUint(1) != 0
	r.ReadUint(2) // reserved
	return m, nil
}
