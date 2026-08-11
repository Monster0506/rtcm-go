package rtcm

type Msg1033 struct {
	Msg1008
	ReceiverType    string
	FirmwareVersion string
	ReceiverSerial  string
}

func DecodeMsg1033(payload []byte) (*Msg1033, error) {
	r := NewBitReader(payload)
	m1008 := decodeMsg1008Fields(r)
	receiverType := readLengthPrefixedString(r)
	firmware := readLengthPrefixedString(r)
	receiverSerial := readLengthPrefixedString(r)
	return &Msg1033{
		Msg1008:         m1008,
		ReceiverType:    receiverType,
		FirmwareVersion: firmware,
		ReceiverSerial:  receiverSerial,
	}, nil
}
