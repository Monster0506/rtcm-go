package rtcm



type Msg1007 struct {
	MessageType       int
	StationID         int
	AntennaDescriptor string
	SetupID           int
}

func DecodeMsg1007(payload []byte) (*Msg1007, error) {
	r := NewBitReader(payload)
	m := &Msg1007{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.AntennaDescriptor = readLengthPrefixedString(r)
	m.SetupID = int(r.ReadUint(8))
	return m, nil
}
