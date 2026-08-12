package rtcm

var ssrUpdateIntervals = [16]float64{
	1, 2, 5, 10, 15, 30, 60, 120, 240, 300, 600, 900, 1800, 3600, 7200, 10800,
}

type SSRHeader struct {
	MessageType              int
	EpochTimeS               int
	UpdateIntervalCode       int
	UpdateIntervalS          float64
	MultipleMessageIndicator bool
	IODSSR                   int
	SSRProviderID            int
	SSRSolutionID            int
	NumSatellites            int
}

func decodeSSRHeader(r *BitReader, epochBits int) SSRHeader {
	var h SSRHeader
	h.MessageType = int(r.ReadUint(12))
	h.EpochTimeS = int(r.ReadUint(epochBits))
	h.UpdateIntervalCode = int(r.ReadUint(4))
	h.UpdateIntervalS = ssrUpdateIntervals[h.UpdateIntervalCode]
	h.MultipleMessageIndicator = r.ReadUint(1) != 0
	h.IODSSR = int(r.ReadUint(4))
	h.SSRProviderID = int(r.ReadUint(16))
	h.SSRSolutionID = int(r.ReadUint(4))
	h.NumSatellites = int(r.ReadUint(6))
	return h
}

type SSROrbitHeader struct {
	MessageType              int
	EpochTimeS               int
	UpdateIntervalCode       int
	UpdateIntervalS          float64
	MultipleMessageIndicator bool
	SatelliteReferenceDatum  bool
	IODSSR                   int
	SSRProviderID            int
	SSRSolutionID            int
	NumSatellites            int
}

func decodeSSROrbitHeader(r *BitReader, epochBits int) SSROrbitHeader {
	var h SSROrbitHeader
	h.MessageType = int(r.ReadUint(12))
	h.EpochTimeS = int(r.ReadUint(epochBits))
	h.UpdateIntervalCode = int(r.ReadUint(4))
	h.UpdateIntervalS = ssrUpdateIntervals[h.UpdateIntervalCode]
	h.MultipleMessageIndicator = r.ReadUint(1) != 0
	h.SatelliteReferenceDatum = r.ReadUint(1) != 0
	h.IODSSR = int(r.ReadUint(4))
	h.SSRProviderID = int(r.ReadUint(16))
	h.SSRSolutionID = int(r.ReadUint(4))
	h.NumSatellites = int(r.ReadUint(6))
	return h
}

type SSROrbitCorrection struct {
	SatelliteID             int
	IOD                     int
	DeltaRadialM            float64
	DeltaAlongTrackM        float64
	DeltaCrossTrackM        float64
	DotDeltaRadialMPerS     float64
	DotDeltaAlongTrackMPerS float64
	DotDeltaCrossTrackMPerS float64
}

func decodeSSROrbitCorrection(r *BitReader, satIDBits int) SSROrbitCorrection {
	var c SSROrbitCorrection
	c.SatelliteID = int(r.ReadUint(satIDBits))
	c.IOD = int(r.ReadUint(8))
	c.DeltaRadialM = float64(r.ReadInt(22)) * 0.0001
	c.DeltaAlongTrackM = float64(r.ReadInt(20)) * 0.0004
	c.DeltaCrossTrackM = float64(r.ReadInt(20)) * 0.0004
	c.DotDeltaRadialMPerS = float64(r.ReadInt(21)) * 0.000001
	c.DotDeltaAlongTrackMPerS = float64(r.ReadInt(19)) * 0.000004
	c.DotDeltaCrossTrackMPerS = float64(r.ReadInt(19)) * 0.000004
	return c
}

type SSRClockCorrection struct {
	SatelliteID        int
	DeltaClockC0M      float64
	DeltaClockC1MPerS  float64
	DeltaClockC2MPerS2 float64
}

func decodeSSRClockCorrection(r *BitReader, satIDBits int) SSRClockCorrection {
	var c SSRClockCorrection
	c.SatelliteID = int(r.ReadUint(satIDBits))
	c.DeltaClockC0M = float64(r.ReadInt(22)) * 0.0001
	c.DeltaClockC1MPerS = float64(r.ReadInt(21)) * 0.000001
	c.DeltaClockC2MPerS2 = float64(r.ReadInt(27)) * 0.00000002
	return c
}

type SSRCodeBias struct {
	SignalTrackingMode int
	CodeBiasM          float64
}

type SSRSatelliteCodeBiases struct {
	SatelliteID int
	CodeBiases  []SSRCodeBias
}

func decodeSSRSatelliteCodeBiases(r *BitReader, satIDBits int) SSRSatelliteCodeBiases {
	var s SSRSatelliteCodeBiases
	s.SatelliteID = int(r.ReadUint(satIDBits))
	numBiases := int(r.ReadUint(5))
	s.CodeBiases = make([]SSRCodeBias, numBiases)
	for i := 0; i < numBiases; i++ {
		s.CodeBiases[i] = SSRCodeBias{
			SignalTrackingMode: int(r.ReadUint(5)),
			CodeBiasM:          float64(r.ReadInt(14)) * 0.01,
		}
	}
	return s
}

type SSRURA struct {
	SatelliteID int
	URACode     int
	URAClass    int
	URAValue    int
}

func decodeSSRURA(r *BitReader, satIDBits int) SSRURA {
	var u SSRURA
	u.SatelliteID = int(r.ReadUint(satIDBits))
	u.URACode = int(r.ReadUint(6))
	u.URAClass = u.URACode >> 3
	u.URAValue = u.URACode & 0x7
	return u
}

type SSRCombinedCorrection struct {
	SatelliteID             int
	IOD                     int
	DeltaRadialM            float64
	DeltaAlongTrackM        float64
	DeltaCrossTrackM        float64
	DotDeltaRadialMPerS     float64
	DotDeltaAlongTrackMPerS float64
	DotDeltaCrossTrackMPerS float64
	DeltaClockC0M           float64
	DeltaClockC1MPerS       float64
	DeltaClockC2MPerS2      float64
}

func decodeSSRCombinedCorrection(r *BitReader, satIDBits int) SSRCombinedCorrection {
	var c SSRCombinedCorrection
	c.SatelliteID = int(r.ReadUint(satIDBits))
	c.IOD = int(r.ReadUint(8))
	c.DeltaRadialM = float64(r.ReadInt(22)) * 0.0001
	c.DeltaAlongTrackM = float64(r.ReadInt(20)) * 0.0004
	c.DeltaCrossTrackM = float64(r.ReadInt(20)) * 0.0004
	c.DotDeltaRadialMPerS = float64(r.ReadInt(21)) * 0.000001
	c.DotDeltaAlongTrackMPerS = float64(r.ReadInt(19)) * 0.000004
	c.DotDeltaCrossTrackMPerS = float64(r.ReadInt(19)) * 0.000004
	c.DeltaClockC0M = float64(r.ReadInt(22)) * 0.0001
	c.DeltaClockC1MPerS = float64(r.ReadInt(21)) * 0.000001
	c.DeltaClockC2MPerS2 = float64(r.ReadInt(27)) * 0.00000002
	return c
}

type SSRHighRateClock struct {
	SatelliteID              int
	HighRateClockCorrectionM float64
}

func decodeSSRHighRateClock(r *BitReader, satIDBits int) SSRHighRateClock {
	var c SSRHighRateClock
	c.SatelliteID = int(r.ReadUint(satIDBits))
	c.HighRateClockCorrectionM = float64(r.ReadInt(22)) * 0.0001
	return c
}
