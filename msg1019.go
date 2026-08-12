package rtcm

type Msg1019 struct {
	MessageType             int
	SatelliteID             int
	WeekNumber              int
	SVAccuracy              int
	CodeOnL2                int
	IDOTSemiCirclesPerS     float64
	IODE                    int
	TocS                    float64
	Af2SPerS2               float64
	Af1SPerS                float64
	Af0S                    float64
	IODC                    int
	CrsM                    float64
	DeltaNSemiCirclesPerS   float64
	M0SemiCircles           float64
	CucRad                  float64
	Eccentricity            float64
	CusRad                  float64
	SqrtA                   float64
	ToeS                    float64
	CicRad                  float64
	Omega0SemiCircles       float64
	CisRad                  float64
	I0SemiCircles           float64
	CrcM                    float64
	OmegaSemiCircles        float64
	OmegaDotSemiCirclesPerS float64
	TgdS                    float64
	SVHealth                int
	L2PDataFlag             bool
	FitInterval             bool
}

func DecodeMsg1019(payload []byte) (*Msg1019, error) {
	r := NewBitReader(payload)
	m := &Msg1019{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(6))
	m.WeekNumber = int(r.ReadUint(10))
	m.SVAccuracy = int(r.ReadUint(4))
	m.CodeOnL2 = int(r.ReadUint(2))
	m.IDOTSemiCirclesPerS = float64(r.ReadInt(14)) * twoPow(-43)
	m.IODE = int(r.ReadUint(8))
	m.TocS = float64(r.ReadUint(16)) * 16
	m.Af2SPerS2 = float64(r.ReadInt(8)) * twoPow(-55)
	m.Af1SPerS = float64(r.ReadInt(16)) * twoPow(-43)
	m.Af0S = float64(r.ReadInt(22)) * twoPow(-31)
	m.IODC = int(r.ReadUint(10))
	m.CrsM = float64(r.ReadInt(16)) * twoPow(-5)
	m.DeltaNSemiCirclesPerS = float64(r.ReadInt(16)) * twoPow(-43)
	m.M0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CucRad = float64(r.ReadInt(16)) * twoPow(-29)
	m.Eccentricity = float64(r.ReadUint(32)) * twoPow(-33)
	m.CusRad = float64(r.ReadInt(16)) * twoPow(-29)
	m.SqrtA = float64(r.ReadUint(32)) * twoPow(-19)
	m.ToeS = float64(r.ReadUint(16)) * 16
	m.CicRad = float64(r.ReadInt(16)) * twoPow(-29)
	m.Omega0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CisRad = float64(r.ReadInt(16)) * twoPow(-29)
	m.I0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CrcM = float64(r.ReadInt(16)) * twoPow(-5)
	m.OmegaSemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.OmegaDotSemiCirclesPerS = float64(r.ReadInt(24)) * twoPow(-43)
	m.TgdS = float64(r.ReadInt(8)) * twoPow(-31)
	m.SVHealth = int(r.ReadUint(6))
	m.L2PDataFlag = r.ReadUint(1) != 0
	m.FitInterval = r.ReadUint(1) != 0
	return m, nil
}
