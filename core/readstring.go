package core

func ReadLengthPrefixedString(r *BitReader) string {
	return ReadLengthPrefixedStringN(r, 8)
}

func ReadLengthPrefixedStringN(r *BitReader, lengthBits int) string {
	n := int(r.ReadUint(lengthBits))
	b := make([]byte, n)
	for i := range b {
		b[i] = byte(r.ReadUint(8))
	}
	return string(b)
}
