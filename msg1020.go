package rtcm

// Msg1020 holds the raw ICD field values from a GLONASS ephemeris message
// (Message Type 1020). Values are NOT scaled to physical units. Signed
// fields use GLONASS ICD sign-magnitude encoding (ReadSignMagnitude), not
// two's complement -- see RTCM10403.3's "intS" data type note.
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
	Tb                           int
	XnFirstDerivative            int64
	Xn                           int64
	XnSecondDerivative           int64
	YnFirstDerivative            int64
	Yn                           int64
	YnSecondDerivative           int64
	ZnFirstDerivative            int64
	Zn                           int64
	ZnSecondDerivative           int64
	P3                           bool
	GammaN                       int64
	MP                           int
	MLn3rdString                 bool
	TauN                         int64
	MDeltaTau                    int64
	En                           int
	MP4                          bool
	MFT                          int
	MNT                          int
	MM                           int
	AvailabilityOfAdditionalData bool
	NA                           int
	TauC                         int64
	MN4                          int
	MTauGPS                      int64
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
	m.Tb = int(r.ReadUint(7))
	m.XnFirstDerivative = r.ReadSignMagnitude(24)
	m.Xn = r.ReadSignMagnitude(27)
	m.XnSecondDerivative = r.ReadSignMagnitude(5)
	m.YnFirstDerivative = r.ReadSignMagnitude(24)
	m.Yn = r.ReadSignMagnitude(27)
	m.YnSecondDerivative = r.ReadSignMagnitude(5)
	m.ZnFirstDerivative = r.ReadSignMagnitude(24)
	m.Zn = r.ReadSignMagnitude(27)
	m.ZnSecondDerivative = r.ReadSignMagnitude(5)
	m.P3 = r.ReadUint(1) != 0
	m.GammaN = r.ReadSignMagnitude(11)
	m.MP = int(r.ReadUint(2))
	m.MLn3rdString = r.ReadUint(1) != 0
	m.TauN = r.ReadSignMagnitude(22)
	m.MDeltaTau = r.ReadSignMagnitude(5)
	m.En = int(r.ReadUint(5))
	m.MP4 = r.ReadUint(1) != 0
	m.MFT = int(r.ReadUint(4))
	m.MNT = int(r.ReadUint(11))
	m.MM = int(r.ReadUint(2))
	m.AvailabilityOfAdditionalData = r.ReadUint(1) != 0
	m.NA = int(r.ReadUint(11))
	m.TauC = r.ReadSignMagnitude(32)
	m.MN4 = int(r.ReadUint(5))
	m.MTauGPS = r.ReadSignMagnitude(22)
	m.MLn5thString = r.ReadUint(1) != 0
	r.ReadUint(7) // reserved
	return m, nil
}
