package rtcm

import "fmt"

func ParseFrame(data []byte) (payload []byte, consumed int, err error) {
	if len(data) < 3 {
		return nil, 0, fmt.Errorf("rtcm: %d bytes requested, %d bytes returned", 3, len(data))
	}
	if data[0] != 0xD3 {
		return nil, 0, fmt.Errorf("rtcm: no sync byte found")
	}
	length := int(data[1]&0x03)<<8 | int(data[2])
	total := 3 + length + 3
	if len(data) < total {
		return nil, 0, fmt.Errorf("rtcm: %d bytes requested, %d bytes returned", total, len(data))
	}
	return data[3 : 3+length], total, nil
}
