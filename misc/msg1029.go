package misc

import "github.com/Monster0506/rtcm-go/core"

type Msg1029 struct {
	MessageType     int
	StationID       int
	MJDNumber       int
	SecondsOfDayUTC int
	NumCharacters   int
	Text            string
}

func DecodeMsg1029(payload []byte) (*Msg1029, error) {
	r := core.NewBitReader(payload)
	m := &Msg1029{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.MJDNumber = int(r.ReadUint(16))
	m.SecondsOfDayUTC = int(r.ReadUint(17))
	m.NumCharacters = int(r.ReadUint(7))
	m.Text = core.ReadLengthPrefixedString(r)
	return m, nil
}
