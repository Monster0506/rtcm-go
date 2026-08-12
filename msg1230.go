package rtcm





type Msg1230 struct {
	MessageType            int
	StationID              int
	CodePhaseBiasIndicator bool
	FDMASignalsMask        int
	L1CACodePhaseBiasM     *float64
	L1PCodePhaseBiasM      *float64
	L2CACodePhaseBiasM     *float64
	L2PCodePhaseBiasM      *float64
}

func DecodeMsg1230(payload []byte) (*Msg1230, error) {
	r := NewBitReader(payload)
	m := &Msg1230{}
	m.MessageType = int(r.ReadUint(12))
	m.StationID = int(r.ReadUint(12))
	m.CodePhaseBiasIndicator = r.ReadUint(1) != 0
	r.ReadUint(3) 
	m.FDMASignalsMask = int(r.ReadUint(4))

	readBiasIfPresent := func(bitPos uint) *float64 {
		if m.FDMASignalsMask&(1<<bitPos) == 0 {
			return nil
		}
		v := float64(r.ReadInt(16)) * 0.02
		return &v
	}
	m.L1CACodePhaseBiasM = readBiasIfPresent(3)
	m.L1PCodePhaseBiasM = readBiasIfPresent(2)
	m.L2CACodePhaseBiasM = readBiasIfPresent(1)
	m.L2PCodePhaseBiasM = readBiasIfPresent(0)
	return m, nil
}
