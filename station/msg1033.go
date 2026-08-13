package station

import "github.com/Monster0506/rtcm-go/core"

type Msg1033 struct {
	Msg1008
	ReceiverType    string
	FirmwareVersion string
	ReceiverSerial  string
}

func DecodeMsg1033(payload []byte) (*Msg1033, error) {
	r := core.NewBitReader(payload)
	m1008 := decodeMsg1008Fields(r)
	receiverType := core.ReadLengthPrefixedString(r)
	firmware := core.ReadLengthPrefixedString(r)
	receiverSerial := core.ReadLengthPrefixedString(r)
	return &Msg1033{
		Msg1008:         m1008,
		ReceiverType:    receiverType,
		FirmwareVersion: firmware,
		ReceiverSerial:  receiverSerial,
	}, nil
}
