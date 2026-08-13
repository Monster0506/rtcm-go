package ephemeris

import "github.com/Monster0506/rtcm-go/core"

type Msg1044 struct {
	MessageType             int
	SatelliteID             int
	TocS                    float64
	Af2SPerS2               float64
	Af1SPerS                float64
	Af0S                    float64
	IODE                    int
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
	IDOTSemiCirclesPerS     float64
	CodesOnL2               int
	WeekNumber              int
	URA                     int
	SVHealth                int
	TgdS                    float64
	IODC                    int
	FitInterval             bool
}

func DecodeMsg1044(payload []byte) (*Msg1044, error) {
	r := core.NewBitReader(payload)
	m := &Msg1044{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(4))
	m.TocS = float64(r.ReadUint(16)) * 16
	m.Af2SPerS2 = float64(r.ReadInt(8)) * core.TwoPow(-55)
	m.Af1SPerS = float64(r.ReadInt(16)) * core.TwoPow(-43)
	m.Af0S = float64(r.ReadInt(22)) * core.TwoPow(-31)
	m.IODE = int(r.ReadUint(8))
	m.CrsM = float64(r.ReadInt(16)) * core.TwoPow(-5)
	m.DeltaNSemiCirclesPerS = float64(r.ReadInt(16)) * core.TwoPow(-43)
	m.M0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	m.CucRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	m.Eccentricity = float64(r.ReadUint(32)) * core.TwoPow(-33)
	m.CusRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	m.SqrtA = float64(r.ReadUint(32)) * core.TwoPow(-19)
	m.ToeS = float64(r.ReadUint(16)) * 16
	m.CicRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	m.Omega0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	m.CisRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	m.I0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	m.CrcM = float64(r.ReadInt(16)) * core.TwoPow(-5)
	m.OmegaSemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	m.OmegaDotSemiCirclesPerS = float64(r.ReadInt(24)) * core.TwoPow(-43)
	m.IDOTSemiCirclesPerS = float64(r.ReadInt(14)) * core.TwoPow(-43)
	m.CodesOnL2 = int(r.ReadUint(2))
	m.WeekNumber = int(r.ReadUint(10))
	m.URA = int(r.ReadUint(4))
	m.SVHealth = int(r.ReadUint(6))
	m.TgdS = float64(r.ReadInt(8)) * core.TwoPow(-31)
	m.IODC = int(r.ReadUint(10))
	m.FitInterval = r.ReadUint(1) != 0
	return m, nil
}
