package rtcm

// Msg1042 holds the raw ICD field values from a BeiDou ephemeris message
// (Message Type 1042). Values are NOT scaled to physical units.
type Msg1042 struct {
	MessageType  int
	SatelliteID  int
	WeekNumber   int
	SVURAI       int
	IDOT         int64
	AODE         int
	Toc          int
	A2           int64
	A1           int64
	A0           int64
	AODC         int
	Crs          int64
	DeltaN       int64
	M0           int64
	Cuc          int64
	Eccentricity uint64
	Cus          int64
	SqrtA        uint64
	Toe          int
	Cic          int64
	Omega0       int64
	Cis          int64
	I0           int64
	Crc          int64
	Omega        int64
	OmegaDot     int64
	Tgd1         int64
	Tgd2         int64
	SVHealth     bool
}

func DecodeMsg1042(payload []byte) (*Msg1042, error) {
	r := NewBitReader(payload)
	m := &Msg1042{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(6))
	m.WeekNumber = int(r.ReadUint(13))
	m.SVURAI = int(r.ReadUint(4))
	m.IDOT = r.ReadInt(14)
	m.AODE = int(r.ReadUint(5))
	m.Toc = int(r.ReadUint(17))
	m.A2 = r.ReadInt(11)
	m.A1 = r.ReadInt(22)
	m.A0 = r.ReadInt(24)
	m.AODC = int(r.ReadUint(5))
	m.Crs = r.ReadInt(18)
	m.DeltaN = r.ReadInt(16)
	m.M0 = r.ReadInt(32)
	m.Cuc = r.ReadInt(18)
	m.Eccentricity = r.ReadUint(32)
	m.Cus = r.ReadInt(18)
	m.SqrtA = r.ReadUint(32)
	m.Toe = int(r.ReadUint(17))
	m.Cic = r.ReadInt(18)
	m.Omega0 = r.ReadInt(32)
	m.Cis = r.ReadInt(18)
	m.I0 = r.ReadInt(32)
	m.Crc = r.ReadInt(18)
	m.Omega = r.ReadInt(32)
	m.OmegaDot = r.ReadInt(24)
	m.Tgd1 = r.ReadInt(10)
	m.Tgd2 = r.ReadInt(10)
	m.SVHealth = r.ReadUint(1) != 0
	return m, nil
}
