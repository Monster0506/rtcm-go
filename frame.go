package rtcm

import "fmt"

func ParseFrame(data []byte) (payload []byte, consumed int, err error) {
	skip := 0
	for {
		if skip >= len(data) {
			return nil, 0, fmt.Errorf("rtcm: no sync byte found")
		}
		if data[skip] != 0xD3 {
			skip++
			continue
		}
		if skip+1 >= len(data) {
			return nil, 0, fmt.Errorf("rtcm: serial stream terminated unexpectedly: %d bytes requested, %d bytes returned", 2, len(data)-skip)
		}
		if data[skip+1]&0xFC != 0 {
			skip += 2
			continue
		}
		break
	}
	data = data[skip:]
	if len(data) < 3 {
		return nil, 0, fmt.Errorf("rtcm: serial stream terminated unexpectedly: %d bytes requested, %d bytes returned", 3, len(data))
	}
	length := int(data[1])<<8 | int(data[2])
	if length == 0 {
		return nil, 0, fmt.Errorf("rtcm: invalid payload size 0 bytes")
	}
	if len(data)-3 < length {
		return nil, 0, fmt.Errorf("rtcm: serial stream terminated unexpectedly: %d bytes requested, %d bytes returned", length, len(data)-3)
	}
	if len(data)-3-length < 3 {
		return nil, 0, fmt.Errorf("rtcm: serial stream terminated unexpectedly: %d bytes requested, %d bytes returned", 3, len(data)-3-length)
	}
	total := 3 + length + 3
	if CRC24Q(data[:total]) != 0 {
		return nil, skip + total, fmt.Errorf("rtcm: RTCM3 message invalid - failed CRC")
	}
	return data[3 : 3+length], skip + total, nil
}
