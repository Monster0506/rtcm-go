package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type GLONASSObservationHeader struct {
	MessageType         int
	StationID           int
	GLONASSEpochTkMs    uint64
	SynchronousGNSSFlag bool
	SmoothingIndicator  bool
	SmoothingInterval   int
}

func decodeGLONASSObservationHeader(r *core.BitReader) (GLONASSObservationHeader, int) {
	var h GLONASSObservationHeader
	h.MessageType = int(r.ReadUint(12))
	h.StationID = int(r.ReadUint(12))
	h.GLONASSEpochTkMs = r.ReadUint(27)
	h.SynchronousGNSSFlag = r.ReadUint(1) != 0
	numSats := int(r.ReadUint(5))
	h.SmoothingIndicator = r.ReadUint(1) != 0
	h.SmoothingInterval = int(r.ReadUint(3))
	return h, numSats
}

type GLONASSL1Observation struct {
	SatelliteID            int
	L1CodeIndicator        bool
	FrequencyChannelNumber int
	L1PseudorangeM         float64
	L1PhaserangeM          float64
	L1LockTimeIndicator    int
}

func decodeGLONASSL1Observation(r *core.BitReader) GLONASSL1Observation {
	var o GLONASSL1Observation
	o.SatelliteID = int(r.ReadUint(6))
	o.L1CodeIndicator = r.ReadUint(1) != 0
	o.FrequencyChannelNumber = int(r.ReadUint(5))
	o.L1PseudorangeM = float64(r.ReadUint(25)) * 0.02
	o.L1PhaserangeM = float64(r.ReadInt(20)) * 0.0005
	o.L1LockTimeIndicator = int(r.ReadUint(7))
	return o
}

type GLONASSL1ExtendedObservation struct {
	GLONASSL1Observation
	L1PseudorangeAmbiguityM float64
	L1CNRDbHz               float64
}

func decodeGLONASSL1ExtendedObservation(r *core.BitReader) GLONASSL1ExtendedObservation {
	var o GLONASSL1ExtendedObservation
	o.GLONASSL1Observation = decodeGLONASSL1Observation(r)
	o.L1PseudorangeAmbiguityM = float64(r.ReadUint(7)) * 599584.916
	o.L1CNRDbHz = float64(r.ReadUint(8)) * 0.25
	return o
}

type GLONASSL2Observation struct {
	L2CodeIndicator       int
	L2MinusL1PseudorangeM float64
	L2PhaserangeM         float64
	L2LockTimeIndicator   int
}

func decodeGLONASSL2Observation(r *core.BitReader) GLONASSL2Observation {
	var o GLONASSL2Observation
	o.L2CodeIndicator = int(r.ReadUint(2))
	o.L2MinusL1PseudorangeM = float64(r.ReadInt(14)) * 0.02
	o.L2PhaserangeM = float64(r.ReadInt(20)) * 0.0005
	o.L2LockTimeIndicator = int(r.ReadUint(7))
	return o
}

type GLONASSL2ExtendedObservation struct {
	GLONASSL2Observation
	L2CNRDbHz float64
}

func decodeGLONASSL2ExtendedObservation(r *core.BitReader) GLONASSL2ExtendedObservation {
	var o GLONASSL2ExtendedObservation
	o.GLONASSL2Observation = decodeGLONASSL2Observation(r)
	o.L2CNRDbHz = float64(r.ReadUint(8)) * 0.25
	return o
}
