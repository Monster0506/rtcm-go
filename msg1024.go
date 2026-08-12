package rtcm



type Msg1024GridPoint struct {
	NorthingResidualM float64
	EastingResidualM  float64
	HeightResidualM   float64
}



type Msg1024 struct {
	MessageType                    int
	SystemIdentificationNumber     int
	HorizontalShiftIndicator       bool
	VerticalShiftIndicator         bool
	NorthingOfOriginM              float64
	EastingOfOriginM               float64
	NSExtensionM                   float64
	EWExtensionM                   float64
	MeanNorthingOffsetM            float64
	MeanEastingOffsetM             float64
	MeanHeightOffsetM              float64
	GridPoints                     [16]Msg1024GridPoint
	HorizontalInterpolationMethod  int
	VerticalInterpolationMethod    int
	HorizontalGridQualityIndicator int
	VerticalGridQualityIndicator   int
	MJDNumber                      int
}

func DecodeMsg1024(payload []byte) (*Msg1024, error) {
	r := NewBitReader(payload)
	m := &Msg1024{}
	m.MessageType = int(r.ReadUint(12))
	m.SystemIdentificationNumber = int(r.ReadUint(8))
	m.HorizontalShiftIndicator = r.ReadUint(1) != 0
	m.VerticalShiftIndicator = r.ReadUint(1) != 0
	m.NorthingOfOriginM = float64(r.ReadInt(25)) * 10
	m.EastingOfOriginM = float64(r.ReadUint(26)) * 10
	m.NSExtensionM = float64(r.ReadUint(12)) * 10
	m.EWExtensionM = float64(r.ReadUint(12)) * 10
	m.MeanNorthingOffsetM = float64(r.ReadInt(10)) * 0.01
	m.MeanEastingOffsetM = float64(r.ReadInt(10)) * 0.01
	m.MeanHeightOffsetM = float64(r.ReadInt(15)) * 0.01
	for i := 0; i < 16; i++ {
		m.GridPoints[i] = Msg1024GridPoint{
			NorthingResidualM: float64(r.ReadInt(9)) * 0.001,
			EastingResidualM:  float64(r.ReadInt(9)) * 0.001,
			HeightResidualM:   float64(r.ReadInt(9)) * 0.001,
		}
	}
	m.HorizontalInterpolationMethod = int(r.ReadUint(2))
	m.VerticalInterpolationMethod = int(r.ReadUint(2))
	m.HorizontalGridQualityIndicator = int(r.ReadUint(3))
	m.VerticalGridQualityIndicator = int(r.ReadUint(3))
	m.MJDNumber = int(r.ReadUint(16))
	return m, nil
}
