package msm

import (
	"fmt"
	"math/bits"

	"github.com/Monster0506/rtcm-go/core"
)

const speedOfLight = 299792458.0

type msmVariantShape struct {
	hasNms                 bool
	hasExtSatInfo          bool
	hasRoughPhaserangeRate bool
	hasFinePseudorange     bool
	fineRangeExtended      bool
	hasFinePhaserange      bool
	hasHalfCycleAmbiguity  bool
	hasCNR                 bool
	hasFinePhaserangeRate  bool
}

var msmVariantShapes = map[int]msmVariantShape{
	1: {hasFinePseudorange: true},
	2: {hasFinePhaserange: true, hasHalfCycleAmbiguity: true},
	3: {hasFinePseudorange: true, hasFinePhaserange: true, hasHalfCycleAmbiguity: true},
	4: {hasNms: true, hasFinePseudorange: true, hasFinePhaserange: true,
		hasHalfCycleAmbiguity: true, hasCNR: true},
	5: {hasNms: true, hasExtSatInfo: true, hasRoughPhaserangeRate: true,
		hasFinePseudorange: true, hasFinePhaserange: true, hasHalfCycleAmbiguity: true,
		hasCNR: true, hasFinePhaserangeRate: true},
	6: {hasNms: true, hasFinePseudorange: true, fineRangeExtended: true,
		hasFinePhaserange: true, hasHalfCycleAmbiguity: true, hasCNR: true},
	7: {hasNms: true, hasExtSatInfo: true, hasRoughPhaserangeRate: true,
		hasFinePseudorange: true, fineRangeExtended: true, hasFinePhaserange: true,
		hasHalfCycleAmbiguity: true, hasCNR: true, hasFinePhaserangeRate: true},
}

type MSMSatelliteData struct {
	SatelliteMaskIndex       int
	RoughRangeMs             *int
	ExtendedSatelliteInfo    *int
	RoughRangeModMs          float64
	RoughPhaserangeRateMPerS *float64
}

type MSMCellData struct {
	SatelliteMaskIndex      int
	SignalMaskIndex         int
	FinePseudorangeMs       *float64
	FinePhaserangeMs        *float64
	LockTimeIndicator       *int
	HalfCycleAmbiguity      *bool
	CNRdBHz                 *float64
	FinePhaserangeRateMPerS *float64
	PseudorangeM            *float64
	PhaserangeM             *float64
	PhaseRangeRateMPerS     *float64
}

type MSM struct {
	MSMHeader
	Satellites []MSMSatelliteData
	Cells      []MSMCellData
}

