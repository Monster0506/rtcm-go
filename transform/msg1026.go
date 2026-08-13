package transform

import "github.com/Monster0506/rtcm-go/core"

type Msg1026 struct {
	MessageType                    int
	SystemIdentificationNumber     int
	ProjectionType                 int
	LatitudeOfFalseOriginDeg       float64
	LongitudeOfFalseOriginDeg      float64
	LatitudeOfStandardParallel1Deg float64
	LatitudeOfStandardParallel2Deg float64
	EastingOfFalseOriginM          float64
	NorthingOfFalseOriginM         float64
}

func DecodeMsg1026(payload []byte) (*Msg1026, error) {
	r := core.NewBitReader(payload)
	m := &Msg1026{}
	m.MessageType = int(r.ReadUint(12))
	m.SystemIdentificationNumber = int(r.ReadUint(8))
	m.ProjectionType = int(r.ReadUint(6))
	m.LatitudeOfFalseOriginDeg = float64(r.ReadInt(34)) * 0.000000011
	m.LongitudeOfFalseOriginDeg = float64(r.ReadInt(35)) * 0.000000011
	m.LatitudeOfStandardParallel1Deg = float64(r.ReadInt(34)) * 0.000000011
	m.LatitudeOfStandardParallel2Deg = float64(r.ReadInt(34)) * 0.000000011
	m.EastingOfFalseOriginM = float64(r.ReadUint(36)) * 0.001
	m.NorthingOfFalseOriginM = float64(r.ReadInt(35)) * 0.001
	return m, nil
}
