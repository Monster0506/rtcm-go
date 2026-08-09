package rtcm

const crc24qPoly = 0x1864CFB

var crc24qTable [256]uint32

func init() {
	for i := range crc24qTable {
		crc := uint32(i) << 16
		for b := 0; b < 8; b++ {
			crc <<= 1
			if crc&0x1000000 != 0 {
				crc ^= crc24qPoly
			}
		}
		crc24qTable[i] = crc & 0xFFFFFF
	}
}

func CRC24Q(data []byte) uint32 {
	var crc uint32
	for _, b := range data {
		crc = ((crc << 8) & 0xFFFFFF) ^ crc24qTable[byte(crc>>16)^b]
	}
	return crc
}
