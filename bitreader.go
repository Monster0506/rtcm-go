package rtcm

type BitReader struct {
	data []byte
	pos  int
}

func NewBitReader(data []byte) *BitReader {
	return &BitReader{data: data}
}
