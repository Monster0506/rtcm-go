package station

import "github.com/Monster0506/rtcm-go/core"

type Msg1008 struct {
	MessageType       int
	StationID         int
	AntennaDescriptor string
	SetupID           int
	AntennaSerial     string
}

func decodeMsg1008Fields(r *core.BitReader) Msg1008 {
	var m Msg1008
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.AntennaDescriptor = core.ReadLengthPrefixedString(r)
	m.SetupID = int(r.ReadUint(8))
	m.AntennaSerial = core.ReadLengthPrefixedString(r)
	return m
}

func DecodeMsg1008(payload []byte) (*Msg1008, error) {
	m := decodeMsg1008Fields(core.NewBitReader(payload))
	return &m, nil
}
