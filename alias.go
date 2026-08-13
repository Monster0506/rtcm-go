package rtcm

import "github.com/Monster0506/rtcm-go/core"

type BitReader = core.BitReader

var (
	NewBitReader = core.NewBitReader
	CRC24Q       = core.CRC24Q
	ParseFrame   = core.ParseFrame
)
