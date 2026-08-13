package rtcm

import "fmt"

func PeekMessageType(payload []byte) int {
	if len(payload) < 2 {
		return -1
	}
	return int(NewBitReader(payload).ReadUint(12))
}

func Decode(payload []byte) (any, error) {
	msgType := PeekMessageType(payload)
	if msgType < 0 {
		return nil, fmt.Errorf("rtcm: payload too short to contain a message type")
	}
	if fn, ok := decoders[msgType]; ok {
		return fn(payload)
	}
	if _, ok := ConstellationForType(msgType); ok {
		return DecodeMSM(payload)
	}
	return nil, fmt.Errorf("rtcm: no decoder registered for message type %d", msgType)
}

var decoders = map[int]func([]byte) (any, error){
	1001: func(p []byte) (any, error) { return DecodeMsg1001(p) },
	1002: func(p []byte) (any, error) { return DecodeMsg1002(p) },
	1003: func(p []byte) (any, error) { return DecodeMsg1003(p) },
	1004: func(p []byte) (any, error) { return DecodeMsg1004(p) },
	1005: func(p []byte) (any, error) { return DecodeMsg1005(p) },
	1006: func(p []byte) (any, error) { return DecodeMsg1006(p) },
	1007: func(p []byte) (any, error) { return DecodeMsg1007(p) },
	1008: func(p []byte) (any, error) { return DecodeMsg1008(p) },
	1009: func(p []byte) (any, error) { return DecodeMsg1009(p) },
	1010: func(p []byte) (any, error) { return DecodeMsg1010(p) },
	1011: func(p []byte) (any, error) { return DecodeMsg1011(p) },
	1012: func(p []byte) (any, error) { return DecodeMsg1012(p) },
	1013: func(p []byte) (any, error) { return DecodeMsg1013(p) },
	1014: func(p []byte) (any, error) { return DecodeMsg1014(p) },
	1015: func(p []byte) (any, error) { return DecodeMsg1015(p) },
	1016: func(p []byte) (any, error) { return DecodeMsg1016(p) },
	1017: func(p []byte) (any, error) { return DecodeMsg1017(p) },
	1019: func(p []byte) (any, error) { return DecodeMsg1019(p) },
	1020: func(p []byte) (any, error) { return DecodeMsg1020(p) },
	1021: func(p []byte) (any, error) { return DecodeMsg1021(p) },
	1022: func(p []byte) (any, error) { return DecodeMsg1022(p) },
	1023: func(p []byte) (any, error) { return DecodeMsg1023(p) },
	1024: func(p []byte) (any, error) { return DecodeMsg1024(p) },
	1025: func(p []byte) (any, error) { return DecodeMsg1025(p) },
	1026: func(p []byte) (any, error) { return DecodeMsg1026(p) },
	1027: func(p []byte) (any, error) { return DecodeMsg1027(p) },
	1029: func(p []byte) (any, error) { return DecodeMsg1029(p) },
	1030: func(p []byte) (any, error) { return DecodeMsg1030(p) },
	1031: func(p []byte) (any, error) { return DecodeMsg1031(p) },
	1032: func(p []byte) (any, error) { return DecodeMsg1032(p) },
	1033: func(p []byte) (any, error) { return DecodeMsg1033(p) },
	1034: func(p []byte) (any, error) { return DecodeMsg1034(p) },
	1035: func(p []byte) (any, error) { return DecodeMsg1035(p) },
	1037: func(p []byte) (any, error) { return DecodeMsg1037(p) },
	1038: func(p []byte) (any, error) { return DecodeMsg1038(p) },
	1039: func(p []byte) (any, error) { return DecodeMsg1039(p) },
	1042: func(p []byte) (any, error) { return DecodeMsg1042(p) },
	1044: func(p []byte) (any, error) { return DecodeMsg1044(p) },
	1045: func(p []byte) (any, error) { return DecodeMsg1045(p) },
	1046: func(p []byte) (any, error) { return DecodeMsg1046(p) },
	1057: func(p []byte) (any, error) { return DecodeMsg1057(p) },
	1058: func(p []byte) (any, error) { return DecodeMsg1058(p) },
	1059: func(p []byte) (any, error) { return DecodeMsg1059(p) },
	1060: func(p []byte) (any, error) { return DecodeMsg1060(p) },
	1061: func(p []byte) (any, error) { return DecodeMsg1061(p) },
	1062: func(p []byte) (any, error) { return DecodeMsg1062(p) },
	1063: func(p []byte) (any, error) { return DecodeMsg1063(p) },
	1064: func(p []byte) (any, error) { return DecodeMsg1064(p) },
	1065: func(p []byte) (any, error) { return DecodeMsg1065(p) },
	1066: func(p []byte) (any, error) { return DecodeMsg1066(p) },
	1067: func(p []byte) (any, error) { return DecodeMsg1067(p) },
	1068: func(p []byte) (any, error) { return DecodeMsg1068(p) },
	1230: func(p []byte) (any, error) { return DecodeMsg1230(p) },
}
