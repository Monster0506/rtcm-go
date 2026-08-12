package rtcm

// Msg1044 holds the raw ICD field values from a QZSS ephemeris message
// (Message Type 1044). Values are NOT scaled to physical units.
type Msg1044 struct {
	MessageType  int
	SatelliteID  int
	Toc          int
	Af2          int64
	Af1          int64
	Af0          int64
	IODE         int
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
	IDOT         int64
	CodesOnL2    int
	WeekNumber   int
	URA          int
	SVHealth     int
	Tgd          int64
	IODC         int
	FitInterval  bool
}

func DecodeMsg1044(payload []byte) (*Msg1044, error) {
	r := NewBitReader(payload)
	m := &Msg1044{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(4))
	m.Toc = int(r.ReadUint(16))
	m.Af2 = r.ReadInt(8)
	m.Af1 = r.ReadInt(16)
	m.Af0 = r.ReadInt(22)
	m.IODE = int(r.ReadUint(8))
	m.Crs = r.ReadInt(16)
	m.DeltaN = r.ReadInt(16)
	m.M0 = r.ReadInt(32)
	m.Cuc = r.ReadInt(16)
	m.Eccentricity = r.ReadUint(32)
	m.Cus = r.ReadInt(16)
	m.SqrtA = r.ReadUint(32)
	m.Toe = int(r.ReadUint(16))
	m.Cic = r.ReadInt(16)
	m.Omega0 = r.ReadInt(32)
	m.Cis = r.ReadInt(16)
	m.I0 = r.ReadInt(32)
	m.Crc = r.ReadInt(16)
	m.Omega = r.ReadInt(32)
	m.OmegaDot = r.ReadInt(24)
	m.IDOT = r.ReadInt(14)
	m.CodesOnL2 = int(r.ReadUint(2))
	m.WeekNumber = int(r.ReadUint(10))
	m.URA = int(r.ReadUint(4))
	m.SVHealth = int(r.ReadUint(6))
	m.Tgd = r.ReadInt(8)
	m.IODC = int(r.ReadUint(10))
	m.FitInterval = r.ReadUint(1) != 0
	return m, nil
}
