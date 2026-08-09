package rtcm

type BitReader struct {
	data []byte
	pos  int
}

func NewBitReader(data []byte) *BitReader {
	return &BitReader{data: data}
}

func (r *BitReader) ReadUint(nbits int) uint64 {
	var v uint64
	for nbits > 0 {
		byteIdx := r.pos >> 3
		bitOff := r.pos & 7
		take := 8 - bitOff
		if take > nbits {
			take = nbits
		}
		shift := 8 - bitOff - take
		mask := byte((1 << take) - 1)
		v = v<<take | uint64((r.data[byteIdx]>>shift)&mask)
		r.pos += take
		nbits -= take
	}
	return v
}
