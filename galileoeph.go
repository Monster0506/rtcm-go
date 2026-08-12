package rtcm

// GalileoEphemerisCommon holds the ~25 fields shared identically by the
// Galileo F/NAV (1045) and I/NAV (1046) ephemeris messages, up through
// BGDE5aE1 where the two message layouts diverge. Values are NOT scaled
// to physical units.
type GalileoEphemerisCommon struct {
	MessageType  int
	SatelliteID  int
	WeekNumber   int
	IODNav       int
	SISA         int
	IDOT         int64
	Toc          int
	Af2          int64
	Af1          int64
	Af0          int64
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
	BGDE5aE1     int64
}

func decodeGalileoEphemerisCommon(r *BitReader) GalileoEphemerisCommon {
	var c GalileoEphemerisCommon
	c.MessageType = int(r.ReadUint(12))
	c.SatelliteID = int(r.ReadUint(6))
	c.WeekNumber = int(r.ReadUint(12))
	c.IODNav = int(r.ReadUint(10))
	c.SISA = int(r.ReadUint(8))
	c.IDOT = r.ReadInt(14)
	c.Toc = int(r.ReadUint(14))
	c.Af2 = r.ReadInt(6)
	c.Af1 = r.ReadInt(21)
	c.Af0 = r.ReadInt(31)
	c.Crs = r.ReadInt(16)
	c.DeltaN = r.ReadInt(16)
	c.M0 = r.ReadInt(32)
	c.Cuc = r.ReadInt(16)
	c.Eccentricity = r.ReadUint(32)
	c.Cus = r.ReadInt(16)
	c.SqrtA = r.ReadUint(32)
	c.Toe = int(r.ReadUint(14))
	c.Cic = r.ReadInt(16)
	c.Omega0 = r.ReadInt(32)
	c.Cis = r.ReadInt(16)
	c.I0 = r.ReadInt(32)
	c.Crc = r.ReadInt(16)
	c.Omega = r.ReadInt(32)
	c.OmegaDot = r.ReadInt(24)
	c.BGDE5aE1 = r.ReadInt(10)
	return c
}
