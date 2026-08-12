package rtcm

func readLengthPrefixedString(r *BitReader) string {
	return readLengthPrefixedStringN(r, 8)
}

func readLengthPrefixedStringN(r *BitReader, lengthBits int) string {
	n := int(r.ReadUint(lengthBits))
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(r.ReadUint(8))
	}
	return string(b)
}
