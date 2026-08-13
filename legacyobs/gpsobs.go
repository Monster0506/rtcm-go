package legacyobs

import "github.com/Monster0506/rtcm-go/core"

type GPSObservationHeader struct {
	MessageType         int
	StationID           int
	GPSEpochTOWMs       uint64
	SynchronousGNSSFlag bool
	SmoothingIndicator  bool
	SmoothingInterval   int
}

func decodeGPSObservationHeader(r *core.BitReader) (GPSObservationHeader, int) {
	var h GPSObservationHeader
	h.MessageType = int(r.ReadUint(12))
	h.StationID = int(r.ReadUint(12))
	h.GPSEpochTOWMs = r.ReadUint(30)
	h.SynchronousGNSSFlag = r.ReadUint(1) != 0
	numSats := int(r.ReadUint(5))
	h.SmoothingIndicator = r.ReadUint(1) != 0
	h.SmoothingInterval = int(r.ReadUint(3))
	return h, numSats
}

type GPSL1Observation struct {
	SatelliteID         int
	L1CodeIndicator     bool
	L1PseudorangeM      float64
	L1PhaserangeM       float64
	L1LockTimeIndicator int
}

func decodeGPSL1Observation(r *core.BitReader) GPSL1Observation {
	var o GPSL1Observation
	o.SatelliteID = int(r.ReadUint(6))
	o.L1CodeIndicator = r.ReadUint(1) != 0
	o.L1PseudorangeM = float64(r.ReadUint(24)) * 0.02
	o.L1PhaserangeM = float64(r.ReadInt(20)) * 0.0005
	o.L1LockTimeIndicator = int(r.ReadUint(7))
	return o
}

type GPSL1ExtendedObservation struct {
	GPSL1Observation
	L1PseudorangeAmbiguityM float64
	L1CNRDbHz               float64
}

func decodeGPSL1ExtendedObservation(r *core.BitReader) GPSL1ExtendedObservation {
	var o GPSL1ExtendedObservation
	o.GPSL1Observation = decodeGPSL1Observation(r)
	o.L1PseudorangeAmbiguityM = float64(r.ReadUint(8)) * 299792.458
	o.L1CNRDbHz = float64(r.ReadUint(8)) * 0.25
	return o
}

type GPSL2Observation struct {
	L2CodeIndicator       int
	L2MinusL1PseudorangeM float64
	L2PhaserangeM         float64
	L2LockTimeIndicator   int
}

func decodeGPSL2Observation(r *core.BitReader) GPSL2Observation {
	var o GPSL2Observation
	o.L2CodeIndicator = int(r.ReadUint(2))
	o.L2MinusL1PseudorangeM = float64(r.ReadInt(14)) * 0.02
	o.L2PhaserangeM = float64(r.ReadInt(20)) * 0.0005
	o.L2LockTimeIndicator = int(r.ReadUint(7))
	return o
}

type GPSL2ExtendedObservation struct {
	GPSL2Observation
	L2CNRDbHz float64
}

func decodeGPSL2ExtendedObservation(r *core.BitReader) GPSL2ExtendedObservation {
	var o GPSL2ExtendedObservation
	o.GPSL2Observation = decodeGPSL2Observation(r)
	o.L2CNRDbHz = float64(r.ReadUint(8)) * 0.25
	return o
}
