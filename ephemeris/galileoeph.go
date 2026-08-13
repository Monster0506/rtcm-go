package ephemeris

import "github.com/Monster0506/rtcm-go/core"

type GalileoEphemerisCommon struct {
	MessageType             int
	SatelliteID             int
	WeekNumber              int
	IODNav                  int
	SISA                    int
	IDOTSemiCirclesPerS     float64
	TocS                    float64
	Af2SPerS2               float64
	Af1SPerS                float64
	Af0S                    float64
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
	BGDE5aE1S               float64
}

func decodeGalileoEphemerisCommon(r *core.BitReader) GalileoEphemerisCommon {
	var c GalileoEphemerisCommon
	c.MessageType = int(r.ReadUint(12))
	c.SatelliteID = int(r.ReadUint(6))
	c.WeekNumber = int(r.ReadUint(12))
	c.IODNav = int(r.ReadUint(10))
	c.SISA = int(r.ReadUint(8))
	c.IDOTSemiCirclesPerS = float64(r.ReadInt(14)) * core.TwoPow(-43)
	c.TocS = float64(r.ReadUint(14)) * 60
	c.Af2SPerS2 = float64(r.ReadInt(6)) * core.TwoPow(-59)
	c.Af1SPerS = float64(r.ReadInt(21)) * core.TwoPow(-46)
	c.Af0S = float64(r.ReadInt(31)) * core.TwoPow(-34)
	c.CrsM = float64(r.ReadInt(16)) * core.TwoPow(-5)
	c.DeltaNSemiCirclesPerS = float64(r.ReadInt(16)) * core.TwoPow(-43)
	c.M0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	c.CucRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	c.Eccentricity = float64(r.ReadUint(32)) * core.TwoPow(-33)
	c.CusRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	c.SqrtA = float64(r.ReadUint(32)) * core.TwoPow(-19)
	c.ToeS = float64(r.ReadUint(14)) * 60
	c.CicRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	c.Omega0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	c.CisRad = float64(r.ReadInt(16)) * core.TwoPow(-29)
	c.I0SemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	c.CrcM = float64(r.ReadInt(16)) * core.TwoPow(-5)
	c.OmegaSemiCircles = float64(r.ReadInt(32)) * core.TwoPow(-31)
	c.OmegaDotSemiCirclesPerS = float64(r.ReadInt(24)) * core.TwoPow(-43)
	c.BGDE5aE1S = float64(r.ReadInt(10)) * core.TwoPow(-32)
	return c
}
