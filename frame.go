package rtcm

import "fmt"

func ParseFrame(data []byte) (payload []byte, consumed int, err error) {
	skip := 0
	for skip < len(data) && data[skip] != 0xD3 {
		skip++
	}
	if skip >= len(data) {
		return nil, 0, fmt.Errorf("rtcm: no sync byte found")
	}
	data = data[skip:]
	if len(data) < 3 {
		return nil, 0, fmt.Errorf("rtcm: %d bytes requested, %d bytes returned", 3, len(data))
	}
	length := int(data[1]&0x03)<<8 | int(data[2])
	if length == 0 {
		return nil, 0, fmt.Errorf("rtcm: invalid payload size 0 bytes")
	}
	total := 3 + length + 3
	if len(data) < total {
		return nil, 0, fmt.Errorf("rtcm: %d bytes requested, %d bytes returned", total, len(data))
	}
	if CRC24Q(data[:total]) != 0 {
		return nil, 0, fmt.Errorf("rtcm: RTCM3 message invalid - failed CRC")
	}
	return data[3 : 3+length], skip + total, nil
}
