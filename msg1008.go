package rtcm

type Msg1008 struct {
	MessageType       int
	StationID         int
	AntennaDescriptor string
	SetupID           int
	AntennaSerial     string
}

func DecodeMsg1008(payload []byte) (*Msg1008, error) {
	r := NewBitReader(payload)
	m := &Msg1008{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.AntennaDescriptor = readLengthPrefixedString(r)
	m.SetupID = int(r.ReadUint(8))
	m.AntennaSerial = readLengthPrefixedString(r)
	return m, nil
}