func DecodeMSM(payload []byte) (*MSM, error) {
	r := core.NewBitReader(payload)
	m := &MSM{}
	m.MessageType = int(r.ReadUint(12))
	constellation, ok := ConstellationForType(m.MessageType)
	if !ok {
		return nil, fmt.Errorf("rtcm: message type %d is not an MSM message", m.MessageType)
	}
	m.Constellation = constellation
	msmNumber := m.MessageType % 10
	shape, ok := msmVariantShapes[msmNumber]
	if !ok {
		return nil, fmt.Errorf("rtcm: message type %d has no known MSM variant (derived MSM number %d)", m.MessageType, msmNumber)
	}

	m.StationID = int(r.ReadUint(12))
	m.EpochTime = r.ReadUint(30)
	m.MultiMessage = r.ReadUint(1) != 0
	m.IODS = int(r.ReadUint(3))
	r.ReadUint(7)
	m.ClockSteering = int(r.ReadUint(2))
	m.ExternalClock = int(r.ReadUint(2))
	m.SmoothingIndicator = r.ReadUint(1) != 0
	m.SmoothingInterval = int(r.ReadUint(3))
	m.SatelliteMask = r.ReadUint(64)
	m.SignalMask = uint32(r.ReadUint(32))
	m.SatelliteCount = bits.OnesCount64(m.SatelliteMask)
	m.SignalCount = bits.OnesCount32(m.SignalMask)

	satIndices := maskSetBitIndices64(m.SatelliteMask, 64)
	sigIndices := maskSetBitIndices32(m.SignalMask, 32)

	cellWidth := m.SatelliteCount * m.SignalCount
	cellMask := make([]bool, cellWidth)
	remaining := cellWidth
	pos := 0
	for remaining > 0 {
		chunk := remaining
		if chunk > 64 {
			chunk = 64
		}
		v := r.ReadUint(chunk)
		for i := 0; i < chunk; i++ {
			cellMask[pos+i] = (v>>(chunk-1-i))&1 != 0
		}
		pos += chunk
		remaining -= chunk
	}
	m.CellCount = 0
	for _, b := range cellMask {
		if b {
			m.CellCount++
		}
	}

	m.Satellites = make([]MSMSatelliteData, m.SatelliteCount)
	for i := 0; i < m.SatelliteCount; i++ {
		m.Satellites[i].SatelliteMaskIndex = satIndices[i]
	}
	if shape.hasNms {
		for i := range m.Satellites {
			v := int(r.ReadUint(8))
			m.Satellites[i].RoughRangeMs = &v
		}
	}
	if shape.hasExtSatInfo {
		for i := range m.Satellites {
			v := int(r.ReadUint(4))
			m.Satellites[i].ExtendedSatelliteInfo = &v
		}
	}
	for i := range m.Satellites {
		m.Satellites[i].RoughRangeModMs = float64(r.ReadUint(10)) * core.TwoPow(-10)
	}
	if shape.hasRoughPhaserangeRate {
		for i := range m.Satellites {
			v := float64(r.ReadInt(14))
			m.Satellites[i].RoughPhaserangeRateMPerS = &v
		}
	}

	m.Cells = make([]MSMCellData, 0, m.CellCount)
	for satPos, sIdx := range satIndices {
		for sigPos, gIdx := range sigIndices {
			if !cellMask[satPos*m.SignalCount+sigPos] {
				continue
			}
			m.Cells = append(m.Cells, MSMCellData{
				SatelliteMaskIndex: sIdx,
				SignalMaskIndex:    gIdx,
			})
		}
	}

	if shape.hasFinePseudorange {
		for i := range m.Cells {
			var v float64
			if shape.fineRangeExtended {
				v = float64(r.ReadInt(20)) * core.TwoPow(-29)
			} else {
				v = float64(r.ReadInt(15)) * core.TwoPow(-24)
			}
			m.Cells[i].FinePseudorangeMs = &v
		}
	}
	if shape.hasFinePhaserange {
		for i := range m.Cells {
			var v float64
			if shape.fineRangeExtended {
				v = float64(r.ReadInt(24)) * core.TwoPow(-31)
			} else {
				v = float64(r.ReadInt(22)) * core.TwoPow(-29)
			}
			m.Cells[i].FinePhaserangeMs = &v
		}
	}
	if shape.hasFinePhaserange {
		for i := range m.Cells {
			var v int
			if shape.fineRangeExtended {
				v = int(r.ReadUint(10))
			} else {
				v = int(r.ReadUint(4))
			}
			m.Cells[i].LockTimeIndicator = &v
		}
	}
	if shape.hasHalfCycleAmbiguity {
		for i := range m.Cells {
			v := r.ReadUint(1) != 0
			m.Cells[i].HalfCycleAmbiguity = &v
		}
	}
	if shape.hasCNR {
		for i := range m.Cells {
			var v float64
			if shape.fineRangeExtended {
				v = float64(r.ReadUint(10)) * core.TwoPow(-4)
			} else {
				v = float64(r.ReadUint(6))
			}
			m.Cells[i].CNRdBHz = &v
		}
	}
	if shape.hasFinePhaserangeRate {
		for i := range m.Cells {
			v := float64(r.ReadInt(15)) * 0.0001
			m.Cells[i].FinePhaserangeRateMPerS = &v
		}
	}

	satByIndex := make(map[int]*MSMSatelliteData, len(m.Satellites))
	for i := range m.Satellites {
		satByIndex[m.Satellites[i].SatelliteMaskIndex] = &m.Satellites[i]
	}
	for i := range m.Cells {
		c := &m.Cells[i]
		sat := satByIndex[c.SatelliteMaskIndex]
		if sat == nil || sat.RoughRangeMs == nil {
			continue
		}
		nms := float64(*sat.RoughRangeMs)
		roughMs := nms + sat.RoughRangeModMs
		if c.FinePseudorangeMs != nil {
			v := speedOfLight / 1000 * (roughMs + *c.FinePseudorangeMs)
			c.PseudorangeM = &v
		}
		if c.FinePhaserangeMs != nil {
			v := speedOfLight / 1000 * (roughMs + *c.FinePhaserangeMs)
			c.PhaserangeM = &v
		}
		if sat.RoughPhaserangeRateMPerS != nil && c.FinePhaserangeRateMPerS != nil {
			v := *sat.RoughPhaserangeRateMPerS + *c.FinePhaserangeRateMPerS
			c.PhaseRangeRateMPerS = &v
		}
	}

	return m, nil
}

func maskSetBitIndices64(mask uint64, width int) []int {
	indices := make([]int, 0, bits.OnesCount64(mask))
	for i := 0; i < width; i++ {
		if mask&(1<<(width-1-i)) != 0 {
			indices = append(indices, i)
		}
	}
	return indices
}

func maskSetBitIndices32(mask uint32, width int) []int {
	indices := make([]int, 0, bits.OnesCount32(mask))
	for i := 0; i < width; i++ {
		if mask&(1<<(width-1-i)) != 0 {
			indices = append(indices, i)
		}
	}
	return indices
}
