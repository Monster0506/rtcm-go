package rtcm

type Msg1020 struct {
	MessageType                  int
	SatelliteID                  int
	FrequencyChannelNumber       int
	AlmanacHealth                bool
	AlmanacHealthAvailability    bool
	P1                           int
	Tk                           int
	MSbOfBn                      bool
	P2                           bool
	TbMin                        float64
	XnFirstDerivativeKmPerS      float64
	XnKm                         float64
	XnSecondDerivativeKmPerS2    float64
	YnFirstDerivativeKmPerS      float64
	YnKm                         float64
	YnSecondDerivativeKmPerS2    float64
	ZnFirstDerivativeKmPerS      float64
	ZnKm                         float64
	ZnSecondDerivativeKmPerS2    float64
	P3                           bool
	GammaN                       float64
	MP                           int
	MLn3rdString                 bool
	TauNS                        float64
	MDeltaTauS                   float64
	En                           int
	MP4                          bool
	MFT                          int
	MNT                          int
	MM                           int
	AvailabilityOfAdditionalData bool
	NA                           int
	TauCS                        float64
	MN4                          int
	MTauGPSS                     float64
	MLn5thString                 bool
}

func DecodeMsg1020(payload []byte) (*Msg1020, error) {
	r := NewBitReader(payload)
	m := &Msg1020{}
	m.MessageType = int(r.ReadUint(12))
	m.SatelliteID = int(r.ReadUint(6))
	m.FrequencyChannelNumber = int(r.ReadUint(5))
	m.AlmanacHealth = r.ReadUint(1) != 0
	m.AlmanacHealthAvailability = r.ReadUint(1) != 0
	m.P1 = int(r.ReadUint(2))
	m.Tk = int(r.ReadUint(12))
	m.MSbOfBn = r.ReadUint(1) != 0
	m.P2 = r.ReadUint(1) != 0
	m.TbMin = float64(r.ReadUint(7)) * 15
	m.XnFirstDerivativeKmPerS = float64(r.ReadSignMagnitude(24)) * twoPow(-20)
	m.XnKm = float64(r.ReadSignMagnitude(27)) * twoPow(-11)
	m.XnSecondDerivativeKmPerS2 = float64(r.ReadSignMagnitude(5)) * twoPow(-30)
	m.YnFirstDerivativeKmPerS = float64(r.ReadSignMagnitude(24)) * twoPow(-20)
	m.YnKm = float64(r.ReadSignMagnitude(27)) * twoPow(-11)
	m.YnSecondDerivativeKmPerS2 = float64(r.ReadSignMagnitude(5)) * twoPow(-30)
	m.ZnFirstDerivativeKmPerS = float64(r.ReadSignMagnitude(24)) * twoPow(-20)
	m.ZnKm = float64(r.ReadSignMagnitude(27)) * twoPow(-11)
	m.ZnSecondDerivativeKmPerS2 = float64(r.ReadSignMagnitude(5)) * twoPow(-30)
	m.P3 = r.ReadUint(1) != 0
	m.GammaN = float64(r.ReadSignMagnitude(11)) * twoPow(-40)
	m.MP = int(r.ReadUint(2))
	m.MLn3rdString = r.ReadUint(1) != 0
	m.TauNS = float64(r.ReadSignMagnitude(22)) * twoPow(-30)
	m.MDeltaTauS = float64(r.ReadSignMagnitude(5)) * twoPow(-30)
	m.En = int(r.ReadUint(5))
	m.MP4 = r.ReadUint(1) != 0
	m.MFT = int(r.ReadUint(4))
	m.MNT = int(r.ReadUint(11))
	m.MM = int(r.ReadUint(2))
	m.AvailabilityOfAdditionalData = r.ReadUint(1) != 0
	m.NA = int(r.ReadUint(11))
	m.TauCS = float64(r.ReadSignMagnitude(32)) * twoPow(-31)
	m.MN4 = int(r.ReadUint(5))
	m.MTauGPSS = float64(r.ReadSignMagnitude(22)) * twoPow(-30)
	m.MLn5thString = r.ReadUint(1) != 0
	r.ReadUint(7) // reserved
	return m, nil
}
