package rtcm



type Msg1014 struct {
	MessageType                 int
	NetworkID                   int
	SubnetworkID                int
	NumAuxStations              int
	MasterReferenceStationID    int
	AuxiliaryReferenceStationID int
	DeltaLatitudeDeg            float64
	DeltaLongitudeDeg           float64
	DeltaHeightM                float64
}

func DecodeMsg1014(payload []byte) (*Msg1014, error) {
	r := NewBitReader(payload)
	m := &Msg1014{}
	m.MessageType = int(r.ReadUint(12))
	m.NetworkID = int(r.ReadUint(8))
	m.SubnetworkID = int(r.ReadUint(4))
	m.NumAuxStations = int(r.ReadUint(5))
	m.MasterReferenceStationID = int(r.ReadUint(12))
	m.AuxiliaryReferenceStationID = int(r.ReadUint(12))
	m.DeltaLatitudeDeg = float64(r.ReadInt(20)) * 25e-6
	m.DeltaLongitudeDeg = float64(r.ReadInt(21)) * 25e-6
	m.DeltaHeightM = float64(r.ReadInt(23)) * 0.001
	return m, nil
}
