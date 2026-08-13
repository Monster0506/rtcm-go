package transform

import "github.com/Monster0506/rtcm-go/core"

type Msg1023GridPoint struct {
	LatitudeResidualDeg  float64
	LongitudeResidualDeg float64
	HeightResidualM      float64
}

type Msg1023 struct {
	MessageType                    int
	SystemIdentificationNumber     int
	HorizontalShiftIndicator       bool
	VerticalShiftIndicator         bool
	LatitudeOfOriginArcSec         float64
	LongitudeOfOriginArcSec        float64
	NSExtensionArcSec              float64
	EWExtensionArcSec              float64
	MeanLatitudeOffsetDeg          float64
	MeanLongitudeOffsetDeg         float64
	MeanHeightOffsetM              float64
	GridPoints                     [16]Msg1023GridPoint
	HorizontalInterpolationMethod  int
	VerticalInterpolationMethod    int
	HorizontalGridQualityIndicator int
	VerticalGridQualityIndicator   int
	MJDNumber                      int
}

func DecodeMsg1023(payload []byte) (*Msg1023, error) {
	r := core.NewBitReader(payload)
	m := &Msg1023{}
	m.MessageType = int(r.ReadUint(12))
	m.SystemIdentificationNumber = int(r.ReadUint(8))
	m.HorizontalShiftIndicator = r.ReadUint(1) != 0
	m.VerticalShiftIndicator = r.ReadUint(1) != 0
	m.LatitudeOfOriginArcSec = float64(r.ReadInt(21)) * 0.5
	m.LongitudeOfOriginArcSec = float64(r.ReadInt(22)) * 0.5
	m.NSExtensionArcSec = float64(r.ReadUint(12)) * 0.5
	m.EWExtensionArcSec = float64(r.ReadUint(12)) * 0.5
	m.MeanLatitudeOffsetDeg = float64(r.ReadInt(8)) * 0.001
	m.MeanLongitudeOffsetDeg = float64(r.ReadInt(8)) * 0.001
	m.MeanHeightOffsetM = float64(r.ReadInt(15)) * 0.01
	for i := 0; i < 16; i++ {
		m.GridPoints[i] = Msg1023GridPoint{
			LatitudeResidualDeg:  float64(r.ReadInt(9)) * 0.00003,
			LongitudeResidualDeg: float64(r.ReadInt(9)) * 0.00003,
			HeightResidualM:      float64(r.ReadInt(9)) * 0.001,
		}
	}
	m.HorizontalInterpolationMethod = int(r.ReadUint(2))
	m.VerticalInterpolationMethod = int(r.ReadUint(2))
	m.HorizontalGridQualityIndicator = int(r.ReadUint(3))
	m.VerticalGridQualityIndicator = int(r.ReadUint(3))
	m.MJDNumber = int(r.ReadUint(16))
	return m, nil
}
