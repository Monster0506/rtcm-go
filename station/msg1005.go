package station

import "github.com/Monster0506/rtcm-go/core"

type Msg1005 struct {
	MessageType               int
	StationID                 int
	ITRFRealizationYear       int
	GPSIndicator              bool
	GLONASSIndicator          bool
	GalileoIndicator          bool
	ReferenceStationIndicator bool
	ECEFXM                    float64
	OscillatorIndicator       bool
	ECEFYM                    float64
	QuarterCycleIndicator     int
	ECEFZM                    float64
}

func decodeMsg1005Fields(r *core.BitReader) Msg1005 {
	var m Msg1005
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.ITRFRealizationYear = int(r.ReadUint(6))
	m.GPSIndicator = r.ReadUint(1) != 0
	m.GLONASSIndicator = r.ReadUint(1) != 0
	m.GalileoIndicator = r.ReadUint(1) != 0
	m.ReferenceStationIndicator = r.ReadUint(1) != 0
	m.ECEFXM = float64(r.ReadBits38()) * 0.0001
	m.OscillatorIndicator = r.ReadUint(1) != 0
	r.ReadUint(1)
	m.ECEFYM = float64(r.ReadBits38()) * 0.0001
	m.QuarterCycleIndicator = int(r.ReadUint(2))
	m.ECEFZM = float64(r.ReadBits38()) * 0.0001
	return m
}

func DecodeMsg1005(payload []byte) (*Msg1005, error) {
	m := decodeMsg1005Fields(core.NewBitReader(payload))
	return &m, nil
}
