package msm

import (
	"fmt"
	"math/bits"

	"github.com/Monster0506/rtcm-go/core"
)

func ConstellationForType(msgType int) (string, bool) {
	switch {
	case msgType >= 1071 && msgType <= 1077:
		return "GPS", true
	case msgType >= 1081 && msgType <= 1087:
		return "GLONASS", true
	case msgType >= 1091 && msgType <= 1097:
		return "Galileo", true
	case msgType >= 1101 && msgType <= 1107:
		return "SBAS", true
	case msgType >= 1111 && msgType <= 1117:
		return "QZSS", true
	case msgType >= 1121 && msgType <= 1127:
		return "BeiDou", true
	default:
		return "", false
	}
}

type MSMHeader struct {
	MessageType        int
	Constellation      string
	StationID          int
	EpochTime          uint64
	MultiMessage       bool
	IODS               int
	ClockSteering      int
	ExternalClock      int
	SmoothingIndicator bool
	SmoothingInterval  int
	SatelliteMask      uint64
	SignalMask         uint32
	SatelliteCount     int
	SignalCount        int
	CellCount          int
}

func DecodeMSMHeader(payload []byte) (*MSMHeader, error) {
	r := core.NewBitReader(payload)
	m := &MSMHeader{}
	m.MessageType = int(r.ReadUint(12))
	constellation, ok := ConstellationForType(m.MessageType)
	if !ok {
		return nil, fmt.Errorf("rtcm: message type %d is not an MSM message", m.MessageType)
	}
	m.Constellation = constellation
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
	m.CellCount = readCellMaskPopcount(r, m.SatelliteCount*m.SignalCount)
	return m, nil
}

func readCellMaskPopcount(r *core.BitReader, width int) int {
	count := 0
	remaining := width
	for remaining > 0 {
		chunk := remaining
		if chunk > 64 {
			chunk = 64
		}
		count += bits.OnesCount64(r.ReadUint(chunk))
		remaining -= chunk
	}
	return count
}
