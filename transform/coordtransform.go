package transform

import "github.com/Monster0506/rtcm-go/core"

type CoordTransformCommon struct {
	MessageType                            int
	SourceName                             string
	TargetName                             string
	SystemIdentificationNumber             int
	UtilizedTransformationMessageIndicator int
	PlateNumber                            int
	ComputationIndicator                   int
	HeightIndicator                        int
	LatitudeOfOriginArcSec                 float64
	LongitudeOfOriginArcSec                float64
	NSExtensionArcSec                      float64
	EWExtensionArcSec                      float64
	TranslationXM                          float64
	TranslationYM                          float64
	TranslationZM                          float64
	RotationXArcSec                        float64
	RotationYArcSec                        float64
	RotationZArcSec                        float64
	ScaleCorrectionPPM                     float64
}

func decodeCoordTransformCommon(r *core.BitReader) CoordTransformCommon {
	var c CoordTransformCommon
	c.MessageType = int(r.ReadUint(12))
	c.SourceName = core.ReadLengthPrefixedStringN(r, 5)
	c.TargetName = core.ReadLengthPrefixedStringN(r, 5)
	c.SystemIdentificationNumber = int(r.ReadUint(8))
	c.UtilizedTransformationMessageIndicator = int(r.ReadUint(10))
	c.PlateNumber = int(r.ReadUint(5))
	c.ComputationIndicator = int(r.ReadUint(4))
	c.HeightIndicator = int(r.ReadUint(2))
	c.LatitudeOfOriginArcSec = float64(r.ReadInt(19)) * 2
	c.LongitudeOfOriginArcSec = float64(r.ReadInt(20)) * 2
	c.NSExtensionArcSec = float64(r.ReadUint(14)) * 2
	c.EWExtensionArcSec = float64(r.ReadUint(14)) * 2
	c.TranslationXM = float64(r.ReadInt(23)) * 0.001
	c.TranslationYM = float64(r.ReadInt(23)) * 0.001
	c.TranslationZM = float64(r.ReadInt(23)) * 0.001
	c.RotationXArcSec = float64(r.ReadInt(32)) * 0.00002
	c.RotationYArcSec = float64(r.ReadInt(32)) * 0.00002
	c.RotationZArcSec = float64(r.ReadInt(32)) * 0.00002
	c.ScaleCorrectionPPM = float64(r.ReadInt(25)) * 0.00001
	return c
}

type CoordTransformTail struct {
	AddSemiMajorAxisSourceM    float64
	AddSemiMinorAxisSourceM    float64
	AddSemiMajorAxisTargetM    float64
	AddSemiMinorAxisTargetM    float64
	HorizontalQualityIndicator int
	VerticalQualityIndicator   int
}

func decodeCoordTransformTail(r *core.BitReader) CoordTransformTail {
	var t CoordTransformTail
	t.AddSemiMajorAxisSourceM = float64(r.ReadUint(24)) * 0.001
	t.AddSemiMinorAxisSourceM = float64(r.ReadUint(25)) * 0.001
	t.AddSemiMajorAxisTargetM = float64(r.ReadUint(24)) * 0.001
	t.AddSemiMinorAxisTargetM = float64(r.ReadUint(25)) * 0.001
	t.HorizontalQualityIndicator = int(r.ReadUint(3))
	t.VerticalQualityIndicator = int(r.ReadUint(3))
	return t
}
