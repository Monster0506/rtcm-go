package rtcm



type Msg1032 struct {
	MessageType                   int
	NonPhysicalReferenceStationID int
	PhysicalReferenceStationID    int
	ITRFEpochYear                 int
	ECEFXM                        float64
	ECEFYM                        float64
	ECEFZM                        float64
}

func DecodeMsg1032(payload []byte) (*Msg1032, error) {
	r := NewBitReader(payload)
	m := &Msg1032{}
	m.MessageType = int(r.ReadUint(12))
	m.NonPhysicalReferenceStationID = int(r.ReadUint(12))
	m.PhysicalReferenceStationID = int(r.ReadUint(12))
	m.ITRFEpochYear = int(r.ReadUint(6))
	m.ECEFXM = float64(r.ReadBits38()) * 0.0001
	m.ECEFYM = float64(r.ReadBits38()) * 0.0001
	m.ECEFZM = float64(r.ReadBits38()) * 0.0001
	return m, nil
}
