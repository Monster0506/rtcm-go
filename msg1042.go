package rtcm

type Msg1042 struct {
	MessageType             int
	SatelliteID             int
	WeekNumber              int
	SVURAI                  int
	IDOTSemiCirclesPerS     float64
	AODE                    int
	TocS                    float64
	A2SPerS2                float64
	A1SPerS                 float64
	A0S                     float64
	AODC                    int
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
	Tgd1S                   float64
	Tgd2S                   float64
	SVHealth                bool
}

func DecodeMsg1042(payload []byte) (*Msg1042, error) {
	r := NewBitReader(payload)
	m := &Msg1042{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(6))
	m.WeekNumber = int(r.ReadUint(13))
	m.SVURAI = int(r.ReadUint(4))
	m.IDOTSemiCirclesPerS = float64(r.ReadInt(14)) * twoPow(-43)
	m.AODE = int(r.ReadUint(5))
	m.TocS = float64(r.ReadUint(17)) * 8
	m.A2SPerS2 = float64(r.ReadInt(11)) * twoPow(-66)
	m.A1SPerS = float64(r.ReadInt(22)) * twoPow(-50)
	m.A0S = float64(r.ReadInt(24)) * twoPow(-33)
	m.AODC = int(r.ReadUint(5))
	m.CrsM = float64(r.ReadInt(18)) * twoPow(-6)
	m.DeltaNSemiCirclesPerS = float64(r.ReadInt(16)) * twoPow(-43)
	m.M0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CucRad = float64(r.ReadInt(18)) * twoPow(-31)
	m.Eccentricity = float64(r.ReadUint(32)) * twoPow(-33)
	m.CusRad = float64(r.ReadInt(18)) * twoPow(-31)
	m.SqrtA = float64(r.ReadUint(32)) * twoPow(-19)
	m.ToeS = float64(r.ReadUint(17)) * 8
	m.CicRad = float64(r.ReadInt(18)) * twoPow(-31)
	m.Omega0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CisRad = float64(r.ReadInt(18)) * twoPow(-31)
	m.I0SemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.CrcM = float64(r.ReadInt(18)) * twoPow(-6)
	m.OmegaSemiCircles = float64(r.ReadInt(32)) * twoPow(-31)
	m.OmegaDotSemiCirclesPerS = float64(r.ReadInt(24)) * twoPow(-43)
	m.Tgd1S = float64(r.ReadInt(10)) * 0.1e-9
	m.Tgd2S = float64(r.ReadInt(10)) * 0.1e-9
	m.SVHealth = r.ReadUint(1) != 0
	return m, nil
}
