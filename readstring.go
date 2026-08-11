package rtcm

func readLengthPrefixedString(r *BitReader) string {
	n := int(r.ReadUint(8))
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(r.ReadUint(8))
	}
	return string(b)
}
