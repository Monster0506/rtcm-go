package transform

import "github.com/Monster0506/rtcm-go/core"

type Msg1027 struct {
	MessageType                     int
	SystemIdentificationNumber      int
	ProjectionType                  int
	RectificationFlag               bool
	LatitudeOfProjectionCenterDeg   float64
	LongitudeOfProjectionCenterDeg  float64
	AzimuthOfInitialLineDeg         float64
	DiffAngleRectifiedToSkewGridDeg float64
	AddScaleFactorPPM               float64
	EastingAtProjectionCenterM      float64
	NorthingAtProjectionCenterM     float64
}

func DecodeMsg1027(payload []byte) (*Msg1027, error) {
	r := core.NewBitReader(payload)
	m := &Msg1027{}
	m.MessageType = int(r.ReadUint(12))
	m.SystemIdentificationNumber = int(r.ReadUint(8))
	m.ProjectionType = int(r.ReadUint(6))
	m.RectificationFlag = r.ReadUint(1) != 0
	m.LatitudeOfProjectionCenterDeg = float64(r.ReadInt(34)) * 0.000000011
	m.LongitudeOfProjectionCenterDeg = float64(r.ReadInt(35)) * 0.000000011
	m.AzimuthOfInitialLineDeg = float64(r.ReadUint(35)) * 0.000000011
	m.DiffAngleRectifiedToSkewGridDeg = float64(r.ReadInt(26)) * 0.000000011
	m.AddScaleFactorPPM = float64(r.ReadUint(30)) * 0.00001
	m.EastingAtProjectionCenterM = float64(r.ReadUint(36)) * 0.001
	m.NorthingAtProjectionCenterM = float64(r.ReadInt(35)) * 0.001
	return m, nil
}
