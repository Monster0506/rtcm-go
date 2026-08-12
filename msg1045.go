package rtcm

type Msg1045 struct {
	GalileoEphemerisCommon
	OSHS  int
	OSDVS bool
}

func DecodeMsg1045(payload []byte) (*Msg1045, error) {
	r := NewBitReader(payload)
	c := decodeGalileoEphemerisCommon(r)
	m := &Msg1045{GalileoEphemerisCommon: c}
	m.OSHS = int(r.ReadUint(2))
	m.OSDVS = r.ReadUint(1) != 0
	r.ReadUint(7) // reserved
	return m, nil
}
