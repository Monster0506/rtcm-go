package networkrtk

import "github.com/Monster0506/rtcm-go/core"

type Msg1013MessageAnnouncement struct {
	MessageID             int
	SyncFlag              bool
	TransmissionIntervalS float64
}

type Msg1013 struct {
	MessageType     int
	StationID       int
	MJDNumber       int
	SecondsOfDayUTC int
	LeapSeconds     int
	Announcements   []Msg1013MessageAnnouncement
}

func DecodeMsg1013(payload []byte) (*Msg1013, error) {
	r := core.NewBitReader(payload)
	m := &Msg1013{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.MJDNumber = int(r.ReadUint(16))
	m.SecondsOfDayUTC = int(r.ReadUint(17))
	numMessages := int(r.ReadUint(5))
	m.LeapSeconds = int(r.ReadUint(8))
	m.Announcements = make([]Msg1013MessageAnnouncement, numMessages)
	for i := 0; i < numMessages; i++ {
		m.Announcements[i] = Msg1013MessageAnnouncement{
			MessageID:             int(r.ReadUint(12)),
			SyncFlag:              r.ReadUint(1) != 0,
			TransmissionIntervalS: float64(r.ReadUint(16)) * 0.1,
		}
	}
	return m, nil
}
